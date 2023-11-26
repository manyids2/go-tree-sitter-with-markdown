package elm_test

import (
	"context"
	"testing"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/elm"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("import Html exposing (text)"), elm.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(file (import_clause (import) moduleName: (upper_case_qid (upper_case_identifier)) exposing: (exposing_list (exposing) (exposed_value (lower_case_identifier)))))",
		n.String(),
	)
}
