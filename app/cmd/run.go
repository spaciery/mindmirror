package cmd

import (
	"github.com/spf13/cobra"
	"samin.dev/mindmirror/server"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the sever",
	Long: `unless provided application will run on port 8080 with the pre-build content
content have to be build before running the server`,
	Run: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	app := &server.AppServer{
		Port:       "8080",
		PageSrcDir: "../public",
		StylesDir:  "../styles",
		ScriptsDir: "../scripts",
	}
	app.Serve()
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.Flags().StringP("port", "p", "8080", "port to run mindmirror on")
}
