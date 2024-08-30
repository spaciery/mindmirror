package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const VERSION = "v0.1"
const VERSION_DESC = "Experimental"
const APP_INFO = `
   •   ┓   •
┏┳┓┓┏┓┏┫┏┳┓┓┏┓┏┓┏┓┏┓
┛┗┗┗┛┗┗┻┛┗┗┗┛ ┛ ┗┛┛

mindmirror dont know what to do ..
use mindmirror --help for more information on how to use mindmirror
	`

var rootCmd = &cobra.Command{
	Use:   "mindmirror",
	Short: "mindmirror is a minimal static site generator intented to use as a documentation/notes viewer",
	Long: `
Mindmirror is a static site generator that converts markdown files to webpages,
it can generate pages from diffrent sources (local folder or a git repository).
once generated it can also deploy the pages on a desired host
	`,
	Run: runcommand,
}

func runcommand(cmd *cobra.Command, args []string) {
	showVer, _ := cmd.Flags().GetBool("version")
	if showVer {
		fmt.Printf("Mindmirror - %s\nyou are using an %s version\n", VERSION, VERSION_DESC)
	} else {
		fmt.Println(APP_INFO)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show current version")
}
