package c

/*
#include <string.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// Malloc 开辟内存
func Malloc(size int) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

// Free 释放内存
func Free(data unsafe.Pointer) {
	C.free(data)
}

// Memmove 内存移动
func Memmove(dest, src unsafe.Pointer, length int) {
	C.memmove(dest, src, C.size_t(length))
}

// Memcpy 内存复制
func Memcpy(dest unsafe.Pointer, src []byte, length int) {
	srcData := C.CBytes(src)
	C.memcpy(dest, srcData, C.size_t(length))
}
