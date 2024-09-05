package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type AppServer struct {
	Port       string
	PageSrcDir string
	StylesDir  string
	ScriptsDir string
}

var pageSourceDir string

func (app *AppServer) Serve() {

	pageSourceDir = app.PageSrcDir

	http.HandleFunc("/", handleRequest)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir(app.StylesDir))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir(app.ScriptsDir))))

	log.Printf("Serving site on http://localhost:%s\n", app.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", app.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		serveIndex(w, r)
		return
	}

	filePath := filepath.Join(pageSourceDir, r.URL.Path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		filePath = filePath + ".html"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	http.ServeFile(w, r, filePath)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {

	files, err := getHTMLFiles(pageSourceDir)
	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}

	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MindMirror</title>
   <link rel="stylesheet" href="/styles/water-auto.css">
   <link rel="stylesheet" href="/styles/imhome.css">
</head>
<body>
    <div id='nav'>
    <h1>MindMirror_</h1>
    <h2>↙ Index</h1>
    </div>
    <div id="body-cont">
    <div id="image">
    <img id="randomImage" alt="Random Image" style="max-width: 100%; height: auto;">
    </div>
    <div id="file-tree">
        {{.}}
        </div>
    </div>
    <script src="/scripts/tree.js"></script>
    <script src="/scripts/image.js"></script>
</body>
</html>
`

	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = t.Execute(w, template.HTML(buildTree(files)))
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func getHTMLFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			relPath, _ := filepath.Rel(dir, path)
			files = append(files, relPath)
		}
		return nil
	})
	sort.Strings(files)
	return files, err
}

func buildTree(files []string) string {

	var tree strings.Builder
	tree.WriteString("<ul>")

	var currentPath []string
	for _, file := range files {
		parts := strings.Split(file, string(filepath.Separator))

		i := 0
		for i < len(currentPath) && i < len(parts) && currentPath[i] == parts[i] {
			i++
		}

		for j := len(currentPath); j > i; j-- {
			tree.WriteString("</ul></li>")
		}

		for j := i; j < len(parts)-1; j++ {
			tree.WriteString("<li><span class='folder'>" + parts[j] + "</span><ul class='nested-item'>")
		}

		fileName := parts[len(parts)-1]
		displayName := strings.TrimSuffix(fileName, ".html")
		tree.WriteString("<li class='html-file'><a href='" + file + "'>" + displayName + "</a></li>")

		currentPath = parts[:len(parts)-1]
	}

	for range currentPath {
		tree.WriteString("</ul></li>")
	}

	tree.WriteString("</ul>")
	return tree.String()
}
