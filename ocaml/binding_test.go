package ocaml_test

import (
	"context"
	"testing"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/ocaml"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`print_endline "Hello World!"`), ocaml.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(compilation_unit (expression_item (application_expression function: (value_path (value_name)) argument: (string (string_content)))))",
		n.String(),
	)
}
