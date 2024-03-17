package cmd

import (
	"fmt"
	"os"

	"github.com/puffyguy/go-web-scrapper/scrapper"
	"github.com/spf13/cobra"
)

var url string
var classes []string

var rootCmd = &cobra.Command{
	Use:   "go-web-scrapper",
	Short: "Web Scrapper CLI app with Cobra",
	Long:  `This app will scrape the URL provided by YOU for YOU.`,
}

var selectorCmd = &cobra.Command{
	Use:   "selector",
	Short: "CSS Selector",
	Long:  `Provide custom css selector`,
	Run: func(cmd *cobra.Command, args []string) {
		classes, _ = cmd.Flags().GetStringSlice("class")
		// fmt.Printf("URL: %s, Class: %s\n", url, classes)

		scrapper.ScrapeTheURL(url, classes)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "URL to scrape")
	rootCmd.MarkPersistentFlagRequired("url")

	selectorCmd.Flags().StringSliceP("class", "c", []string{}, "CSS selectors to be used")
	selectorCmd.MarkFlagRequired("class")

	rootCmd.AddCommand(selectorCmd)
}
