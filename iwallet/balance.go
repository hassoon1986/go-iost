// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package iwallet

import (
	"fmt"

	"github.com/spf13/cobra"
)

// accountInfoCmd represents the balance command.
var accountInfoCmd = &cobra.Command{
	Use:   "balance",
	Short: "Check balance information of a specified account",
	Long:  `Check balance information of a specified account, must followed by an account ID`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Please enter the account ID")
		}
		// Do some checks for arg[0] here.
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		id := args[0]
		info, err := sdk.getAccountInfo(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(marshalTextString(info))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountInfoCmd)
}
