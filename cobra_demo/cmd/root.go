package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Verbose bool
var Source string
var Region string
var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run hugo ...")
		fmt.Printf("Verbose: %v\n", Verbose)
		fmt.Printf("Source: %v\n", Source)
		fmt.Printf("Region: %v\n", Region)
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
