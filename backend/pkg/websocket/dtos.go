package websocket

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
