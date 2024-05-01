package meta

import "net/http"

// Meta is metadata of merchant integration services.
type Meta struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	ExtraData interface{} `json:"extra,omitempty"`
}

// MetaOK is successfully metadata.
var OK = New(http.StatusOK, nil)

// New returns new Meta.
func New(code int, extra interface{}, msg ...string) Meta {
	cd := 200
	if code > 0 {
		cd = code
	}

	m := http.StatusText(cd)
	if len(msg) > 0 {
		m = msg[0]
	}

	return Meta{
		Code:      cd,
		Message:   m,
		ExtraData: extra,
	}
}

// Error metadata.
type Error struct {
	Meta Meta `json:"meta"`
}
