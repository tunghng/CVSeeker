package services

import (
	"fmt"
	"go.uber.org/dig"
	"grabber-match/internal/meta"
	"grabber-match/pkg/gpt"
	"strings"
)

type IDataProcessingService interface {
	SummarizeResume(fullText string) (*meta.BasicResponse, error)
}

type DataProcessingService struct {
	gptClient gpt.IGptAdaptorClient
}

type DataProcessingServiceArgs struct {
	dig.In
	GptClient gpt.IGptAdaptorClient
}

func NewDataProcessingService(args DataProcessingServiceArgs) IDataProcessingService {
	return &DataProcessingService{
		gptClient: args.GptClient,
	}
}

func (_this *DataProcessingService) SummarizeResume(fullText string) (*meta.BasicResponse, error) {
	prompt := generatePrompt(fullText)
	model := "gpt-3.5-turbo" // Or any specific model you wish to use.

	responseText, err := _this.gptClient.AskGPT(prompt, model)
	if err != nil {
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    200,
			Message: "Resume summarized successfully",
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
