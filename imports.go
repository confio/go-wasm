package wasm

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
//
// extern int32_t repeat(void *context, int32_t pointer, int32_t length, int32_t count);
import "C"

import (
	"unsafe"
	"strings"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

var curInstance *wasm.Instance
 
//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

func importsWithSum() (*wasm.Imports, error) {
	return wasm.NewImports().Append("sum", sum, C.sum)
}

//export repeat
func repeat(context unsafe.Pointer, pointer int32, length int32, count int32) int32 {
	var instanceContext = wasm.IntoInstanceContext(context)
	var memory = instanceContext.Memory().Data()
	text := string(memory[pointer : pointer+length])

	res := strings.Repeat(text, int(count))

	outputPtr := prepareString(*curInstance, res)
	return outputPtr
}

func importsWithRepeatAndSum() (*wasm.Imports, error) {
	imp, err := wasm.NewImports().Append("repeat", repeat, C.repeat)
	if err != nil {
		return nil, err
	}
	return imp.Append("sum", sum, C.sum)
}