package infra

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setup() (*http.ServeMux, *httptest.Server) {
	// Test server
	testMux := http.NewServeMux()
	testServer := httptest.NewServer(testMux)
	return testMux, testServer
}

func Test_searchRepository_Search(t *testing.T) {
	testMux, testServer := setup()
	defer testServer.Close()
	testMux.HandleFunc("/rest/api/2/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Method not allowed. %v", r.Method)
		}

		wantURL := "/rest/api/2/search?jql=type+%3D+Bug+and+Status+NOT+IN+%28Resolved%29"
		if got := r.URL.String(); !strings.HasPrefix(got, wantURL) {
			t.Errorf("Request URL: %v, want %v", got, wantURL)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"expand": "schema,names","startAt": 1,"maxResults": 40,"total": 6,"issues": [{"expand": "html","id": "10230","self": "http://kelpie9:8081/rest/api/2/issue/BULK-62","key": "BULK-62","fields": {"summary": "testing","timetracking": null,"issuetype": {"self": "http://kelpie9:8081/rest/api/2/issuetype/5","id": "5","description": "The sub-task of the issue","iconUrl": "http://kelpie9:8081/images/icons/issue_subtask.gif","name": "Sub-task","subtask": true},"customfield_10071": null}},{"expand": "html","id": "10004","self": "http://kelpie9:8081/rest/api/2/issue/BULK-47","key": "BULK-47","fields": {"summary": "Cheese v1 2.0 issue","timetracking": null,"issuetype": {"self": "http://kelpie9:8081/rest/api/2/issuetype/3","id": "3","description": "A task that needs to be done.","iconUrl": "http://kelpie9:8081/images/icons/task.gif","name": "Task","subtask": false}}}]}`)
	})

	testRepository, _ := NewJiraRepository(nil, testServer.URL, "", "")
	result, err := testRepository.SearchIssues("type = Bug and Status NOT IN (Resolved)")

	if err != nil {
		t.Errorf("Error given: %s", err)
		return
	}

	if len(result.Issues) != 2 {
		t.Errorf("Issues size is not match: %v", result.Issues)
		return
	}

}
