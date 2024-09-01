package builder

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type Builder struct {
	TempDir    string
	SourceDir  string
	HeaderHtml string
	StyleSheet string
	md         goldmark.Markdown
	FileInfo   map[string]time.Time
}

func (b *Builder) GenerateSite() {

	err := copyCSSFile(b.TempDir, b.StyleSheet)
	if err != nil {
		log.Printf("Error copying CSS file: %v\n", err)
		return
	}

	b.initializeGoldmark()
	b.loadFileInfo()

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

	b.saveFileInfo()
}

func copyCSSFile(tempDir string, stylesheet string) error {

	err := os.MkdirAll(filepath.Join(tempDir, "styles"), os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Println("using theme : ", stylesheet)
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
		// goldmark.WithRendererOptions(
		// 	html.WithHardWraps(),
		// 	html.WithXHTML(),
		// ),
	)
}

func (b *Builder) processMarkdownFile(path string) error {

	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to get file info for %s: %w", path, err)
	}

	if lastMod, exists := b.FileInfo[path]; exists && lastMod.Equal(info.ModTime()) {
		fmt.Printf("Skipping unchanged file: %s\n", path)
		return nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	preProcessedContent := b.preprocessMarkdown(string(content))

	var buf bytes.Buffer
	if err := b.md.Convert([]byte(preProcessedContent), &buf); err != nil {
		return fmt.Errorf("failed to convert markdown: %w", err)
	}

	relPath, err := filepath.Rel(b.SourceDir, path)
	if err != nil {
		return fmt.Errorf("failed to get relative path: %w", err)
	}

	outPath := filepath.Join(b.TempDir, strings.TrimSuffix(relPath, ".md")+".html")
	if err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	title := strings.TrimSuffix(filepath.Base(path), ".md")
	fullHTML := fmt.Sprintf(b.HeaderHtml, title, buf.String())

	if err := os.WriteFile(outPath, []byte(fullHTML), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outPath, err)
	}

	b.FileInfo[path] = info.ModTime()

	fmt.Printf("Generated: %s\n", outPath)
	return nil
}

func (b *Builder) preprocessMarkdown(content string) string {
	re := regexp.MustCompile(`\[(.*?)\]\((.*?\.md)\)`)
	return re.ReplaceAllStringFunc(content, func(link string) string {
		return strings.Replace(link, ".md", ".html", 1)
	})
}

func (b *Builder) loadFileInfo() error {
	b.FileInfo = make(map[string]time.Time)
	infoFile := filepath.Join(b.TempDir, "content.modtime")
	file, err := os.Open(infoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), ":", 2)
		if len(parts) == 2 {
			if t, err := time.Parse(time.RFC3339Nano, parts[1]); err == nil {
				b.FileInfo[parts[0]] = t
			}
		}
	}
	return scanner.Err()
}

func (b *Builder) saveFileInfo() error {
	infoFile := filepath.Join(b.TempDir, "content.modtime")
	file, err := os.Create(infoFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for path, modTime := range b.FileInfo {
		_, err := fmt.Fprintf(file, "%s:%s\n", path, modTime.Format(time.RFC3339Nano))
		if err != nil {
			return err
		}
	}
	return nil
}
