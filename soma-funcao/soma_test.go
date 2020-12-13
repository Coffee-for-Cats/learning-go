package main

import (
	"reflect"
	"testing"
)

func verificaSomas(t *testing.T, resultado, esperado []int) {
	t.Helper()
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v esperado %v", resultado, esperado)
	}
}

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})
}

func TestSomaTudo(t *testing.T) {
	recebido := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(recebido, esperado) {
		t.Errorf("resultado %v esperado %v", recebido, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {

	t.Run("faz a soma do resto", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		verificaSomas(t, resultado, esperado)
	})
}
