/*
Copyright © 2023 Syro team <info@syro.com>
*/

package cmd

import (
	"fmt"
	"syro/api"
	"syro/util"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run your app with the Syro secrets injected",
	Long:  "Fetch the latest secrets from your chosen project and inject it to your app via a .env file.",
	Run: func(cmd *cobra.Command, args []string) {
		isConfigLoaded, config, err := util.LoadConfigFromProjectConfigFile()
		if err != nil {
			fmt.Println("Something went wrong while loading items from your project config file.")
		}
		if isConfigLoaded == true {
			projectId := ""
			if len(config.ProjectId) == 0 {
				projectId, err = getProjectIdAndUpdateProjectConfigFile(config.CompanyId, config.MemberId, config.SessionToken)
				if err != nil {
					return
				}
			} else {
				projectId = config.ProjectId
			}
			items, err := api.FetchProjectItems(config.CompanyId, config.MemberId, projectId, config.SessionToken)
			if err != nil {
				fmt.Println("Something went wrong while fetching project secrets.")
				return
			}

			err = util.SaveSecretsToEnvFile(items)
			if err != nil {
				fmt.Println("Something went wrong while saving project secrets to .env file.")
				return
			}

		} else {
			fmt.Println("The Syro CLI is not properly configured yet for this project. Kindly complete the set up first by using the login command.")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
