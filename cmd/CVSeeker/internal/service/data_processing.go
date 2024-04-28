package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"grabber-match/cmd/CVSeeker/internal/cfg"
	"grabber-match/internal/ginLogger"
	"grabber-match/internal/meta"
	"grabber-match/internal/models"
	"grabber-match/internal/repositories"
	"grabber-match/pkg/db"
	"grabber-match/pkg/elasticsearch"
	"grabber-match/pkg/gpt"
	"strings"
	"time"
)

type IDataProcessingService interface {
	ProcessData(c *gin.Context, fullText string, file []byte) (*meta.BasicResponse, error)
}

type DataProcessingService struct {
	db         *db.DB
	gptClient  gpt.IGptAdaptorClient
	resumeRepo repositories.IResumeRepository
	elkClient  elasticsearch.CoreElkClient
}

type DataProcessingServiceArgs struct {
	dig.In
	DB         *db.DB `name:"talentAcquisitionDB"`
	GptClient  gpt.IGptAdaptorClient
	ResumeRepo repositories.IResumeRepository
	ElkClient  elasticsearch.CoreElkClient
}

func NewDataProcessingService(args DataProcessingServiceArgs) IDataProcessingService {
	return &DataProcessingService{
		db:         args.DB,
		gptClient:  args.GptClient,
		resumeRepo: args.ResumeRepo,
		elkClient:  args.ElkClient,
	}
}

func (_this *DataProcessingService) ProcessData(c *gin.Context, fullText string, file []byte) (*meta.BasicResponse, error) {
	prompt := generatePrompt(fullText)
	model := "gpt-3.5-turbo"

	// Summarize resume text by making request to OpenAI
	responseText, err := _this.gptClient.AskGPT(prompt, model)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to summarize using GPT: %v", err)
		return nil, err
	}

	// Document for elasticsearch
	elkData := map[string]interface{}{
		"full_text":        fullText,
		"download_link":    "http://example.com/mockresume.pdf", // Mock download link
		"vector_embedding": "MockVectorEmbedding123",            // Mock vector embedding
		"created_at":       time.Now().Format(time.RFC3339),
		"updated_at":       time.Now().Format(time.RFC3339),
	}

	// Upload to Elasticsearch
	err = _this.elkClient.SaveToElasticsearch(c, viper.GetString(cfg.ElasticsearchDocumentIndex), elkData)
	if err != nil {
		return nil, fmt.Errorf("failed to upload resume data to Elasticsearch: %v", err)
	}

	//
	resume := &models.Resume{
		FullText:        fullText,
		VectorEmbedding: "MockVectorEmbedding123",            // Mock vector embedding
		DownloadLink:    "http://example.com/mockresume.pdf", // Mock download link
	}

	// Create resume in database
	_, err = _this.resumeRepo.Create(_this.db, resume)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create resume record: %v", err)
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
