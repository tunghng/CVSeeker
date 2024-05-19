package elasticsearch

type ElkResumeDTO struct {
	Content   ResumeSummaryDTO `json:"content"`
	Embedding []float32        `json:"embedding"`
}

type ResumeSummaryDTO struct {
	Id                string              `json:"id"`
	Summary           string              `json:"summary"`
	Skills            []string            `json:"skills"`
	BasicInfo         BasicInfo           `json:"basic_info"`
	WorkExperience    []WorkExperience    `json:"work_experience"`
	ProjectExperience []ProjectExperience `json:"project_experience"`
	Award             []Award             `json:"award"`
	URL               string              `json:"url"`
	Point             float64             `json:"point"`
}

type BasicInfo struct {
	FullName       string   `json:"full_name"`
	University     string   `json:"university"`
	EducationLevel string   `json:"education_level"` // BS, MS, or PhD
	Majors         []string `json:"majors"`
	GPA            float64  `json:"gpa"`
}

type WorkExperience struct {
	JobTitle   string `json:"job_title"`
	Company    string `json:"company"`
	Location   string `json:"location"`
	Duration   string `json:"duration"` // Could be changed to a more structured format if necessary
	JobSummary string `json:"job_summary"`
}

type ProjectExperience struct {
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
}

type Award struct {
	AwardName string `json:"award_name"`
}
