package main

import (
	"os"
	"path/filepath"

	_ "embed"
)

//go:embed .github/ISSUE_TEMPLATE/bug_report.md
var bugReportTemplate []byte

//go:embed .github/ISSUE_TEMPLATE/feature_request.md
var featureRequestTemplate []byte

//go:embed .github/pull_request_template.md
var pullRequestTemplate []byte

func createIssueTemplates() error {
	issueTemplateDir := filepath.Join(".github", "ISSUE_TEMPLATE")

	err := os.MkdirAll(issueTemplateDir, os.ModePerm)
	if err != nil {
		return err
	}

	bugReportFile, err := os.Create(filepath.Join(issueTemplateDir, "bug_report.md"))
	if err != nil {
		return err
	}
	defer bugReportFile.Close()

	featureRequestFile, err := os.Create(filepath.Join(issueTemplateDir, "feature_request.md"))
	if err != nil {
		return err
	}
	defer featureRequestFile.Close()

	if _, err = bugReportFile.Write(bugReportTemplate); err != nil {
		return err
	}
	_, err = featureRequestFile.Write(featureRequestTemplate)
	return err
}

func createPullRequestTemplate() error {
	err := os.MkdirAll(".github", os.ModePerm)
	if err != nil {
		return err
	}

	pullRequestFile, err := os.Create(filepath.Join(".github", "pull_request_template.md"))
	if err != nil {
		return err
	}

	_, err = pullRequestFile.Write(pullRequestTemplate)
	return err
}
