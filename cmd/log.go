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
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

/*type LogMessage struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
} */

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
		fmt.Println("log called")
		message, _ := cmd.Flags().GetString("message")
		if message == "" {
			fmt.Println("Error. Cannot write an empty message.")
		} else {
			now := time.Now()
			formattedtime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())

			f, err := os.OpenFile("audit.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			w := csv.NewWriter(f)
			w.Write([]string{formattedtime, message})
			w.Flush()

			/*if _, err := os.Stat("audti.csv"); err == nil {
				// file exists append message to it
				f, err := os.OpenFile("audit.csv", os.O_WRONLY|os.O_APPEND, 0644)
				w := csv.NewWriter(f)
				w.Write()
				w.Flush()

			} else if os.IsNotExist(err) {
				// file does not exist
				//create file
				// append message
			} else {
				// WTF how did we get here
			}*/

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
	logCmd.Flags().StringP("message", "m", "", "Add a message to log to your notebook")
}