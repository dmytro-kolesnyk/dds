package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Files list from storage",
	Long:  `Files list from storage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Files LIST")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
