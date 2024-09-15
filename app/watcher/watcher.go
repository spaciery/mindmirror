package watcher

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"samin.dev/mindmirror/builder"
)

type Watcher struct {
	Source string
}

func (w *Watcher) WatchAndBuild() error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if strings.HasSuffix(filepath.Base(event.Name), "~") {
					log.Println("Ignoring tempfile")
					continue
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file:", event.Name)

					b := builder.Builder{
						TempDir:    viper.GetString("app.content.dest"),
						SourceDir:  event.Name,
						StyleSheet: viper.GetString("app.styles.path") + "/" + viper.GetString("app.styles.stylesheets.pages"),
					}

					b.GenerateSite()

				} else if event.Op&fsnotify.Create == fsnotify.Create {

					log.Println("Created file or directory:", event.Name)
					fi, err := os.Stat(event.Name)
					if err == nil && fi.IsDir() {
						log.Println("Adding new directory to watcher:", event.Name)
						err = watcher.Add(event.Name)
						if err != nil {
							log.Println("Error adding directory to watcher:", err)
						}
					}

				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("Removed file:", event.Name)
					// Handle file removal
					//
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	err = filepath.Walk(w.Source, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	log.Printf("Watching directory: %s\n", w.Source)
	<-done
	return nil
}
