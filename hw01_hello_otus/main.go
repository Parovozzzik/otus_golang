package main

import (
	"fmt"

	strutil "golang.org/x/example/stringutil"
)

func main() {
	a := strutil.Reverse("Hello, OTUS!")
	fmt.Println(a)
}
