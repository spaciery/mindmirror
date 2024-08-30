package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Port:       viper.GetString("app.port"),
		PageSrcDir: viper.GetString("app.content.dest"),
		StylesDir:  viper.GetString("app.styles.path"),
		ScriptsDir: viper.GetString("app.scripts.path"),
	}
	app.Serve()
}

func init() {
	runCmd.Flags().StringP("port", "p", "8080", "port to run mindmirror on")
	viper.BindPFlag("app.port", runCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(runCmd)
}
