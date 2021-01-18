package main

import (
	"fmt"
	"io"
	"net/http"
)

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "<div style='background-color: blue; width: 100; height: 100'>olá</div>")
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "mundo")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		fmt.Println(err)
	}
}
