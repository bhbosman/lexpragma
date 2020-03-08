package lexpragma

type Lexem struct {
	TokenName string
	TypeKind  int
	Value     string
	Eof       bool
}

func (l Lexem) CheckTarget() bool {
	return false
}

type LexemCollection struct {
	LexemValues []Lexem
}

func NewLexemEofValue() Lexem {

	return Lexem{
		Eof: true,
	}
}

func NewLexemNoValue(typeKind int, tokenName func(int) string) Lexem {
	return Lexem{
		TypeKind:  typeKind,
		TokenName: tokenName(typeKind),
		Eof:       false,
	}
}

func NewLexemStringValue(typeKind int, tokenName func(int) string, stringValue string) Lexem {
	return Lexem{
		TypeKind:  typeKind,
		TokenName: tokenName(typeKind),
		Value:     stringValue,
		Eof:       false,
	}
}
