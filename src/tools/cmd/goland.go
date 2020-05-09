package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"microservice/src/tools/helper"
)

var golandCmd = &cobra.Command{
	Use:   "goland",
	Short: "scanner file",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := helper.ScannerService(args)
		if err != nil {
			fmt.Println(fmt.Sprintf("Scan dir error %v", err))
		}

		//err = BuildBuddyYml(services)
		//
		//if err != nil {
		//	fmt.Println(fmt.Sprintf("Error build yml file %v", err))
		//}
	},
}

func init() {
	rootCmd.AddCommand(golandCmd)
}
