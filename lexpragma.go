package lexpragma

import (
	"bufio"
	"bytes"
	"github.com/bhbosman/gocommon"
	"github.com/bhbosman/yaccpragma"
	"os"
)

type Handler struct {
	buf             []byte
	readerCloser    gocommon.ByteReaderCloser
	current         byte
	startCond       int
	prevCond        int
	currentAssigned bool
}

func NewHandler(byteReaderCloser gocommon.ByteReaderCloser) (*Handler, error) {
	handler := &Handler{
		buf:          nil,
		readerCloser: byteReaderCloser,
		current:      0,
		startCond:    0,
		prevCond:     0,
		//definedFiles:    nil,
		currentAssigned: false,
	}
	handler.GetChar()
	return handler, nil
}

func NewHandlerFromFileName(fileName string) (*Handler, error) {
	f, e := os.Open(fileName)
	if e != nil {
		return nil, e
	}
	return NewHandler(gocommon.NewByteReaderWithCloser(f, bufio.NewReaderSize(f, 4096)))
}

func NewPragmaLexFromData(streamName, s string) (*Handler, error) {
	return NewHandler(gocommon.NewByteReaderNoCloser(bytes.NewBufferString(s)))
}

func (self *Handler) initStream() {
	self.currentAssigned = false
	self.current = 0
}

func (self *Handler) GetChar() byte {
	if self.current != 0 {
		self.buf = append(self.buf, self.current)
	}
	self.current = 0

	var b byte
	var e error
	if b, e = self.readerCloser.ReadByte(); e == nil {
		self.current = b
	}

	return self.current
}

type Error struct {
	s string
}

func NewError(s string) *Error {
	return &Error{s: s}
}

func (e Error) Error() string {
	return e.s
}

func createError(s string) error {
	return NewError(s)
}

func (self *Handler) ReadLexem() Lexem {

	lexem := self.readLexem()
	return lexem
}

func (self Handler) TokenName(tokenId int) string {
	if tokenId < 255 {
		return string(byte(tokenId))
	}

	return yaccpragma.YaccPragmaTokname(yaccpragma.YaccPragmaTok2[tokenId-yaccpragma.YaccPragmaPrivate])
}
