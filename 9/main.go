package main

import (
	
	"fmt"
)


func main() {
   fmt.Print (sum(1,2,3,4,5,6,7,8,9,12334,32423534,4576578578,897668,231231231))
}

func sum (numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }
    return total
}
