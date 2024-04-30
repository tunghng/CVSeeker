package services

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/internal/models"
	"CVSeeker/internal/repositories"
	"CVSeeker/pkg/db"
	"CVSeeker/pkg/elasticsearch"
	"CVSeeker/pkg/gpt"
	"CVSeeker/pkg/huggingface"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"strings"
)

type IDataProcessingService interface {
	ProcessData(c *gin.Context, fullText string, file []byte) (*meta.BasicResponse, error)
}

type DataProcessingService struct {
	db            *db.DB
	gptClient     gpt.IGptAdaptorClient
	resumeRepo    repositories.IResumeRepository
	elasticClient elasticsearch.IElasticsearchClient
	hfClient      huggingface.IHuggingFaceClient
}

type DataProcessingServiceArgs struct {
	dig.In
	DB            *db.DB `name:"talentAcquisitionDB"`
	GptClient     gpt.IGptAdaptorClient
	ResumeRepo    repositories.IResumeRepository
	ElasticClient elasticsearch.IElasticsearchClient
	HfClient      huggingface.IHuggingFaceClient
}

func NewDataProcessingService(args DataProcessingServiceArgs) IDataProcessingService {
	return &DataProcessingService{
		db:            args.DB,
		gptClient:     args.GptClient,
		resumeRepo:    args.ResumeRepo,
		elasticClient: args.ElasticClient,
		hfClient:      args.HfClient,
	}
}

func (_this *DataProcessingService) ProcessData(c *gin.Context, fullText string, file []byte) (*meta.BasicResponse, error) {
	prompt := generatePrompt(fullText)
	model := viper.GetString(cfg.ChatGptModel)
	elasticDocumentName := viper.GetString(cfg.ElasticsearchDocumentIndex)
	textEmbeddingModel := viper.GetString(cfg.HuggingfaceModel)

	mockDownloadLink := "http://example.com/mockresume.pdf"

	// Summarize resume text by making request to OpenAI
	responseText, err := _this.gptClient.AskGPT(prompt, model)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to summarize using GPT: %v", err)
		return nil, err
	}

	// Prepare resume for database
	resume := &models.Resume{
		FullText:     responseText,
		DownloadLink: mockDownloadLink,
	}

	// Create resume in database
	databaseResume, err := _this.resumeRepo.Create(_this.db, resume)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create resume record: %v", err)
		return nil, err
	}

	// Create the vector representation of text
	vectorEmbedding, err := _this.hfClient.GetTextEmbedding(responseText, textEmbeddingModel)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to get text embedding: %v", err)
		return nil, err
	}

	// Prepare document for Elasticsearch, including the database resumeId
	elkResume := map[string]interface{}{
		"content":   responseText,
		"embedding": vectorEmbedding,
		"url":       databaseResume.DownloadLink,
	}

	// Index resume in Elasticsearch
	err = _this.elasticClient.AddDocument(c, elasticDocumentName, fmt.Sprintf("%d", databaseResume.ResumeId), elkResume)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to upload resume data to Elasticsearch: %v", err)
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    200,
			Message: "Resume processed and file uploaded successfully",
		},
		Data: responseText,
	}

	return response, nil
}

func generatePrompt(fullText string) string {
	var sb strings.Builder
	sb.WriteString("Given the full text of a resume below, please perform a detailed summary focusing on essential information that will be used for text embedding models aimed at similarity matching with job descriptions. Output the summary in a structured and concise format suitable for embedding, which includes the following sections:\n\n")
	sb.WriteString(fmt.Sprintf("Full text of the resume:\n%s\n\n", fullText))
	sb.WriteString("Summary should include sections on these, you MUST NOT MAKE UP ANY INFORMATION, do not add the section if there is no content related:\n")
	sb.WriteString("1. Header: Candidate's full name, email address, phone number, and any professional online profiles such as LinkedIn or GitHub.\n")
	sb.WriteString("2. Summary: Concise professional summary or objective encapsulating the candidate’s career goals and main qualifications.\n")
	sb.WriteString("3. Education: Each educational qualification including the degree, institution, year of graduation, and any honors or special recognitions.\n")
	sb.WriteString("4. Work Experience:\n    * Include for each job the employer, job title, dates of employment, and a bulleted list of key responsibilities and significant achievements.\n")
	sb.WriteString("5. Skills: List all relevant technical and soft skills, specific technologies or tools, formatted as a comma-separated list.\n")
	sb.WriteString("6. Certifications: Any relevant certifications with the name, issuing organization, and date of certification.\n")
	sb.WriteString("7. Projects: Significant projects including title, brief description, technologies used, and project's impact or outcome.\n")
	sb.WriteString("8. Publications: Any publications including title, co-authors, publication venue, and publication date.\n")
	sb.WriteString("9. Languages: Languages known along with proficiency level (e.g., fluent, intermediate).\n")
	sb.WriteString("10. Professional Affiliations: Memberships or roles in professional organizations, including the organization name, role nature, and dates of involvement.\n")
	return sb.String()
}
