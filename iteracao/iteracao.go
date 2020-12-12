package iteracao

// Repetir é util
func Repetir(caractere string, quantasVezes int) string {
	var repeticoes string
	for i := 0; i < quantasVezes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}
