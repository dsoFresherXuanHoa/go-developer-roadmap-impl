/*
dsoFresherXuanHoa <dso.intern.xuanhoa@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webcrawl",
	Short: "Simple Crawling Tool using Golang",
	Long:  `Simple Crawling Tool using Golang`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
