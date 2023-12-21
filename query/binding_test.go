package query_test

import (
	"context"
	"testing"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/query"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("(type_identifier) @type"), query.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (named_node name: (identifier) (capture name: (identifier))))",
		n.String(),
	)
}
