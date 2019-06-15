package wasm

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	input, err := ioutil.ReadFile("wasm/fib_recursive.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(input, "app_main", nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res != 9227465 { // fib(35)
		t.Fatalf("Unexpected result: %d", res)
	}
}