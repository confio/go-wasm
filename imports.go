package wasm

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
import "C"

import (
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"unsafe"
)

//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}

func importsWithSum() (*wasm.Imports, error) {
	return wasm.NewImports().Append("sum", sum, C.sum)
}