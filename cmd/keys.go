// Copyright © 2016 Brett Smith <bc.smith@sas.com>
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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Return all top level keys",
	Long:  `Return all the top level keys`,
	Run:   getKeys,
}

func getKeys(cmd *cobra.Command, args []string) {

	raw, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var data interface{}

	err = json.Unmarshal(raw, &data)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

    fmt.Println("KEYS MODULE")

	m := data.(map[string]interface{})

	for k, _ := range m {
		switch kk := k.(type) {
			case string:
				fmt.Println(k)
			case int:
				fmt.Println(k)
			case []interface{}:
				for i, u := range k {
					fmt.Println(i)
			}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
	}

	os.Exit(0)
}

func init() {
	RootCmd.AddCommand(keysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
