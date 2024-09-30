package cmd

import (
	"awsss/pkg/graph"
	"fmt"

	"github.com/spf13/cobra"
)

var fileType string

var graphCmd = &cobra.Command{
    Use:   "graph",
    Short: "Generate a trust relationship graph",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Please provide an output file name (without extension).")
            return
        }
        outputFile := args[0]
        err := graph.GenerateTrustGraph(outputFile, fileType)
        if err != nil {
            fmt.Println("Error generating graph:", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(graphCmd)
    graphCmd.Flags().StringVarP(&fileType, "type", "t", "svg", "Specify output file type (svg or png)")
}
