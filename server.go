package main

import (
	"fmt"
	"hiring_test/servicelayer"
)

func main() {
	fmt.Println("Server started.")
	servicelayer.HandleRequests()
	fmt.Println("Program exited.")
}
