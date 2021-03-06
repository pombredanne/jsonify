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

// valsCmd represents the vals command
var valsCmd = &cobra.Command{
	Use:   "vals",
	Short: "Return all the top level values",
	Long:  `Return all the top level values`,
	Run:   getVals,
}

func getVals(cmd *cobra.Command, args []string) {

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

	m := data.(map[string]interface{})

	for _, v := range m {
		switch val := v.(type) {
            case string:
			    fmt.Println(v)
		    case int:
			    fmt.Println(v)
		    case []interface{}:
			    for i, u := range v {
				    fmt.Println(u)
			    }
		    default:
			    fmt.Println(v, "is of a type I don't know how to handle")
		}
	}
	os.Exit(0)
}

func init() {
	RootCmd.AddCommand(valsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// valsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// valsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
