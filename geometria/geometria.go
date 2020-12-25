package main

import "math"

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

// func Area(retangulo Retangulo) float64 {
// 	return retangulo.Largura * retangulo.Altura
// }

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

//Forma é uma interface; os testes chamam essa interface e seu método área, eles chegam aqui e acham esse "Area() float64", então procuram todos os metodos "Area" e que retorna um float64 (que estão acima por sinal) e executam eles.
type Forma interface {
	Area() float64
}
