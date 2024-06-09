package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Carro struct {
	Nome string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano int `json:"-"`
}

func (c Carro) Andar() {
	fmt.Println("O carro " + c.Nome + " está andando!")
}

func (c Carro) Parar() {
	fmt.Println("O carro " + c.Nome + " está parando!")
}


func main() {
	carro := Carro{Nome: "ford", Modelo: "Mustang", Ano: 1969}

	http.HandleFunc("/", func( w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(carro)
	})

	http.ListenAndServe(":3333", nil)
}