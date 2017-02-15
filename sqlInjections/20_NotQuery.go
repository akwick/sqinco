package sqlInjections

import "fmt"

func NotQuerry() {
	a := "(1) lovely"
	// @ExpectWarning false
	s := message("SELECT a gopher=" + a)
	fmt.Print(s)
}

func message(a string) string {
	return a
}
