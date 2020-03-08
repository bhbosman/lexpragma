package lexpragma

import (
	"github.com/bhbosman/yaccpragma"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_Hex(t *testing.T) {
	t.Run("Test Identifier", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "a")
		assert.NoError(t, e)
		lexem := handler.ReadLexem()
		assert.Equal(t, "a", lexem.Value)
		assert.Equal(t, yaccpragma.Identifier, lexem.TypeKind)
	})

	t.Run("Test #pragma", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "#pragma")
		assert.NoError(t, e)
		lexem := handler.ReadLexem()
		assert.Equal(t, yaccpragma.RwPragma, lexem.TypeKind)
	})

	t.Run("Test prefix", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "prefix")
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.RwPrefix, lexem.TypeKind)
	})

	t.Run("Test ID", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "ID")
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.RwId, lexem.TypeKind)
	})

	t.Run("Test string", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "\"abc\"")
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.StringLiteral, lexem.TypeKind)
	})

	t.Run("Test Whitespace", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "\t\t\n ")
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.WhiteSpace, lexem.TypeKind)
	})

	t.Run("Test Whitespace", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "version")
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.RwVersion, lexem.TypeKind)
	})

}
