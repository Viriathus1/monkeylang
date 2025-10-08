package main

import (
	"flag"
	"fmt"
	"monkeylang/compiler"
	"monkeylang/evaluator"
	"monkeylang/lexer"
	"monkeylang/object"
	"monkeylang/parser"
	"monkeylang/vm"
	"time"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'engine'")

var input = `
let fibonacci = fn(x) {
	if (x == 0) {
		return 0;
	} else {
		if (x == 1) {
			return 1;
		} else {
			fibonacci(x - 1) + fibonacci(x - 2)
		}
	}
};
fibonacci(35);
`

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if *engine == "vm" {
		comp := compiler.New()
		if err := comp.Compile(program); err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.Bytecode())

		start := time.Now()

		if err := machine.Run(); err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		duration = time.Since(start)
		result = machine.LastPoppedStackElem()
	} else {
		env := object.NewEnvironment()
		start := time.Now()
		result = evaluator.Eval(program, env)
		duration = time.Since(start)
	}

	fmt.Printf(
		"engine=%s\nresult=%sduration=%s\n",
		*engine,
		result.Inspect(),
		duration,
	)
}
