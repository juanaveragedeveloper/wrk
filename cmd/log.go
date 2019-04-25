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
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

type conf struct {
	Basepath        string `yaml:"basePath"`
	Currentnotebook string `yaml:"currentNotebook"`
}

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log your message to a notebook",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")
		tag, _ := cmd.Flags().GetString("tag")
		all, _ := cmd.Flags().GetString("all")
		find, _ := cmd.Flags().GetString("find")

		var configuration conf
		workingdirectory, _ := os.Getwd()
		configFile := workingdirectory + "\\.config\\config.yml"
		yamlFile, err := ioutil.ReadFile(configFile)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = yaml.Unmarshal(yamlFile, &configuration)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(configuration.Currentnotebook)
		path := strings.Split(configuration.Currentnotebook, "\\")
		length := len(path)
		notebook := "\\" + path[length-1] + ".csv"

		if message != "" {
			if all != "" || find != "" {
				fmt.Println("Error cannot use find or all when logging a message")
				return
			}

			now := time.Now()
			formattedtime := fmt.Sprintf("%d/%02d/%02dT%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())

			f, err := os.OpenFile(configuration.Currentnotebook+notebook, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			w := csv.NewWriter(f)
			w.Write([]string{formattedtime, message, tag})
			w.Flush()

			fmt.Println("Successfully logged message to notebook")
			return
		}

		if find != "" {
			f, _ := os.Open(configuration.Currentnotebook + notebook)

			r := csv.NewReader(bufio.NewReader(f))

			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				message := record[1]
				recordedtag := record[2]
				if strings.Contains(message, find) {
					if tag != "" && recordedtag == tag {
						fmt.Println(record)
					} else {
						fmt.Println(record)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	logCmd.Flags().StringP("tag", "t", "", "tag your message")
	logCmd.Flags().StringP("message", "m", "", "Add a message to log to your notebook")
	logCmd.Flags().StringP("all", "a", "", "Print all logs for a notebook")
	logCmd.Flags().StringP("find", "f", "", "Find logs")
}
