package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use: "version",
  Short: "Print the job service version",
  Run: func(cmd *cobra.Command, args[]string) {
    fmt.Println("job service v0.1")
  },
}
