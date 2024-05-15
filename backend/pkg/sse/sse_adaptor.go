package sse

import (
	"fmt"
	"net/http"
	"time"

	"CVSeeker/pkg/gpt"
)

type ISSEClient interface {
	StreamRunResponse(w http.ResponseWriter, responseChan <-chan gpt.RunResponse, errorChan <-chan error)
}

type SSEClient struct{}

func NewSSEClient() ISSEClient {
	return &SSEClient{}
}

func (client *SSEClient) StreamRunResponse(w http.ResponseWriter, responseChan <-chan gpt.RunResponse, errorChan <-chan error) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		select {
		case runResp := <-responseChan:
			if runResp.Status == "completed" || runResp.Status == "failed" {
				fmt.Fprintf(w, "event: %s\n", runResp.Status)
				fmt.Fprintf(w, "data: %s\n\n", runResp)
				flusher.Flush()
				return
			}
			fmt.Fprintf(w, "event: update\n")
			fmt.Fprintf(w, "data: %s\n\n", runResp)
			flusher.Flush()
		case err := <-errorChan:
			fmt.Fprintf(w, "event: error\n")
			fmt.Fprintf(w, "data: %s\n\n", err.Error())
			flusher.Flush()
			return
		case <-time.After(2 * time.Second):
			// Send a keep-alive comment to keep the connection open
			fmt.Fprintf(w, ": keep-alive\n\n")
			flusher.Flush()
		}
	}
}
