package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var printFlag string
var printCmd = &cobra.Command{
	Use: "print [OPTIONS] [COMMANDS]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run print...")
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
