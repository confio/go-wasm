package wasm

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

// Read loads a wasm file
func Read(filename string) ([]byte, error) {
	return wasm.ReadBytes(filename)
}

type ResultParser func(wasm.Instance, wasm.Value) (interface{}, error)

func AsInt32(_ wasm.Instance, res wasm.Value) (interface{}, error) {
	return res.ToI32(), nil
}

func AsInt64(_ wasm.Instance, res wasm.Value) (interface{}, error) {
	return res.ToI64(), nil
}

func AsString(instance wasm.Instance, res wasm.Value) (interface{}, error) {
	outputPointer := res.ToI32()
	memory := instance.Memory.Data()[outputPointer:]
	nth := 0
	var output strings.Builder

	for {
		if memory[nth] == 0 {
			break
		}

		output.WriteByte(memory[nth])
		nth++
	}

	lengthOfOutput := nth

	// Deallocate the subject, and the output.
	deallocate := instance.Exports["deallocate"]
	// TODO
	// deallocate(inputPointer, lengthOfSubject)
	deallocate(outputPointer, lengthOfOutput)

	return output.String(), nil
}

// Run will execute the named function on the wasm bytes with the passed arguments.
// Returns the result or an error
func Run(code []byte, call string, args []interface{}, parse ResultParser) (interface{}, error) {
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

	fArgs := prepareArgs(instance, args)

	ret, err := f(fArgs...)
	if err != nil {
		return nil, errors.Wrap(err, "Execution failure")
	}
	fmt.Printf("%v: %v\n", ret.GetType(), ret)

	return parse(instance, ret)
}

func prepareArgs(instance wasm.Instance, args []interface{}) []interface{} {
	out := make([]interface{}, len(args))

	for i, arg := range args {
		switch t := arg.(type) {
		case int32, int64:
			out[i] = arg
		case string:
			out[i] = prepareString(instance, t)
		case []byte:
			out[i] = prepareString(instance, string(t))
		default:
			panic(fmt.Sprintf("Unsupported type: %T", arg))
		}
	}
	return out
}

func prepareString(instance wasm.Instance, arg string) int32 {
	l := len(arg)
	allocateResult, _ := instance.Exports["allocate"](l)
	inputPointer := allocateResult.ToI32()

	// Write the subject into the memory.
	memory := instance.Memory.Data()[inputPointer:]

	for nth := 0; nth < l; nth++ {
		memory[nth] = arg[nth]
	}

	// C-string terminates by NULL.
	memory[l] = 0

	return inputPointer
}
