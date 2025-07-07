package testinternal

import (
	"fmt"
)

// SayHelloFromInternal is an exported function that can be called from other packages within the same module.
func SayHelloFromInternal() {
	fmt.Println("Hello from the internal test package!")
}