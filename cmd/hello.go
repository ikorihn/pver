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

	"github.com/spf13/cobra"
)

type helloOption struct {
	age      int
	lastName string
}

var helloOpt = helloOption{}

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Long:  `Longer desription for hello command. Say hello to whom specified in arg.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Name is required")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Hello, %s\n", args[0])
		if helloOpt.lastName != "" {
			fmt.Printf("Name: %s %s\n", args[0], helloOpt.lastName)
		}
		if helloOpt.age >= 0 {
			fmt.Printf("Age: %v\n", helloOpt.age)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	helloCmd.Flags().IntVarP(&helloOpt.age, "age", "a", -1, "Age")
	helloCmd.Flags().StringVarP(&helloOpt.lastName, "last-name", "l", "", "Last name")
}
