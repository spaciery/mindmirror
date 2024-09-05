package cmd

import (
	"context"
	"log"
	"os"
	"strings"
	"sync"

	"fyne.io/systray"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"samin.dev/mindmirror/builder"
	"samin.dev/mindmirror/server"
)

var trayCmd = &cobra.Command{
	Use:   "tray",
	Short: "open mindmirror in system tray",
	Long:  `system tray can be used to access clean , build and run in a convinient way`,
	Run:   trayCommand,
}

func trayCommand(cmd *cobra.Command, args []string) {
	systray.Run(onReady, onExit)
}

var (
	app             *server.AppServer
	serverCtx       context.Context
	serverCancel    context.CancelFunc
	serverWg        sync.WaitGroup
	isServerRunning bool
	iconStop        []byte
	iconRunning     []byte
)

func loadIcon(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Failed to load icon: %v", err)
		return nil
	}
	return b
}

func onReady() {

	iconStop = loadIcon(viper.GetString("app.styles.path") + "/rec.png")
	iconRunning = loadIcon(viper.GetString("app.styles.path") + "/lotus.png")

	app = &server.AppServer{
		Port:       viper.GetString("app.port"),
		PageSrcDir: viper.GetString("app.content.dest"),
		StylesDir:  viper.GetString("app.styles.path"),
		ScriptsDir: viper.GetString("app.scripts.path"),
	}
	themedHTML := strings.Replace(SCRIPT_HTML, "{{THEME}}", "/styles/"+viper.GetString("app.styles.stylesheets.pages"), 1)
	b := builder.Builder{
		TempDir:    viper.GetString("app.content.dest"),
		SourceDir:  viper.GetString("app.content.source"),
		HeaderHtml: themedHTML,
		StyleSheet: viper.GetString("app.styles.path") + "/" + viper.GetString("app.styles.stylesheets.pages"),
	}

	systray.SetTitle("Mindmirror")
	systray.SetIcon(iconStop)
	systray.SetTooltip("mindmirror system tray utility")

	mBuild := systray.AddMenuItem("Build", "Build the application")
	mRun := systray.AddMenuItem("Run", "Run the application")
	mStop := systray.AddMenuItem("Stop", "Stop the server")
	mStop.Disable()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	go func() {
		for {
			select {
			case <-mBuild.ClickedCh:
				b.GenerateSite()
			case <-mRun.ClickedCh:
				if !isServerRunning {
					systray.SetIcon(iconRunning)
					startServer()
					mRun.Disable()
					mStop.Enable()
				}
			case <-mStop.ClickedCh:
				if isServerRunning {
					systray.SetIcon(iconStop)
					stopServer()
					mStop.Disable()
					mRun.Enable()
				}
			case <-mQuit.ClickedCh:
				if isServerRunning {
					stopServer()
				}
				systray.Quit()
				return
			}
		}
	}()
}

func startServer() {
	serverCtx, serverCancel = context.WithCancel(context.Background())
	serverWg.Add(1)
	go func() {
		defer serverWg.Done()
		app.ServeWithContext(serverCtx)
	}()
	isServerRunning = true
}

func stopServer() {
	if serverCancel != nil {
		serverCancel()
		serverWg.Wait()
		isServerRunning = false
	}
}

func onExit() {
	if isServerRunning {
		stopServer()
	}
}

func init() {
	trayCmd.Flags().StringP("port", "p", "8080", "port to run mindmirror on")
	viper.BindPFlag("app.port", trayCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(trayCmd)
}
