// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "sets up the expected folder structure for the tool to wrk.",
	Long: ` wrk setup creates a file structure like this 
				- $PWD/nb/names/hello/hello.cs`,
	Run: func(cmd *cobra.Command, args []string) {
		// Getting the paths to be used
		rootPath, err := os.Getwd()
		fmt.Println("rootPath ", rootPath)
		if err != nil {
			// fmt.Printf("[ERR] Getting path %v", err)
		}
		path := filepath.Join(rootPath, "nb")
		configPath := filepath.Join(rootPath, ".config")
		// make the default  nb dir and nb hello
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("[ERR] Mkdir %v", err)
		}
		// make the default config file
		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			fmt.Printf("[ERR] Mkdir %v", err)
		}

		// now create the default yaml file
		os.Chdir(".config")
		str := fmt.Sprintf("basePath: %s\ncurrentNotebook: %s", rootPath, path)
		data := []byte(str)
		err = ioutil.WriteFile("config.yaml", data, 0644)
		if err != nil {
			fmt.Printf("[ERR] %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
