package meta

type BasicResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
	Plan interface{} `json:"plan,omitempty"`
}
