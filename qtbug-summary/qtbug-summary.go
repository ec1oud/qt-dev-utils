// Copyright (C) 2024 Shawn Rutledge <s@ecloud.org>
// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("give the issue number")
	}
	jiraClient, _ := jira.NewClient(nil, "https://bugreports.qt.io/")
	issue, _, _ := jiraClient.Issue.Get(os.Args[1], nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
}

