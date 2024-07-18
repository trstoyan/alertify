package main

import (
	"fmt"

	"github.com/trstoyan/alertify/api"
)

func init() {
	fmt.Println("Initializing...")
}

func main() {
	fmt.Println(api.Myname("John Doe"))
	fmt.Println("Hello, World!")
}
