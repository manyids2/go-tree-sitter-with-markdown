package csharp_test

import (
	"context"
	"testing"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/csharp"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("using static System.Math;"), csharp.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(compilation_unit (using_directive (qualified_name (identifier) (identifier))))",
		n.String(),
	)
}
