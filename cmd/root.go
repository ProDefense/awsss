package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "awsss",
    Short: "AWS Security Scanner",
    Long:  "A simple AWS Security Scanner (awsss) tool to analyze AWS IAM roles.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
