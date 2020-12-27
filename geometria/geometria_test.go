package geometria

import "testing"

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
