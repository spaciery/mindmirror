package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var cfgFile string
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
	cobra.OnInitialize(initConfigs)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file to be used (default : $XDG_CONFIG_DIR/mindmirror/config.toml)")
	rootCmd.Flags().BoolP("version", "v", false, "Show current version")
}

func initConfigs() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		configpath, err := os.UserConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(configpath + "/mindmirror/")
		viper.AddConfigPath("/etc/mindmirror/")
		viper.AddConfigPath(".")
		viper.SetConfigName("mindmirror")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found,exiting")
		} else {
			fmt.Println("error loading config", err)
		}
		os.Exit(1)
	}

}
