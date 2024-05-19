package services

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/internal/dtos"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/internal/models"
	"CVSeeker/internal/repositories"
	"CVSeeker/pkg/aws"
	"CVSeeker/pkg/db"
	"CVSeeker/pkg/elasticsearch"
	"CVSeeker/pkg/huggingface"
	"CVSeeker/pkg/summarizer"
	"CVSeeker/pkg/websocket"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type IDataProcessingService interface {
	ProcessData(c *gin.Context, fullText string, file string, uuid string) (*meta.BasicResponse, error)
	ProcessDataBatch(c *gin.Context, resumes []dtos.ResumeData, isLinkedin bool) (*meta.BasicResponse, error)
	GetAllUploads(c *gin.Context) (*meta.BasicResponse, error)
}

type DataProcessingService struct {
	db            *db.DB
	gptClient     summarizer.ISummarizerAdaptorClient
	resumeRepo    repositories.IResumeRepository
	uploadRepo    repositories.IUploadRepository
	elasticClient elasticsearch.IElasticsearchClient
	hfClient      huggingface.IHuggingFaceClient
	s3Client      *aws.S3Client
}

type DataProcessingServiceArgs struct {
	dig.In
	DB            *db.DB `name:"talentAcquisitionDB"`
	GptClient     summarizer.ISummarizerAdaptorClient
	ResumeRepo    repositories.IResumeRepository
	UploadRepo    repositories.IUploadRepository
	ElasticClient elasticsearch.IElasticsearchClient
	HfClient      huggingface.IHuggingFaceClient
	S3Client      *aws.S3Client
}

func NewDataProcessingService(args DataProcessingServiceArgs) IDataProcessingService {
	return &DataProcessingService{
		db:            args.DB,
		gptClient:     args.GptClient,
		resumeRepo:    args.ResumeRepo,
		uploadRepo:    args.UploadRepo,
		elasticClient: args.ElasticClient,
		hfClient:      args.HfClient,
		s3Client:      args.S3Client,
	}
}

func (_this *DataProcessingService) ProcessData(c *gin.Context, fullText string, file string, uuid string) (*meta.BasicResponse, error) {
	// This method now schedules the processing in the background and immediately returns a response
	go func() {
		elasticDocumentName := viper.GetString(cfg.ElasticsearchDocumentIndex)

		initialUpload := &models.Upload{
			Status: "Processing", // Initial status
			UUID:   uuid,
		}

		createdUpload, err := _this.uploadRepo.Create(_this.db, initialUpload)
		if err != nil {
			ginLogger.Gin(c).Errorf("Failed to log initial upload: %v", err)
			return
		}

		// Assume createElkResume is an existing method that prepares the data for Elasticsearch
		elkResume, err := _this.createElkResume(c, fullText, file, false)
		if err != nil {
			_this.uploadRepo.Update(_this.db, &models.Upload{ID: createdUpload.ID, Status: "Failed"})
			ginLogger.Gin(c).Errorf("failed to create elastic document: %v", err)
			return
		}

		// Add document to Elasticsearch and handle the response
		documentID, err := _this.elasticClient.AddDocument(c, elasticDocumentName, elkResume)
		if err != nil {
			_this.uploadRepo.Update(_this.db, &models.Upload{ID: createdUpload.ID, Status: "Failed"})
			ginLogger.Gin(c).Errorf("failed to upload resume data to Elasticsearch: %v", err)
			return
		}

		_this.uploadRepo.Update(_this.db, &models.Upload{ID: createdUpload.ID, DocumentID: documentID, Status: "Success"})

		websocket.BroadcastNotification("All documents have been processed successfully.")
	}()

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Processing request received and is being processed",
		},
		Data: nil,
	}

	// Return an immediate response to indicate that processing has started
	return response, nil
}

func (_this *DataProcessingService) ProcessDataBatch(c *gin.Context, resumes []dtos.ResumeData, isLinkedin bool) (*meta.BasicResponse, error) {
	elasticDocumentName := viper.GetString(cfg.ElasticsearchDocumentIndex)

	if isLinkedin {
		linkedInUrls := make([]string, len(resumes))
		for i, resume := range resumes {
			linkedInUrls[i] = resume.FileBytes // Assuming FileBytes contains the LinkedIn URL
		}

		processedResumes, err := fetchLinkedInData(linkedInUrls)
		if err != nil {
			return nil, err
		}
		resumes = processedResumes // Replace or merge as necessary
	}

	// Start processing in the background
	go func() {
		var wg sync.WaitGroup
		results := make(chan *dtos.ResumeProcessingResult, len(resumes))
		errors := make(chan error, len(resumes))

		for _, resume := range resumes {
			wg.Add(1)
			go func(res dtos.ResumeData) {
				defer wg.Done()

				// Create initial upload record for each document
				initialUpload := &models.Upload{
					Status: "Processing",
				}

				createdUpload, err := _this.uploadRepo.Create(_this.db, initialUpload)
				if err != nil {
					ginLogger.Gin(c).Errorf("Failed to log initial upload: %v", err)
					return
				}

				elkResume, err := _this.createElkResume(c, res.Content, res.FileBytes, isLinkedin)
				if err != nil {
					_this.uploadRepo.Update(_this.db, &models.Upload{ID: createdUpload.ID, Status: "Failed"})
					ginLogger.Gin(c).Errorf("failed to create elk resume: %v", err)
					errors <- err
					return
				}

				documentID, err := _this.elasticClient.AddDocument(c, elasticDocumentName, elkResume)
				if err != nil {
					_this.uploadRepo.Update(_this.db, &models.Upload{ID: createdUpload.ID, Status: "Failed"})
					ginLogger.Gin(c).Errorf("failed to upload resume data to Elasticsearch: %v", err)
					errors <- err
					return
				}

				_this.uploadRepo.Update(_this.db, &models.Upload{ID: createdUpload.ID, DocumentID: documentID, Status: "Success"})
			}(resume)
		}

		wg.Wait()
		close(results)
		close(errors)

		websocket.BroadcastNotification("All documents have been processed successfully.")
	}()

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Processing request received and is being processed",
		},
		Data: nil,
	}

	// Return an immediate response to indicate that processing has started
	return response, nil
}

func (_this *DataProcessingService) GetAllUploads(c *gin.Context) (*meta.BasicResponse, error) {
	uploads, err := _this.uploadRepo.GetAll(_this.db)
	if err != nil {
		ginLogger.Gin(c).Errorf("Failed to retrieve upload records: %v", err)
		return nil, err
	}

	// Convert uploads to DTOs
	var uploadsDTO []dtos.UploadDTO
	for _, upload := range uploads {
		dto := dtos.UploadDTO{
			DocumentID: upload.DocumentID,
			Status:     upload.Status,
			Name:       upload.Name,
			CreatedAt:  upload.CreatedAt.Unix(),
			UUID:       upload.UUID,
		}
		uploadsDTO = append(uploadsDTO, dto)
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Upload records retrieved successfully",
		},
		Data: uploadsDTO,
	}

	return response, nil
}

func fetchLinkedInData(urls []string) ([]dtos.ResumeData, error) {
	apiUrl := "http://0.0.0.0:8000/api/getfulltext/?list_url=" + strings.Join(urls, ",")
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var linkedInResumes struct {
		Resumes []dtos.ResumeData `json:"resumes"`
	}
	err = json.Unmarshal(body, &linkedInResumes)
	if err != nil {
		return nil, err
	}

	// Here you can process the LinkedIn data further if needed
	// For example, modifying the content structure, or integrating additional data

	return linkedInResumes.Resumes, nil
}

func (_this *DataProcessingService) createElkResume(c *gin.Context, fullText string, file string, isLinkedin bool) (*elasticsearch.ElkResumeDTO, error) {
	var prompt string
	if isLinkedin == false {
		prompt = generatePrompt(fullText)
	} else {
		prompt = generatePromptLinkedin(fullText)
	}
	model := viper.GetString(cfg.ChatGptModel)
	textEmbeddingModel := viper.GetString(cfg.HuggingfaceModel)
	awsBucketName := viper.GetString(cfg.AwsBucket)

	// Parse resume text to JSON format by making request to OpenAI
	responseText, err := _this.gptClient.AskGPT(prompt, model)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to summarize using GPT: %v", err)
		return nil, err
	}

	// Upload file to S3 and get the URL
	key := fmt.Sprintf("%d", time.Now().Unix())

	var fileURL string

	if isLinkedin == false {
		fileBytes, err := base64.StdEncoding.DecodeString(file)
		if err != nil {
			ginLogger.Gin(c).Errorf("failed to decode file: %v", err)
			return nil, err
		}
		fileURL, err = _this.s3Client.UploadFile(c, awsBucketName, key, fileBytes)
		if err != nil {
			ginLogger.Gin(c).Errorf("failed to upload file to S3: %v", err)
			return nil, err
		}
	} else {
		fileURL = file
	}
	var resumeSummary elasticsearch.ResumeSummaryDTO
	if err := json.Unmarshal([]byte(responseText), &resumeSummary); err != nil {
		ginLogger.Gin(c).Errorf("failed to parse JSON response: %v", err)
		return nil, err
	}
	resumeSummary.URL = fileURL

	embeddingText := generateFulltext(resumeSummary)
	// Create the vector representation of text
	vectorEmbedding, err := _this.hfClient.GetTextEmbedding(embeddingText, textEmbeddingModel)
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

func generatePromptLinkedin(fullText string) string {
	var sb strings.Builder
	sb.WriteString("Full text of the resume:\n\n")
	sb.WriteString(fullText)
	sb.WriteString("\n\nPlease transform the above resume text into a well-structured JSON. The JSON should have the following structure and order:\n\n")
	sb.WriteString(`{
  "summary": "[Provide a concise professional summary based on the resume. Include key skills and experiences.]",
  "skills": ["List all relevant skills derived from the resume, each as a separate element in the array."],
  "basic_info": {
    "full_name": "[Full name]",
    "university": "[University name]",
    "education_level": "[Education level, e.g., BS, MS, PhD, appropriate for the resume context]",
    "majors": ["A list of majors that align with the professional background and education level]",
    "GPA": [A GPA as a number that is plausible for the given educational background, or use null if not applicable]
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

func generateFulltext(resume elasticsearch.ResumeSummaryDTO) string {
	var fullTextContent strings.Builder
	fullTextContent.WriteString(fmt.Sprintf("Summary: %s; Skills: %v; ", resume.Summary, resume.Skills))
	fullTextContent.WriteString(fmt.Sprintf("Education: %s, %s, GPA: %.2f; ", resume.BasicInfo.University, resume.BasicInfo.EducationLevel, resume.BasicInfo.GPA))
	fullTextContent.WriteString("Work Experience: ")
	for _, work := range resume.WorkExperience {
		fullTextContent.WriteString(fmt.Sprintf("%s at %s, %s; ", work.JobTitle, work.Company, work.Duration))
	}
	fullTextContent.WriteString("Projects: ")
	for _, project := range resume.ProjectExperience {
		fullTextContent.WriteString(fmt.Sprintf("%s: %s; ", project.ProjectName, project.ProjectDescription))
	}
	fullTextContent.WriteString("Awards: ")
	for _, award := range resume.Award {
		fullTextContent.WriteString(fmt.Sprintf("%s; ", award.AwardName))
	}
	return fullTextContent.String()
}
