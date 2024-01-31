// Copyright (C) 2024 Shawn Rutledge <s@ecloud.org>
// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"bufio"
	"flag"
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"os"
	"regexp"
	"strings"
)

var jiraClient, _ = jira.NewClient(nil, "https://bugreports.qt.io/")
var relatedBugsFlag bool = false

func bugDesc (issue *jira.Issue) string {
	priority, _, _ := strings.Cut(issue.Fields.Priority.Name, ":")
	resolution := "Unrslvd"
	if issue.Fields.Resolution != nil {
		resolution = issue.Fields.Resolution.Name
	}

	return fmt.Sprintf("%s %s %s\t%+v",
		issue.Fields.Status.Name, resolution,
		priority, issue.Fields.Summary)
}

func describe (bugID string, indent string) {
	issue, _, err := jiraClient.Issue.Get(bugID, nil)

	if (err != nil) {
		fmt.Fprintf(os.Stderr, "%s: %s\n", bugID, err)
		return
	}

	fmt.Printf("%s%s\n", indent, bugDesc(issue))
	if (relatedBugsFlag) {
		for _, link := range issue.Fields.IssueLinks {
			if link.InwardIssue != nil {
				fmt.Printf("%s\t%s %s %s\n", indent, link.Type.Name, link.InwardIssue.Key, bugDesc(link.InwardIssue))
			}
			if link.OutwardIssue != nil {
				fmt.Printf("%s\t%s %s %s\n", indent, link.Type.Name, link.OutwardIssue.Key, bugDesc(link.OutwardIssue))
			}
		}
	}
}

func describeWithID (bugID string, indent string) {
	issue, _, err := jiraClient.Issue.Get(bugID, nil)

	if (err != nil) {
		fmt.Fprintf(os.Stderr, "%s: %s\n", bugID, err)
		return
	}

	fmt.Printf("%s%s %s\n", indent, bugID, bugDesc(issue))
	if (relatedBugsFlag) {
		for _, link := range issue.Fields.IssueLinks {
			if link.InwardIssue != nil {
				fmt.Printf("%s\t%s %s %s\n", indent, link.Type.Name, link.InwardIssue.Key, bugDesc(link.InwardIssue))
			}
			if link.OutwardIssue != nil {
				fmt.Printf("%s\t%s %s %s\n", indent, link.Type.Name, link.OutwardIssue.Key, bugDesc(link.OutwardIssue))
			}
		}
	}
}

func main() {
	flag.BoolVar(&relatedBugsFlag, "r", false, "show related bugs")
	flag.Parse()

	re := regexp.MustCompile("QTBUG-[0-9]+")
	if flag.NArg() == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Print(scanner.Text())
			matches := re.FindAllString(scanner.Text(), -1)
			if (len(matches) > 1) {
				fmt.Println()
				for _, bugID := range matches {
					describeWithID(bugID, "\t")
				}
			} else if (len(matches) == 1) {
				describe(matches[0], " ")
			} else {
				fmt.Println()
			}
		}
	} else {
		for _, bugID := range flag.Args() {
			describeWithID(bugID, "")
		}
	}
}
