all: build

build: static/parser.wasm static/wasm_exec.js

.PHONY: static/parser.wasm
static/parser.wasm: main.go
	tinygo build -target wasm -gc=leaking -o $@ $<

static/wasm_exec.js:
	cp $$TINYGOROOT/targets/wasm_exec.js $@
