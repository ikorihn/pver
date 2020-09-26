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

// npmCmd represents the npm command
func newNpmCmd() *cobra.Command {
	var updateVer string

	var npmCmd = &cobra.Command{
		Use:   "npm [-u version]",
		Short: "Show, update package.json's version",
		RunE: func(cmd *cobra.Command, args []string) error {
			if conf.Npm.Filepath == "" {
				conf.Npm.Filepath = "package.json"
			}

			npmFvm.SetConfig(conf)
			version := npmFvm.Version()
			cmd.Printf("Version: %v\n", version)

			if updateVer != "" {
				err := npmFvm.Update(updateVer)
				if err != nil {
					cmd.PrintErrf("update error: %v", err)
					return err
				}
				cmd.Printf("Updated to => %v\n", updateVer)

				err = gitRepository.CommitUpdate(conf.Npm.Filepath, updateVer)
				if err != nil {
					cmd.PrintErrf("commit error: %v", err)
					return err
				}
			}
			return nil
		},
	}

	npmCmd.Flags().StringVarP(&updateVer, "update", "u", "", "update to specified version")

	return npmCmd
}
