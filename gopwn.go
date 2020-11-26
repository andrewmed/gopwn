package gopwn

import "fmt"

func p(err error) {
	if err != nil {
		panic(err)
	}
}

func AddressFmt(i interface{}) uintptr {
	var addr uintptr
	_, err := fmt.Sscanf(fmt.Sprintf("%p", i), "0x%x", &addr)
	p(err)
	return addr
}

