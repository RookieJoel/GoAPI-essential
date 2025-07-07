package jo

import (
	"github.com/RookieJoel/GoAPI-essential/Jo/internal/testInternal" // Importing the internal test package
)

func SayHelloFromInternal() {
	// Calling the exported function from the internal package
	testinternal.SayHelloFromInternal()
}