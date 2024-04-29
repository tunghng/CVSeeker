package elasticsearch

type ElasticResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	URL     string `json:"url"`
}
