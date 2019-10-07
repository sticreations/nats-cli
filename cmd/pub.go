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
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publishes a Message to a Topic",
	Long:  `To Publish a Message to a Topic use nats-cli pub TOPIC MESSAGE.`,
	Run: func(cmd *cobra.Command, args []string) {

		c, err := nats.Connect(natsServer)
		if err != nil {
			fmt.Errorf("Could not Connect to Nats Server: %v", err)
		}
		if len(args) == 0 {
			fmt.Errorf("Please add a Topic to Publish to")
		}
		topic := args[0]
		if len(args) > 1 {
			message := strings.Join(args[1:], " ")
			c.Request(topic, []byte(message), nats.DefaultTimeout)
			fmt.Printf("Published %v on Topic %v\n", message, topic)
		} else {
			c.Request(topic, []byte(nil), nats.DefaultTimeout)
			fmt.Printf("Published on Topic %v\n", topic)
		}
		c.Close()
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
