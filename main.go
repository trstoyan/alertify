package main

import (
	"fmt"
)

func myname(name string) string {
	return name
}

func main() {
	fmt.Printf("Hello, World! %s", myname("Sarath"))
}
