package meta

type BasicResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
	Plan interface{} `json:"plan,omitempty"`
}

type Plan struct {
	Plans []PlanDetail `json:"plans"`
}

type PlanDetail struct {
	IsCurrentPlan bool    `json:"isCurrentPlan"`
	ExpireTime    *int64  `json:"expireTime"`
	PlanId        int64   `json:"planId"`
	Type          string  `json:"type"`
	Unit          string  `json:"unit"`
	Label         string  `json:"label"`
	Price         float64 `json:"price"`
	PriceOrg      float64 `json:"priceOrg"`
}
