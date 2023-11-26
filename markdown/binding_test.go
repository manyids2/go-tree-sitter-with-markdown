package markdown_test

import (
	"context"
	"testing"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/markdown"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("# hi\n## hello"), markdown.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (atx_heading (atx_h1_marker) (heading_content (text))) (atx_heading (atx_h2_marker) (heading_content (text))))",
		n.String(),
	)
}
