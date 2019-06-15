package wasm

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/perlin-network/life/exec"
)

type Resolver struct{}

func (r *Resolver) ResolveFunc(module, field string) exec.FunctionImport {
	switch module {
	case "env":
		switch field {
		case "__life_log":
			return func(vm *exec.VirtualMachine) int64 {
				ptr := int(uint32(vm.GetCurrentFrame().Locals[0]))
				msgLen := int(uint32(vm.GetCurrentFrame().Locals[1]))
				msg := vm.Memory[ptr : ptr+msgLen]
				fmt.Printf("[app] %s\n", string(msg))
				return 0
			}
		default:
			panic(fmt.Errorf("unknown import resolved: %s", field))
		}
	default:
		panic(fmt.Errorf("unknown module: %s", module))
	}
}

func (r *Resolver) ResolveGlobal(module, field string) int64 {
	panic("we're not resolving global variables for now")
}

var callbacks exec.ImportResolver = &Resolver{}

func TestRun(t *testing.T) {
	input, err := ioutil.ReadFile("examples/fib_recursive/build/fib_recursive.wasm")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	res, err := Run(input, callbacks, "fib", []int64{8}, nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res != 21 { // fib(8)
		t.Fatalf("Unexpected result: %d", res)
	}

	data := []byte(`{"number": 10}`)
	res, err = Run(input, callbacks, "app_main", nil, data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res != 55 { // fib(10)
		t.Fatalf("Unexpected result for fib(10): %d", res)
	}
}
