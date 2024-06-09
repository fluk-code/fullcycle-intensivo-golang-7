package main

import "fmt"

type Carro struct {
	Nome string
	Modelo string
	Ano int
}

func (c Carro) Andar() {
	fmt.Println("O carro " + c.Nome + " está andando!")
}

func (c Carro) Parar() {
	fmt.Println("O carro " + c.Nome + " está parando!")
}


func main() {
	carro := Carro{Nome: "ford", Modelo: "Mustang", Ano: 1969}
	carro2 := Carro{Nome: "chevrolet", Modelo: "Camaro", Ano: 1969}

	carro.Andar()
	carro2.Andar()
	carro.Parar()
}