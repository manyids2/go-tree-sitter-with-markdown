package sexp

//#include "tree_sitter/parser.h"
//TSLanguage *tree_sitter_sexp();
import "C"

import (
	"unsafe"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_sexp())
	return sitter.NewLanguage(ptr)
}
