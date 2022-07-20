all: build

build: static/parser.wasm static/wasm_exec.js

static/parser.wasm: main.go
	tinygo build -target wasm -o $@ $<

static/wasm_exec.js:
	cp $$TINYGOROOT/targets/wasm_exec.js $@
