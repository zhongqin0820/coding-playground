package myPack

import (
	"fmt"
)

var MyPackA int = 1

func hi() {
	fmt.Println("this is private func")
}

func Hi() {
	fmt.Println("this is public func")
}
