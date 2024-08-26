package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

const (
	CONTENT_DIRECTORY = "./content"
	OUT_DIRECTORY     = "./public"
	SCRIPT_HTML       = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
    <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
</head>
<body>
    %s
</body>
</html>
`
)

var (
	port = 8080
	md   goldmark.Markdown
)

func main() {

	generate := flag.Bool("g", false, "Generate the static site")
	serve := flag.Bool("s", false, "Serve the static site")
	flag.Parse()

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

func generateSite() {

	initializeGoldmark()

	os.MkdirAll(OUT_DIRECTORY, os.ModePerm)

	err := filepath.Walk(CONTENT_DIRECTORY, func(path string, info os.FileInfo, err error) error {
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

func initializeGoldmark() {

	md = goldmark.New(
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

// LinkRenderer is a custom renderer for links
type LinkRenderer struct{}

// RegisterFuncs implements NodeRenderer.RegisterFuncs
func (r *LinkRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindLink, r.renderLink)
}

// renderLink renders links, converting .md extensions to .html
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
	relPath, _ := filepath.Rel(CONTENT_DIRECTORY, path)
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
	fs := http.FileServer(http.Dir(OUT_DIRECTORY))
	http.Handle("/", fs)

	log.Printf("Serving site on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
