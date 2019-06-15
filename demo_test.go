package wasm

import (
	"testing"
)

func TestSimple(t *testing.T) {
	simple, err := Read("examples/simple/build/simple.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(simple, "sum", []interface{}{int32(17), int32(102)})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res != 119 {
		t.Fatalf("Unexpected result: %d", res)
	}
}

func TestFib(t *testing.T) {
	fib, err := Read("examples/fib_recursive/build/fib_recursive.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(fib, "fib", []interface{}{int32(8)})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res != 21 { // fib(8)
		t.Fatalf("Unexpected result: %d", res)
	}
}
