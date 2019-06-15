package wasm

import (
	"github.com/pkg/errors"

	"github.com/perlin-network/life/exec"
	wasm_validation "github.com/perlin-network/life/wasm-validation"
)

var defaultConfig = exec.VMConfig{
	DefaultMemoryPages:   128,
	DefaultTableSize:     65536,
	DisableFloatingPoint: true,
}

// Run will execute the named function on the wasm bytes with the passed arguments.
// Returns the result or an error
func Run(wasm []byte, resolver exec.ImportResolver, call string, args []int64) (int64, error) {
	validator, err := wasm_validation.NewValidator()
	if err != nil {
		return 0, errors.Wrap(err, "init validator")
	}

	err = validator.ValidateWasm(wasm)
	if err != nil {
		return 0, errors.Wrap(err, "validate wasm")
	}

	vm, err := exec.NewVirtualMachine(wasm, defaultConfig, resolver, nil)
	if err != nil {
		return 0, errors.Wrap(err, "init vm")
	}

	entryID, ok := vm.GetFunctionExport(call)
	if !ok {
		return 0, errors.Errorf("Entry function %s not found", call)
	}

	ret, err := vm.Run(entryID, args...)
	if err != nil {
		vm.PrintStackTrace()
		return 0, errors.Wrap(err, "Execution failure")
	}

	return ret, nil
}
