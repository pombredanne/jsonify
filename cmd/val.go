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
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// valCmd represents the val command
var valCmd = &cobra.Command{
	Use:   "val",
	Short: "Return keys associated with val",
	Long:  `Return keys associated with given value`,
	Run:   getVal,
}

func getVal(cmd *cobra.Command, args []string) {

	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var data interface{}

	err = json.Unmarshal([]byte(raw), &data)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

    fmt.Println("VAL MODULE")

	m := data.(map[string]interface{})

	for k, v := range m {
        switch val := v.(type) {
            case string:
				if InSlice(v, args) {
					fmt.Println(k)
				}
            case int:
				if InSlice(v, args) {
					fmt.Println(k)
				}
            case []interface{}:
                for i, u := range v {
					if InSlice(i, args) {
						fmt.Println(i)
					}
            }
            default:
                fmt.Println(val, "is of a type I don't know how to handle")
            }
	}
	os.Exit(0)

}

func init() {
	RootCmd.AddCommand(valCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// valCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// valCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
