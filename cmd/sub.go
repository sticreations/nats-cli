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
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := nats.Connect(natsServer)
		if err != nil {
			fmt.Errorf("Could not Connect to Nats Server: %v", err)
		}
		if len(args) < 1 {
			fmt.Errorf("Please define a Topic to Subscribe")
		}
		wg := sync.WaitGroup{}
		wg.Add(1)
		fmt.Printf("Subscribing on Topic %v\n", args[0])
		c.Subscribe(args[0], func(msg *nats.Msg) {
			fmt.Printf("Message on Topic %v: %v\n", msg.Subject, string(msg.Data))
		})
		// interrupt handling
		wg.Wait()
		c.Close()

	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
