package gograveyard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "https://api.github.com"

type Project struct {
	Client *http.Client
	Owner  string
	Repo   string
}

type SearchIssues struct {
	TotalCount int `json:"total_count"` //nolint:tagliatelle
}

// OpenIssueCount will return the number of open issues for a repository
// This API endpoint used was found here:
// https://github.com/isaacs/github/issues/536#issuecomment-532884919
func (p Project) OpenIssuesCount() (int, error) {
	var count int

	url := fmt.Sprintf("%s/search/issues?q=repo:%s/%s+type:issue+state:open&per_page=1",
		BaseURL, p.Owner, p.Repo)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return count, fmt.Errorf("failed to create new request with context: %w", err)
	}

	resp, err := p.Client.Do(req)
	if err != nil {
		return count, fmt.Errorf("http GET failed for open issue count: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return count, fmt.Errorf("failed to read response body: %w", err)
	}

	var s SearchIssues

	err = json.Unmarshal(body, &s)
	if err != nil {
		return count, fmt.Errorf("failed to unmarshal open issue count: %w", err)
	}

	return s.TotalCount, nil
}
