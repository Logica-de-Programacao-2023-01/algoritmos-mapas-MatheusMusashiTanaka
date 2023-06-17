package main

import (
	"fmt"
	"strings"
)

func contarPalavras(s string) map[string]int {

	frequenciaPalavras := make(map[string]int)

	strings.Fields(s)

	for _, ranS := range strings.Fields(s) {

		frequenciaPalavras[ranS] += 1

	}

	return frequenciaPalavras

}

func main() {

	s := "oba oba he he kkkkkkkkkkkkkk oba"

	fmt.Println(contarPalavras(s))

}
