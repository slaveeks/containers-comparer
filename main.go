package main

import (
	"awesomeProject/src/comparer"
	"fmt"
)

func main() {
	fmt.Println("123")
	c := comparer.CreateUtilTester("podman", "./test-app/check")
	c.BuildImage()
}
