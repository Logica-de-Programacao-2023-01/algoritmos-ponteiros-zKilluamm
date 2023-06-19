package main

import "fmt"

func atualizarValor(ponteiro *int) {
	numero := *ponteiro
	digito1 := numero % 10
	digito2 := (numero / 10) % 10
	soma := digito1 + digito2
	*ponteiro = soma
}

func main() {
	valor := 1234
	atualizarValor(&valor)
	fmt.Println("Valor atualizado:", valor)
}
