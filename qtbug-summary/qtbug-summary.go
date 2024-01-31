// Copyright (C) 2024 Shawn Rutledge <s@ecloud.org>
// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"bufio"
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"os"
	"regexp"
	"strings"
)

var jiraClient, _ = jira.NewClient(nil, "https://bugreports.qt.io/")

func bugDesc (bugID string) string {
	issue, _, err := jiraClient.Issue.Get(bugID, nil)

	if (err != nil) {
		panic(err)
	}

	priority, _, _ := strings.Cut(issue.Fields.Priority.Name, ":")

	return fmt.Sprintf("%s %s %s: %+v", issue.Key, issue.Fields.Status.Name, priority, issue.Fields.Summary)
}

func main() {
	re := regexp.MustCompile("QTBUG-[0-9]+")
	if len(os.Args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			matches := re.FindAllString(scanner.Text(), -1)
			for _, bugID := range matches {
				fmt.Printf("%s\n", bugDesc(bugID))
			}
		}
	} else {
		for _, bugID := range os.Args[1:] {
			fmt.Printf("%s\n", bugDesc(bugID))
		}
	}
}
