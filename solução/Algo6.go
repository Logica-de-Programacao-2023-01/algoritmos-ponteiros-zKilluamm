package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func atualizarTitulo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro X",
		Autor:  "Anônimo",
	}
	atualizarTitulo(&livro)
	fmt.Println("Título atualizado:", livro.Titulo)
}
