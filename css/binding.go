package css

//#include "parser.h"
//TSLanguage *tree_sitter_css();
import "C"
import (
	"unsafe"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_css())
	return sitter.NewLanguage(ptr)
}
