package main

import "testing"

// func TestArea(t *testing.T) {

// verificaArea := func(t *testing.T, forma Forma, esperado float64) {
// 	t.Helper()
// 	resultado := forma.Area()

// 	if resultado != esperado {
// 		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
// 	}
// }

// t.Run("retângulos", func(t *testing.T) {
// 	retangulo := Retangulo{12.0, 6.0}
// 	verificaArea(t, retangulo, 72.0)
// })

// t.Run("círculos", func(t *testing.T) {
// 	circulo := Circulo{10}
// 	verificaArea(t, circulo, 314.1592653589793)
// })
// }

func TestArea(t *testing.T) {
	testesArea := []struct {
		forma   Forma
		temArea float64
	}{
		{forma: Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
		{forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.temArea {
			t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.temArea)
		}
	}
}
