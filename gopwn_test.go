package gopwn

import (
	"testing"
	"unsafe"
)

func TestAddress(t *testing.T) {
	val := 1
	addr := AddressFmt(&val)
	addr1 := uintptr(unsafe.Pointer(&val))
	if addr != addr1 {
		t.Fatalf("expect equal, got: %d and %d", addr, addr1)
	}
}

