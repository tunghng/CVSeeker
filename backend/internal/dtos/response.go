package dtos

import (
	"CVSeeker/internal/meta"
	"encoding/json"
)

// Response response custom message
type Response struct {
	Meta meta.Meta   `json:"meta"`
	Data interface{} `json:"data" swaggertype:"object"`
}

type ResumeProcessingResult struct {
	Id     string `json:"id"` // Assume each resume has a reference or ID.
	Status string `json:"status"`
}

type PaginationResponse struct {
	Meta           meta.Meta       `json:"meta"`
	PaginationInfo *PaginationInfo `json:"pagination"`
	Data           interface{}     `json:"data" swaggertype:"object"`
	Suggestion     interface{}     `json:"suggestion,omitempty"`
}

type PaginationInfo struct {
	PageSize     int64 `json:"pageSize"`
	PageOffset   int64 `json:"pageOffset"`
	TotalRecords int64 `json:"totalRecords"`
	TotalPages   int64 `json:"totalPages"`
}

type ListParam struct {
	PageOffset int64
	PageSize   int64
	Pagination bool
	Preload    bool
	OrderBy    *string
}

const (
	QueryValueAll  = "*"
	QueryValueNone = "-"
)

const (
	DefaultPageSize int64 = 20
	MinPageSize     int64 = 10
	MaxPageSize     int64 = 500
)

type PaginationFilter struct {
	PageOffset int64
	PageSize   int64
	UserId     int64
	ChannelIds []int64
	OrderBy    string `json:"orderBy"`
}

type StreamedEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"` // Using RawMessage to handle varying data structures
}

type DeltaMessage struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Delta  struct {
		Content []struct {
			Index int    `json:"index"`
			Type  string `json:"type"`
			Text  struct {
				Value string `json:"value"`
			} `json:"text"`
		} `json:"content"`
	} `json:"delta"`
}
