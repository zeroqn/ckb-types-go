// MIT

// WARNING: This file has automatically been generated on Thu, 24 Oct 2019 15:22:37 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package types

/*
#cgo LDFLAGS: -static -L${SRCDIR}/../ -lckb_ffi -lpthread -ldl
#include "../ckb_ffi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// Buffer as declared in ckb-types-go/ckb_ffi.h:22
type Buffer struct {
	Len            uint
	Data           *byte
	refb0a5a638    *C.buffer_t
	allocsb0a5a638 interface{}
}
