package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

const (
	DEFAULT_CONTENT_DIRECTORY = "./content"
	OUT_DIRECTORY             = "./public"
	DEFAULT_CSS_FILE          = "./styles/default.css"
	SCRIPT_HTML               = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
    		<meta charset="UTF-8">
      		<meta name="viewport" content="width=device-width, initial-scale=1.0">
        	<title>%s</title>
         	<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
          	<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
           	<link rel="stylesheet" href="/styles/default.css">
            <link rel="stylesheet" href="/styles/prism.css">
        </head>
        <body>
        %s
        <script src="/scripts/prism.js"></script>
        </body>
        </html>
`
)

var (
	port             = 8080
	md               goldmark.Markdown
	contentDirectory string
	cssFile          string
)

func main() {

	generate := flag.Bool("g", false, "Generate the static site")
	serve := flag.Bool("s", false, "Serve the static site")
	clean := flag.Bool("c", false, "Clean previous out")
	flag.StringVar(&contentDirectory, "d", DEFAULT_CONTENT_DIRECTORY, "Specify the content directory")
	flag.StringVar(&cssFile, "css", DEFAULT_CSS_FILE, "Specify the CSS file to use")
	flag.Parse()

	if *clean {
		err := cleanPublic()
		if err != nil {
			log.Fatal("cleaning failed , exiting .. ")
		}
	}

	if *generate {
		generateSite()
	}

	if *serve {
		serveSite()
	}

	if !*generate && !*serve {
		fmt.Println("Please specify at least one action: -g to generate, -s to serve")
	}
}

func cleanPublic() error {

	if _, err := os.Stat(OUT_DIRECTORY); os.IsNotExist(err) {
		log.Printf("Directory %s does not exist, nothing to clean.\n", OUT_DIRECTORY)
		return nil
	}

	err := os.RemoveAll(OUT_DIRECTORY)
	if err != nil {
		log.Printf("Error removing directory %s: %v\n", OUT_DIRECTORY, err)
		return err
	}

	log.Printf("Successfully removed directory %s\n", OUT_DIRECTORY)

	return nil
}

func generateSite() {

	err := copyCSSFile()
	if err != nil {
		log.Printf("Error copying CSS file: %v\n", err)
		return
	}

	initializeGoldmark()

	os.MkdirAll(OUT_DIRECTORY, os.ModePerm)

	err = filepath.Walk(contentDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		return processMarkdownFile(path)
	})

	if err != nil {
		log.Printf("Error generating site: %v\n", err)
	} else {
		log.Println("Site generated successfully")
	}
}

func copyCSSFile() error {

	err := os.MkdirAll(filepath.Join(OUT_DIRECTORY, "styles"), os.ModePerm)
	if err != nil {
		return err
	}

	source, err := os.Open(cssFile)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(filepath.Join(OUT_DIRECTORY, "styles", "default.css"))
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func initializeGoldmark() {

	md = goldmark.New(
		goldmark.WithExtensions(mathjax.MathJax),
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
		goldmark.WithRendererOptions(
			renderer.WithNodeRenderers(
				util.Prioritized(NewLinkRenderer(), 1),
			),
		),
	)
}

func NewLinkRenderer() renderer.NodeRenderer {
	return &LinkRenderer{}
}

type LinkRenderer struct{}

func (r *LinkRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindLink, r.renderLink)
}

func (r *LinkRenderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Link)
	dest := string(n.Destination)
	if strings.HasSuffix(dest, ".md") {
		dest = strings.TrimSuffix(dest, ".md") + ".html"
	}
	_, _ = w.WriteString("<a href=\"")
	_, _ = w.WriteString(dest)
	_, _ = w.WriteString("\">")
	return ast.WalkContinue, nil
}

func processMarkdownFile(path string) error {

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		return err
	}
	html := buf.Bytes()
	relPath, _ := filepath.Rel(contentDirectory, path)
	outPath := filepath.Join(OUT_DIRECTORY, strings.TrimSuffix(relPath, ".md")+".html")
	os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	fullHTML := fmt.Sprintf(SCRIPT_HTML, strings.TrimSuffix(filepath.Base(path), ".md"), string(html))
	err = os.WriteFile(outPath, []byte(fullHTML), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Generated: %s\n", outPath)
	return nil
}

func serveSite() {
	http.HandleFunc("/", handleRequest)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("./scripts"))))

	log.Printf("Serving site on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		serveIndex(w, r)
		return
	}

	filePath := filepath.Join(OUT_DIRECTORY, r.URL.Path)
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

	files, err := getHTMLFiles(OUT_DIRECTORY)
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
    <link rel="stylesheet" href="/styles/server.css">
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
