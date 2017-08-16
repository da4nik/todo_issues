package integrations

import (
	"github.com/da4nik/todo_issues/integrations/github"
	"github.com/da4nik/todo_issues/types"
)

// Issue integrations interface
type Issue interface {
	CreateIssue(string, string, int) types.IntegrationResponse
}

// TODO: #15 Optain access_token from params/config/env vars
// TODO: #15 https://api.github.com/repos/da4nik/todo_issues/issues/15
var integration = github.New("access_token", "da4nik", "todo_issues")

// CreateIssue created issue, only for github now
func CreateIssue(title, filename string, lineNumber int) types.IntegrationResponse {
	// TODO: #16 Allow to create issues in multiple systems
	// TODO: #16 https://api.github.com/repos/da4nik/todo_issues/issues/16
	return integration.CreateIssue(title, filename, lineNumber)
}
