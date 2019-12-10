package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Long:  "CLI utility for work with distribute data storage",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
