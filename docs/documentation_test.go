package docs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderHTML(t *testing.T) {
	html, err := RenderHTML("test.md")
	require.Nil(t, err)

	expectedHtml := "<h1 id=\"test\">Test</h1>\n\n<p>This is a test</p>\n"

	require.Equal(t, expectedHtml, html)
}
