package main

import "fmt"
func main() {
	var minhaVar interface{} = "Luciano Sales"

	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
