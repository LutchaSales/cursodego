package main


import "fmt"

func main() {
	salarios := map[string]int{"lutcha": 1000, "joão": 20000, "maria": 3000}

	fmt.Println(salarios["lutcha"])
    delete(salarios, "lutcha")
    salarios["luciano"] = 5000
    fmt.Println(salarios["luciano"])

    //sal := make(map[string]int)
    //sal1 := map[string]int{}
    //sal1["patricia"] = 15000

    for nome, salario := range salarios {
        fmt.Printf("o salario de %s é %d\n", nome, salario)
    }

    for _, salario := range salarios {
        fmt.Printf("o salario é %d\n", salario)
    }
}
