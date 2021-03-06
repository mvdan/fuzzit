/*
Copyright © 2019 fuzzit.dev, inc.

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
	"log"
)

// targetCmd represents the target command
var targetCmd = &cobra.Command{
	Use:   "target [target_name]",
	Short: "Create new fuzzing target",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating target...")

		seed, err := cmd.Flags().GetString("seed")
		if err != nil {
			log.Fatal(err)
		}
		_, err = gFuzzitClient.CreateTarget(args[0], seed)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created new target: %s successfully", args[0])
	},
}

func init() {
	createCmd.AddCommand(targetCmd)

	targetCmd.Flags().String("seed", "", "path to .tar.gz seed corpus")
}
