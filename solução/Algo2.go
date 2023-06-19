package main

import "fmt"

func verificarParImpar(ponteiro *int) {
	if *ponteiro%2 == 0 {
		*ponteiro = 0 // Par
	} else {
		*ponteiro = 1 // Ímpar
	}
}

func main() {
	numero := 7
	verificarParImpar(&numero)
	fmt.Println("Número atualizado:", numero)
}
