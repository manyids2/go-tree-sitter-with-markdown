package markdown_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/markdown"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("# hi\n## hello"), markdown.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(module (expression_statement (call function: (identifier) arguments: (argument_list (integer)))))",
		n.String(),
	)
}
