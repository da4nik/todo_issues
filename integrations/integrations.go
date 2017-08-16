package integrations

import (
	"github.com/da4nik/todo_issues/integrations/github"
	"github.com/da4nik/todo_issues/types"
)

// Issue integrations interface
type Issue interface {
	CreateIssue(string, string, int) types.IntegrationResponse
}

// TODO: #19 Optain access_token from params/config/env vars
// TODO: #19 https://github.com/da4nik/todo_issues/issues/19
var integration = github.New("access_key", "da4nik", "todo_issues")

// CreateIssue created issue, only for github now
func CreateIssue(title, filename string, lineNumber int) types.IntegrationResponse {
	// TODO: #20 Allow to create issues in multiple systems
	// TODO: #20 https://github.com/da4nik/todo_issues/issues/20
	return integration.CreateIssue(title, filename, lineNumber)
}
