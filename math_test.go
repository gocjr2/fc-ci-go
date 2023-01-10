package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestDivisao(t *testing.T) {

	total := divisao(15, 15)

	if total != 1 {
		t.Errorf("Resultado da divisao é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}
