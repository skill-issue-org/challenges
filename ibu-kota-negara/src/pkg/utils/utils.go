package utils

import (
	"unsafe"
)

func EAX() uint8 {
	return uint8(unsafe.Sizeof(true))
}
