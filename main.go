package main

import (
	"syscall/js"

	"github.com/open-policy-agent/opa/ast"
)

func main() {
}

//export update
func update() {
	document := js.Global().Get("document")
	document.Call("getElementById", "error").Set("value", "error goes here") // debug startup
	aStr := document.Call("getElementById", "rego").Get("value").String()
	x, err := ast.ParseModule("policy.rego", aStr)
	if err != nil {
		document.Call("getElementById", "error").Set("value", err.Error())
		return
	}
	document.Call("getElementById", "result").Set("value", x.String())
}