package gograveyard

import (
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
	TotalCount int `json:"total_count"`
}

// OpenIssueCount will return the number of open issues for a repository
// This API endpoint used was found here: https://github.com/isaacs/github/issues/536#issuecomment-532884919
func (p Project) OpenIssuesCount() (int, error) {
	var count int

	resp, err := p.Client.Get(fmt.Sprintf("%s/search/issues?q=repo:%s/%s+type:issue+state:open&per_page=1", BaseURL, p.Owner, p.Repo))
	if err != nil {
		return count, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return count, err
	}

	var s SearchIssues

	err = json.Unmarshal(body, &s)
	if err != nil {
		return count, err
	}

	return s.TotalCount, nil
}
