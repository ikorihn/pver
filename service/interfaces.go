package service

import "net/http"

type JiraRepository interface {
	NewRequest(method, urlStr string, body interface{}) (*http.Request, error)
	Do(req *http.Request, v interface{}) (*http.Response, error)
}
