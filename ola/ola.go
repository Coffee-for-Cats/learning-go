package main

import "fmt"

const stringPortugues = "Olá, "
const stringEspanhol = "Hola, "
const stringFrances = "Bonjour, "

func Ola(nome string, idioma string) string {

	if nome == "" {
		nome = "Mundo"
	}

	return qualPrefixo(idioma) + nome
}

func qualPrefixo(idioma string) (prefixo string) {

	switch idioma {
	case "frances":
		prefixo = stringFrances
	case "espanhol":
		prefixo = stringEspanhol
	default:
		prefixo = stringPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
