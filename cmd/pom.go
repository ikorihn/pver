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
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
)

// pomCmd represents the pom command
func newPomCmd(fvm FileVersionManager) *cobra.Command {
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

			fvm.SetConfig(conf)
			version := fvm.Version()
			cmd.Printf("Version: %v\n", version)

			if updateVer != "" {
				err := fvm.Update(updateVer)
				if err != nil {
					cmd.PrintErrf("update error: %v", err)
					return err
				}
				cmd.Printf("Updated to => %v\n", updateVer)

				wd := "./"
				repo, err := git.PlainOpen(wd)
				if err != nil {
					if errors.Is(err, git.ErrRepositoryNotExists) {
						// make error message git-like
						fmt.Println("fatal: not a git repository")
						os.Exit(1)
					}
					return nil
				}

				w, err := repo.Worktree()
				if err != nil {
					return err
				}
				_, err = w.Add(conf.Pom.Filepath)
				if err != nil {
					return err
				}

				email, err := gitconfig.Email()
				if err != nil {
					return err
				}
				name, err := gitconfig.Username()
				if err != nil {
					return err
				}

				fmt.Printf("commit: version up to %s\n", updateVer)
				w.Commit(fmt.Sprintf("version up to %s", updateVer), &git.CommitOptions{
					Author: &object.Signature{
						Email: email,
						Name:  name,
						When:  time.Now(),
					},
				})
				return nil
			}
			return nil
		},
	}

	pomCmd.Flags().StringVarP(&updateVer, "update", "u", "", "update to specified version")

	return pomCmd
}
