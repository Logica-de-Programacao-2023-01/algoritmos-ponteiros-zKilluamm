package main

import (
	"fmt"
	"math"
)

func atualizarValor(ponteiro *float64) {
	pi := math.Pi
	valorAtual := *ponteiro
	media := (valorAtual + pi) / 2
	*ponteiro = media
}

func main() {
	valor := 3.14
	atualizarValor(&valor)
	fmt.Println("Valor atualizado:", valor)
}
