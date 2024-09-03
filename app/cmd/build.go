package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"samin.dev/mindmirror/builder"
)

const SCRIPT_HTML = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
   		<meta charset="UTF-8">
     	<meta name="viewport" content="width=device-width, initial-scale=1.0">
       	<title>%s</title>
        <link rel="icon" href="/styles/favicon.ico" type="image/x-icon">
        <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
        <link rel="stylesheet" href="/styles/highlight/styles/base16/horizon-dark.css">
        <script src="/styles/highlight/highlight.min.js"></script>
        <script src="/styles/highlight/languages/java.min.js"></script>
        <link href="https://fonts.googleapis.com/css2?family=Source+Sans+Pro:wght@400;600;700&family=Source+Code+Pro&display=swap" rel="stylesheet">
        <link rel="stylesheet" href="{{THEME}}">
       </head>
       <body>
       %s
        <script>hljs.highlightAll();</script>
       </body>
       </html>
`

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build html pages from the markdown content",
	Long: `Using build html files are generated from markdown ,
the source and destination can be provided`,
	Run: runBuild,
}

func runBuild(cmd *cobra.Command, args []string) {

	themedHTML := strings.Replace(SCRIPT_HTML, "{{THEME}}", "/styles/"+viper.GetString("app.styles.stylesheets.pages"), 1)

	b := builder.Builder{
		TempDir:    viper.GetString("app.content.dest"),
		SourceDir:  viper.GetString("app.content.source"),
		HeaderHtml: themedHTML,
		StyleSheet: viper.GetString("app.styles.path") + "/" + viper.GetString("app.styles.stylesheets.pages"),
	}

	b.GenerateSite()
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
