package jo

import (
	"fmt"
)

// if the first letter of a function is capitalized, it is exported (Public)
func SayHello() {
	fmt.Println("Hello from the Jo package!")
}

// if the first letter of a function is not capitalized, it is unexported (Private)
//this function is private to the package (same package name can access it)
func sayGoodbye() {
	fmt.Println("Goodbye from the Jo package!") //this function is private to the package
}
