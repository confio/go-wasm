package wasm

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
//
// extern int32_t repeat(void *context, int32_t pointer, int32_t length, int32_t count, int32_t outputPtr, int32_t outputLen);
import "C"

import (
	"fmt"
	"unsafe"
	"strings"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

func importsWithSum() (*wasm.Imports, error) {
	return wasm.NewImports().Append("sum", sum, C.sum)
}

//export repeat
func repeat(context unsafe.Pointer, pointer int32, length int32, count int32, outputPtr int32, outputLen int32) int32 {
	var instanceContext = wasm.IntoInstanceContext(context)
	var memory = instanceContext.Memory().Data()
	text := string(memory[pointer : pointer+length])

	res := strings.Repeat(text, int(count))
	fmt.Println(res)

	// zero append string
	write := append([]byte(res), byte(0))
	copy(memory[outputPtr:outputPtr+outputLen], write)
	return int32(len(res))
}

func importsWithRepeatAndSum() (*wasm.Imports, error) {
	imp, err := wasm.NewImports().Append("repeat", repeat, C.repeat)
	if err != nil {
		return nil, err
	}
	return imp.Append("sum", sum, C.sum)
}