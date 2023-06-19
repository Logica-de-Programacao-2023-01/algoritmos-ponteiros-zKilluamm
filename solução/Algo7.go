package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValor(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   1000.0,
		Titular: "João",
	}

	valorAdicional := 500.0
	adicionarValor(&conta, valorAdicional)

	fmt.Println("Novo saldo:", conta.Saldo)
}
