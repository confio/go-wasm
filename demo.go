package wasm

import (
	"fmt"

	"github.com/pkg/errors"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

// Read loads a wasm file
func Read(filename string) ([]byte, error) {
	return wasm.ReadBytes(filename)
}

// Run will execute the named function on the wasm bytes with the passed arguments.
// Returns the result or an error
func Run(code []byte, call string, args []interface{}) (*wasm.Value, error) {

	// Instantiates the WebAssembly module.
	instance, err := wasm.NewInstance(code)
	if err != nil {
		return nil, errors.Wrap(err, "init wasmer")
	}
	defer instance.Close()

	f, ok := instance.Exports[call]
	if !ok {
		return nil, errors.Errorf("Function %s not in Exports", call)
	}

	ret, err := f(args...)
	if err != nil {
		return nil, errors.Wrap(err, "Execution failure")
	}
	fmt.Printf("%v: %v\n", ret.GetType(), ret)

	return &ret, nil
}
