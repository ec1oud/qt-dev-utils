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

func main() {
	re := regexp.MustCompile("QTBUG-[0-9]+")
	bugID := ""
	if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		matches := re.FindAllString(line, -1)
		bugID = matches[0]
	} else {
		bugID = os.Args[1]
	}
	jiraClient, _ := jira.NewClient(nil, "https://bugreports.qt.io/")
	issue, _, err := jiraClient.Issue.Get(bugID, nil)

	if (err != nil) {
		panic(err)
	}

	priority, _, _ := strings.Cut(issue.Fields.Priority.Name, ":")

	fmt.Printf("%s: %s: %+v\n", issue.Key, priority, issue.Fields.Summary)
}
