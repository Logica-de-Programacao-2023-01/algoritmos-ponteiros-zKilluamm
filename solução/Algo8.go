package main

import (
	"fmt"
	"unsafe"
)

func modificarValor(ponteiro *int) {
	novoValor := 42
	*ponteiro = novoValor
}

func main() {
	valor := 10
	ponteiro := &valor

	fmt.Println("Antes da modificação:", *ponteiro)

	modificarValor(ponteiro)

	fmt.Println("Após a modificação:", *ponteiro)

	// Liberando a memória após o uso
	ponteiro = nil
	valor = 0

	// Verificando se o ponteiro foi liberado
	if ponteiro == nil {
		fmt.Println("Memória liberada.")
	} else {
		fmt.Println("A memória ainda está alocada.")
	}

	// Verificando o tamanho do ponteiro em bytes
	tamanho := unsafe.Sizeof(ponteiro)
	fmt.Println("Tamanho do ponteiro:", tamanho, "bytes")
}
