package main

import (
	"fmt"
)

type Car struct {
	Name  string
	Model string
}

func (c Car) Andar() {
	fmt.Println("O Carro", c.Name, "está andando")
}

func main() {

	carro := Car{"Fusca", "VW"}
	carro.Model = "Audi"
	fmt.Println(carro.Model)

	result, err := soma(1, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result, err)
	carro.Andar()
}

func soma(a, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("soma maior que 10")
	}
	return a + b, nil
}
