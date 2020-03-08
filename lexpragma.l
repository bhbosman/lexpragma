%{
package lexpragma

import "strings"

import yaccPragmaToken "github.com/bhbosman/yaccpragma"

	const (
		INITIAL = iota
	)


func (self *Handler) readLexem() Lexem {

%}

%yyc self.current
%yyn self.current = self.GetChar()
%yyt self.startCond

D  [0-9]+
E  [eE][-+]?{D}
F  {D}"."{D}?{E}?|{D}{E}?|"."{D}{E}?

%%
    self.buf = self.buf[:0]


"#"pragma return NewLexemStringValue(yaccPragmaToken.RwPragma, self.TokenName, string(self.buf))
prefix return NewLexemStringValue(yaccPragmaToken.RwPrefix, self.TokenName, string(self.buf))
ID return NewLexemStringValue(yaccPragmaToken.RwId, self.TokenName, string(self.buf))
version return NewLexemStringValue(yaccPragmaToken.RwVersion, self.TokenName, string(self.buf))

L?\"(\\.|[^\\"])*\"
    {
        s := strings.TrimSpace(string(self.buf[1:len(self.buf)-1]))
        return NewLexemStringValue(yaccPragmaToken.StringLiteral, self.TokenName, s)
    }
[ \t\r\n]+       return NewLexemStringValue(yaccPragmaToken.WhiteSpace, self.TokenName, string(self.buf))
"."              return NewLexemNoValue(int('.'), self.TokenName)
[a-zA-Z_][a-zA-Z0-9_]*
        s := strings.TrimSpace(string(self.buf[:len(self.buf)]))
		return NewLexemStringValue(yaccPragmaToken.Identifier, self.TokenName, s)
%%
    return NewLexemEofValue()
}





