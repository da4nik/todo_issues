package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/da4nik/todo_issues/config"
	"github.com/da4nik/todo_issues/integrations"
)

var validFileName = regexp.MustCompile(`\.(go|txt)$`)
var todoText = regexp.MustCompile(`(#|\/\/)\s*TODO:\s*(.+)$`)
var issueExists = regexp.MustCompile(`(#|\/\/)\s*TODO:\s*#\S+\s(.+)$`)

func visit(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fi.IsDir() ||
		!validFileName.MatchString(fi.Name()) ||
		strings.HasPrefix(path, "vendor/") {

		return nil
	}

	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bytes.NewReader(read))
	contents := make([]string, 0)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++

		// Leaving existing issues as is
		if issueExists.MatchString(scanner.Text()) {
			contents = append(contents, scanner.Text())
			continue
		}

		matches := todoText.FindStringSubmatch(scanner.Text())
		if len(matches) < 3 {
			contents = append(contents, scanner.Text())
			continue
		}

		comment := string(matches[1])
		title := string(matches[2])
		if title == "" {
			contents = append(contents, scanner.Text())
			continue
		}

		fmt.Printf("Creating issue: %s\n", title)
		issue := integrations.CreateIssue(title, path, lineNumber)
		if issue.ID == "" {
			return nil
		}

		lineWithID := fmt.Sprintf("%s TODO: #%s %s", comment, issue.ID, title)
		additionText := fmt.Sprintf("%s TODO: #%s %s", comment, issue.ID, issue.IssueLink)

		updatedLine := todoText.ReplaceAllString(scanner.Text(), lineWithID)

		contents = append(contents, updatedLine)
		if config.Additions && issue.IssueLink != "" {
			addition := todoText.ReplaceAllString(scanner.Text(), additionText)
			contents = append(contents, addition)
		}
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return scannerErr
	}

	// Add new line to the EOF
	contents = append(contents, "")

	err = ioutil.WriteFile(path, []byte(strings.Join(contents, "\n")), fi.Mode().Perm())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	config.Load()

	err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
}
