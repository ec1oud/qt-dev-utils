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

	return fmt.Sprintf("%s %s: %+v", issue.Fields.Status.Name, priority, issue.Fields.Summary)
}

func main() {
	re := regexp.MustCompile("QTBUG-[0-9]+")
	if len(os.Args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Print(scanner.Text())
			matches := re.FindAllString(scanner.Text(), -1)
			if (len(matches) > 1) {
				fmt.Println()
				for _, bugID := range matches {
					fmt.Printf("\t%s %s\n", bugID, bugDesc(bugID))
				}
			} else if (len(matches) == 1) {
				fmt.Printf(" %s\n", bugDesc(matches[0]))
			} else {
				fmt.Println()
			}
		}
	} else {
		for _, bugID := range os.Args[1:] {
			fmt.Printf("%s\n", bugDesc(bugID))
		}
	}
}
