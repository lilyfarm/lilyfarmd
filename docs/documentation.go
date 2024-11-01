package docs

import (
	"embed"
	"fmt"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed markdown/*.md
var markdownFS embed.FS

func convertMarkdownToHTML(md []byte) string {
	// Create a new parser with CommonMarkdown extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	// Render to HTML
	html := markdown.ToHTML(md, p, nil)

	return string(html)
}

// RenderHTML reads a markdown file in the "markdown" directory and returns
// a rendered HTML version of it.
func RenderHTML(markdownFile string) (string, error) {
	file, err := markdownFS.ReadFile("markdown/" + markdownFile)

	if err != nil {
		return "", fmt.Errorf("could not read markdown: %w", err)
	}

	return convertMarkdownToHTML(file), nil
}
