package main

import "fmt"

func preencherPrimos(slice *[]int, tamanho int) {
	numerosPrimos := []int{}
	numero := 2

	for len(numerosPrimos) < tamanho {
		if ehPrimo(numero) {
			numerosPrimos = append(numerosPrimos, numero)
		}
		numero++
	}

	*slice = numerosPrimos
}

func ehPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	tamanho := 10
	slice := []int{}

	preencherPrimos(&slice, tamanho)

	fmt.Println("Slice preenchida com os", tamanho, "primeiros números primos:")
	fmt.Println(slice)
}
