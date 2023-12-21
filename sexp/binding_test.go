package sexp_test

import (
	"context"
	"testing"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/sexp"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("(module (expression_statement (call function: (identifier) arguments: (argument_list (integer)))))"), sexp.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(sexp (list (atom) (list (atom) (list (atom) (atom) (list (atom)) (atom) (list (atom) (list (atom)))))))",
		n.String(),
	)
}
