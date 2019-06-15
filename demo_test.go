package wasm

import (
	"testing"
)

func TestSimple(t *testing.T) {
	simple, err := Read("examples/simple/build/simple.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(simple, nil, "sum", []interface{}{int32(17), int32(102)}, AsInt32)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res.(int32) != 119 {
		t.Fatalf("Unexpected result: %d", res)
	}
}

func TestFib(t *testing.T) {
	fib, err := Read("examples/fib_recursive/build/fib_recursive.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(fib, nil, "fib", []interface{}{int32(8)}, AsInt32)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res.(int32) != 21 { // fib(8)
		t.Fatalf("Unexpected result: %d", res)
	}
}

func TestGreet(t *testing.T) {
	simple, err := Read("examples/greet/build/greet.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(simple, nil, "greet", []interface{}{`{"number": 2}`}, AsString)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res.(string) != "Hello, Hello, !" {
		t.Fatalf("Unexpected result: %d", res)
	}
}

func TestStringJSON(t *testing.T) {
	simple, err := Read("examples/string_json/build/string_json.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(simple, nil, "greet", []interface{}{`{"number": 3, "message": "Gaia "}`}, AsString)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res.(string) != "Hello, Gaia Gaia Gaia !" {
		t.Fatalf("Unexpected result: %d", res)
	}
}

func TestImportFunc(t *testing.T) {
	simple, err := Read("examples/import_func/build/import_func.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	imports, err := importsWithSum()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(simple, imports, "add1", []interface{}{int32(7), int32(9)}, AsInt32)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res.(int32) != 17 {
		t.Fatalf("Unexpected result: %d", res)
	}
}
