GOEXE ?= go

generate:
	$(GOEXE) generate ./...
.PHONY: generate

test:
	$(GOEXE) test ./...
.PHONY: test

clobber:
	$(GOEXE) run ./generate/tools/clobbergen.go ./macos
.PHONY: clobber

example:
	$(GOEXE) run ./macos/_examples/helloworld/main.go
.PHONY: example

generate/symbols.zip:
	cd generate && wget https://github.com/mactypes/symbolsdb/releases/download/1.1/symbols.zip