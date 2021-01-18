package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

const escrita = "escrita"
const pausa = "pausa"

const ultimaPalavra = "Go!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
