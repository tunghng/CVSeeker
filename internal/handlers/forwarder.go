package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ForwardInfo struct {
	Target   *url.URL
	Username string
	Password string
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func NewSingleHostReverseProxy(forwardInfo ForwardInfo, forwardPath string) *httputil.ReverseProxy {
	target := forwardInfo.Target
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.Header.Set("Authorization",
			fmt.Sprintf("Basic %s",
				base64.StdEncoding.EncodeToString(
					[]byte(fmt.Sprintf("%s:%s", forwardInfo.Username, forwardInfo.Password)))))
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, forwardPath)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
	}
	return &httputil.ReverseProxy{Director: director}
}
