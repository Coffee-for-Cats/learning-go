package main

import "fmt"

const stringOla = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return stringOla + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
