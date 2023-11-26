package toml_test

import (
	"context"
	"testing"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/toml"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`key = "value"`), toml.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (pair (bare_key) (string)))",
		n.String(),
	)
}
