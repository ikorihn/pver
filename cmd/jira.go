/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/kyokomi/emoji"
	"github.com/manifoldco/promptui"
	"github.com/r57ty7/pver/service"
	"github.com/spf13/cobra"
)

func newJiraCmd() *cobra.Command {
	var jiraCmd = &cobra.Command{
		Use:   "jira",
		Short: "JIRA related operation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	jiraCmd.AddCommand(newJiraSearchCmd())
	jiraCmd.AddCommand(newJiraOpenCmd())
	jiraCmd.AddCommand(newJiraBranchCmd())

	return jiraCmd
}

func newJiraSearchCmd() *cobra.Command {
	var jiraCmd = &cobra.Command{
		Use:   "search [jql]",
		Short: "Search JIRA ticket",
		RunE: func(cmd *cobra.Command, args []string) error {
			issues, err := searchTicket(args)
			if err != nil {
				cmd.PrintErrf("%v\n", err)
				return err
			}

			for _, v := range issues {
				cmd.Printf("%v\n", v.Key)
			}

			return nil
		},
	}
	return jiraCmd
}

func newJiraOpenCmd() *cobra.Command {
	var jiraCmd = &cobra.Command{
		Use:   "open [jql]",
		Short: "Open browser with JIRA ticket",
		RunE: func(cmd *cobra.Command, args []string) error {
			issues, err := searchTicket(args)
			if err != nil {
				cmd.PrintErrf("%v\n", err)
				return err
			}

			issue, err := selectTicket(issues)
			if err != nil {
				cmd.PrintErrf("%v\n", err)
				return err
			}

			return openTicket(issue)
		},
	}
	return jiraCmd
}

func newJiraBranchCmd() *cobra.Command {
	var jiraCmd = &cobra.Command{
		Use:   "branch [jql]",
		Short: "Create branch from JIRA ticket",
		RunE: func(cmd *cobra.Command, args []string) error {
			issues, err := searchTicket(args)
			if err != nil {
				cmd.PrintErrf("%v\n", err)
				return err
			}

			issue, err := selectTicket(issues)
			if err != nil {
				cmd.PrintErrf("%v\n", err)
				return err
			}

			suffix := inputBranchSuffix()
			err = gitRepository.CreateBranch(fmt.Sprintf("feature/%s_%s", issue.Key, suffix))
			if err != nil {
				cmd.PrintErrf("%v\n", err)
				return err
			}

			message := emoji.Sprintf(":check_mark: Created branch")
			cmd.Println(message)

			return nil
		},
	}
	return jiraCmd
}

func searchTicket(args []string) ([]service.Issue, error) {
	jql := conf.Jira.JQL
	if len(args) > 0 {
		jql = args[0]
	}

	return jiraService.Search(context.Background(), jql)

}

func selectTicket(issues []service.Issue) (service.Issue, error) {
	// select ticket
	idx, err := fuzzyfinder.Find(
		issues,
		func(i int) string {
			return issues[i].Key
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("ID: %s\nSummary: %s",
				issues[i].Key,
				issues[i].Fields.Summary,
			)
		}))

	if err != nil {
		return service.Issue{}, err
	}

	return issues[idx], nil
}

func openTicket(issue service.Issue) error {
	url := conf.Jira.BaseURL + "/browse/" + issue.Key
	return exec.Command("open", url).Start()
}

func inputBranchSuffix() string {
	prompt := promptui.Prompt{
		Label: "Branch suffix: ",
	}
	result, _ := prompt.Run()
	return result
}
