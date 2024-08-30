package builder

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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

type Builder struct {
	TempDir    string
	SourceDir  string
	HeaderHtml string
	StyleSheet string
	md         goldmark.Markdown
}

func (b *Builder) GenerateSite() {

	err := copyCSSFile(b.TempDir, b.StyleSheet)
	if err != nil {
		log.Printf("Error copying CSS file: %v\n", err)
		return
	}

	b.initializeGoldmark()

	os.MkdirAll(b.TempDir, os.ModePerm)

	err = filepath.Walk(b.SourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		return b.processMarkdownFile(path)
	})

	if err != nil {
		log.Printf("Error generating site: %v\n", err)
	} else {
		log.Println("Site generated successfully")
	}
}

func copyCSSFile(tempDir string, stylesheet string) error {

	err := os.MkdirAll(filepath.Join(tempDir, "styles"), os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Println(stylesheet)
	source, err := os.Open(stylesheet)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(filepath.Join(tempDir, "styles", "default.css"))
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func (b *Builder) initializeGoldmark() {

	b.md = goldmark.New(
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

func (b *Builder) processMarkdownFile(path string) error {

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := b.md.Convert(content, &buf); err != nil {
		return err
	}
	html := buf.Bytes()
	relPath, _ := filepath.Rel(b.SourceDir, path)
	outPath := filepath.Join(b.TempDir, strings.TrimSuffix(relPath, ".md")+".html")
	os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	fullHTML := fmt.Sprintf(b.HeaderHtml, strings.TrimSuffix(filepath.Base(path), ".md"), string(html))
	err = os.WriteFile(outPath, []byte(fullHTML), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Generated: %s\n", outPath)
	return nil
}
