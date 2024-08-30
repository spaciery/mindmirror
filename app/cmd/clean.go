package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"samin.dev/mindmirror/cleaner"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean up the files generated during build",
	Long: `mindmirror build command builds the html pages into a provided/default directory
this directory has to be cleaned before doing a refresh. or simply use this command to delete
files generated by mindmirror`,
	Run: runClean,
}

func runClean(cmd *cobra.Command, args []string) {

	clnr := &cleaner.Cleaner{
		Directory: viper.GetString("app.content.dest"),
	}

	err := clnr.Clean()
	if err != nil {
		if errors.Is(err, cleaner.NothingToClean) {
			fmt.Println("Nothing to clean")
		} else {
			fmt.Println("cleaning failed", err)
		}
	}
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}