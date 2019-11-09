// Copyright © 2019 Karl Hepworth <Karl.Hepworth@gmail.com>
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
	"github.com/fubarhouse/pygmy/v1/service/library"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Example: "pygmy status",
	Short: "Report status of the pygmy services",
	Long: `Loop through all of pygmy's services and identify the present state.
This includes the docker services, the resolver and SSH key status`,
	Run: func(cmd *cobra.Command, args []string) {

		library.Status(c)

	},
}

func init() {

	rootCmd.AddCommand(statusCmd)

}
