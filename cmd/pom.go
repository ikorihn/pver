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
	"github.com/spf13/cobra"
)

// pomCmd represents the pom command
func newPomCmd() *cobra.Command {
	var updateVer string

	var pomCmd = &cobra.Command{
		Use:   "pom [pom.xml]",
		Short: "Show, update pom.xml's version",
		RunE: func(cmd *cobra.Command, args []string) error {
			if conf.Pom.Filepath == "" {
				conf.Pom.Filepath = "pom.xml"
			}
			if conf.Pom.Indent == "" {
				conf.Pom.Indent = "  "
			}

			pomFvm.SetConfig(conf)
			version := pomFvm.Version()
			cmd.Printf("Version: %v\n", version)

			if updateVer != "" {
				err := pomFvm.Update(updateVer)
				if err != nil {
					cmd.PrintErrf("update error: %v", err)
					return err
				}
				cmd.Printf("Updated to => %v\n", updateVer)

				gitRepository.CommitUpdate(conf.Pom.Filepath, updateVer)
			}
			return nil
		},
	}

	pomCmd.Flags().StringVarP(&updateVer, "update", "u", "", "update to specified version")

	return pomCmd
}
