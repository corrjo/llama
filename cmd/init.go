/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a serverless project",
	Long: `Use llama init to start a project. llama will generate the basic file
        structure and some basic code to get you started.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateProject(args)
	},
}

func CreateProject(args []string) {
	current_path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Creating project %s\n", args[0])
	os.Mkdir(current_path+"/"+args[0], 0755)
	subDirectories := []string{"src", "deployment"}
	for _, dir := range subDirectories {
		os.Mkdir(current_path+"/"+args[0]+"/"+dir, 0755)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("lang", "l", "go", "Language of the serverless function")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
