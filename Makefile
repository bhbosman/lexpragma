all: clean lexdata build test


build:
	go build -v
clean:
	rm -f handler.*.go
	rm -f golexidl

lexdata: lexpragma.l
	golex -o handler.ReadLexem.temp         lexpragma.l
	gofmt  handler.ReadLexem.temp > lexpragma.l.go
	rm -f handler.ReadLexem.temp
test:
	go test -v