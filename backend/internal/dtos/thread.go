package dtos

type Thread struct {
	ID        string `json:"id"`
	UpdatedAt int64  `json:"updated_at"`
	Name      string `json:"name"`
}
