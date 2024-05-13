package services

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/internal/dtos"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/internal/repositories"
	"CVSeeker/pkg/aws"
	"CVSeeker/pkg/db"
	"CVSeeker/pkg/elasticsearch"
	"CVSeeker/pkg/huggingface"
	"CVSeeker/pkg/summarizer"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"net/http"
	"strings"
	"sync"
)

type IDataProcessingService interface {
	ProcessData(c *gin.Context, fullText string, file string) (*meta.BasicResponse, error)
	ProcessDataBatch(c *gin.Context, resumes []dtos.ResumeData) (*meta.BasicResponse, error)
}

type DataProcessingService struct {
	db            *db.DB
	gptClient     summarizer.ISummarizerAdaptorClient
	resumeRepo    repositories.IResumeRepository
	elasticClient elasticsearch.IElasticsearchClient
	hfClient      huggingface.IHuggingFaceClient
	s3Client      *aws.S3Client
}

type DataProcessingServiceArgs struct {
	dig.In
	DB            *db.DB `name:"talentAcquisitionDB"`
	GptClient     summarizer.ISummarizerAdaptorClient
	ResumeRepo    repositories.IResumeRepository
	ElasticClient elasticsearch.IElasticsearchClient
	HfClient      huggingface.IHuggingFaceClient
	S3Client      *aws.S3Client
}

func NewDataProcessingService(args DataProcessingServiceArgs) IDataProcessingService {
	return &DataProcessingService{
		db:            args.DB,
		gptClient:     args.GptClient,
		resumeRepo:    args.ResumeRepo,
		elasticClient: args.ElasticClient,
		hfClient:      args.HfClient,
		s3Client:      args.S3Client,
	}
}

func (_this *DataProcessingService) ProcessData(c *gin.Context, fullText string, file string) (*meta.BasicResponse, error) {
	elasticDocumentName := viper.GetString(cfg.ElasticsearchDocumentIndex)

	// Create the ElkResume DTO
	elkResume, err := _this.createElkResume(c, fullText, file)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create elastic document: %v", err)
		return nil, err
	}

	// Upload the document and get its ID
	documentID, err := _this.elasticClient.AddDocument(c.Request.Context(), elasticDocumentName, elkResume)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to upload resume data to Elasticsearch: %v", err)
		return nil, err
	}

	// Prepare the result using the ResumeProcessingResult struct
	result := dtos.ResumeProcessingResult{
		Id:     documentID,
		Status: "Success", // Set status as "Success" since there was no error
	}

	// Prepare the basic response with the result included
	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Resume processed and file uploaded successfully",
		},
		Data: result,
	}

	return response, nil
}

func (_this *DataProcessingService) ProcessDataBatch(c *gin.Context, resumes []dtos.ResumeData) (*meta.BasicResponse, error) {
	var wg sync.WaitGroup
	results := make(chan *dtos.ResumeProcessingResult, len(resumes))
	errors := make(chan error, len(resumes))

	elasticDocumentName := viper.GetString(cfg.ElasticsearchDocumentIndex)

	for _, resume := range resumes {
		wg.Add(1)
		go func(res dtos.ResumeData) {
			defer wg.Done()
			elkResume, err := _this.createElkResume(c, res.Content, res.FileBytes)
			if err != nil {
				ginLogger.Gin(c).Errorf("failed to create elk resume: %v", err)
				errors <- err
				results <- &dtos.ResumeProcessingResult{Status: "Failed"}
				return
			}

			documentID, err := _this.elasticClient.AddDocument(c.Request.Context(), elasticDocumentName, elkResume)
			if err != nil {
				ginLogger.Gin(c).Errorf("failed to upload resume data to Elasticsearch: %v", err)
				errors <- err
				results <- &dtos.ResumeProcessingResult{Id: documentID, Status: "Failed"}
				return
			}

			results <- &dtos.ResumeProcessingResult{Id: documentID, Status: "Success"}
		}(resume)
	}

	wg.Wait()
	close(results)
	close(errors)

	// Check for errors and aggregate results
	var finalResults []dtos.ResumeProcessingResult
	for err := range errors {
		if err != nil {
			return nil, err
		}
	}
	for res := range results {
		finalResults = append(finalResults, *res)
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Batch processing completed successfully",
		},
		Data: finalResults,
	}

	return response, nil
}

func (_this *DataProcessingService) createElkResume(c *gin.Context, fullText string, file string) (*elasticsearch.ElkResumeDTO, error) {
	prompt := generatePrompt(fullText)
	model := viper.GetString(cfg.ChatGptModel)
	textEmbeddingModel := viper.GetString(cfg.HuggingfaceModel)
	//awsBucketName := viper.GetString(cfg.AwsBucket)

	//fileBytes, err := base64.StdEncoding.DecodeString(file)
	//if err != nil {
	//	ginLogger.Gin(c).Errorf("failed to decode file: %v", err)
	//	return nil, err
	//}

	// Parse resume text to JSON format by making request to OpenAI
	responseText, err := _this.gptClient.AskGPT(prompt, model)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to summarize using GPT: %v", err)
		return nil, err
	}

	// Upload file to S3 and get the URL
	//key := fmt.Sprintf("%d.docx", time.Now().Unix())
	//
	//fileURL, err := _this.s3Client.UploadFile(c.Request.Context(), awsBucketName, key, file)
	//if err != nil {
	//	ginLogger.Gin(c).Errorf("failed to upload file to S3: %v", err)
	//	return nil, err
	//}

	var resumeSummary elasticsearch.ResumeSummaryDTO
	//var resumeSummary map[string]interface{}
	if err := json.Unmarshal([]byte(responseText), &resumeSummary); err != nil {
		ginLogger.Gin(c).Errorf("failed to parse JSON response: %v", err)
		return nil, err
	}
	resumeSummary.URL = "https://cvseeker-bucket.s3.ap-southeast-2.amazonaws.com/1714643484.pdf"

	// Create the vector representation of text
	vectorEmbedding, err := _this.hfClient.GetTextEmbedding(fullText, textEmbeddingModel)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to get text embedding: %v", err)
		return nil, err
	}

	// Prepare the document for Elasticsearch
	elkResume := &elasticsearch.ElkResumeDTO{
		Content:   resumeSummary,
		Embedding: vectorEmbedding,
	}

	return elkResume, nil
}

func generatePrompt(fullText string) string {
	var sb strings.Builder
	sb.WriteString("Full text of the resume:\n\n")
	sb.WriteString(fullText)
	sb.WriteString("\n\nPlease transform the above resume text into a well-structured JSON. The JSON should have the following structure and order:\n\n")
	sb.WriteString(`{
  "summary": "[Provide a concise professional summary based on the resume. Include key skills and experiences.]",
  "skills": ["List all relevant skills derived from the resume, each as a separate element in the array."],
  "basic_info": {
    "full_name": "[Invent a full name that sounds realistic and appropriate for the professional field]",
    "university": "[Generate a university name that fits the education level and field of study]",
    "education_level": "[Assign an education level, e.g., BS, MS, PhD, appropriate for the resume context]",
    "majors": ["Create a list of majors that align with the professional background and education level]",
    "GPA": [Generate a GPA as a number that is plausible for the given educational background, or use null if not applicable]
  },
  "work_experience": [
    {
      "job_title": "[Title of the position]",
      "company": "[Name of the company]",
      "location": "[Location of the job]",
      "duration": "[Duration of the job in years or months, e.g., '2 years']",
      "job_summary": "[A brief summary of job responsibilities and achievements]"
    }
  ],
  "project_experience": [
    {
      "project_name": "[Name of the project]",
      "project_description": "[A detailed description of the project, including technologies used and outcomes]"
    }
  ],
  "award": [
    {
      "award_name": "[Name of any award received, empty array if none]"
    }
  ]
}`)
	sb.WriteString("\n\nAll details in the 'basic_info' section should be invented but must sound logical and realistic, appropriate for the professional context. Ensure the details are consistent with typical professional and educational backgrounds relevant to the data in the rest of the resume. For other sections, ensure all entries are derived from the resume's content, maintaining consistency and accuracy with the original information. Provide clear, precise language to avoid ambiguities and ensure data types match the expected format.")
	return sb.String()
}
