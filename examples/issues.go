package main

import (
	"context"
	"log"
	"os"

	"github.com/fgazat/tracker"
	"github.com/fgazat/tracker/entities"
)

// https://yandex.cloud/en/docs/tracker/local-fields
type IssueWithLocalFields struct {
	entities.Issue
	PullRequestURL string `json:"670bd2973c363d6b91754268--prUrl,omitempty"`
}

func main() {
	token := os.Getenv("TRACKER_TOKEN")
	orgID := os.Getenv("XCLOUD_ORG_ID")

	client := tracker.New(token, tracker.WithXCloudOrgID(orgID))

	var issue IssueWithLocalFields
	err := client.GetIssueByKey(context.Background(), "TEST-1", &issue)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("URL: %s", issue.PullRequestURL)
	log.Printf("Issue: %+v", issue)
}
