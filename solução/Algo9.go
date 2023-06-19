package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	desconto := 0.10 // 10% de desconto
	livro.Preco = livro.Preco - (livro.Preco * desconto)
}

func main() {
	livro := Livro{
		Titulo: "Livro X",
		Autor:  "Autor Y",
		Preco:  100.0,
	}

	fmt.Println("Preço original:", livro.Preco)

	aplicarDesconto(&livro)

	fmt.Println("Preço com desconto:", livro.Preco)
}
