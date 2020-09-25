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

	return jiraCmd
}

func newJiraSearchCmd() *cobra.Command {
	var jiraCmd = &cobra.Command{
		Use:   "search",
		Short: "Search JIRA ticket",
		RunE: func(cmd *cobra.Command, args []string) error {
			issues, err := jiraService.Search(context.Background(), "")
			if err != nil {
				cmd.PrintErrf("%v\n", err)
				return err
			}

			for _, v := range issues {
				fmt.Printf("%v\n", v)
			}

			return nil
		},
	}
	return jiraCmd
}
