package main

import "fmt"

var x int = 10

func main() {
	x := 42
	// x = 42 上面的方式和这中结果不同
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	fmt.Println(x)
}
