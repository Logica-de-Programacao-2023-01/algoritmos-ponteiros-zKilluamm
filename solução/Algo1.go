package main

import "fmt"

func atualizarValor(ponteiro *int, n int) {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	*ponteiro = soma
}

func main() {
	valor := 0
	n := 5
	atualizarValor(&valor, n)
	fmt.Println("Valor atualizado:", valor)
}
