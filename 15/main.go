package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello world!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("o tipo da variável é %T e o valor é %v\n", t, t)
}
