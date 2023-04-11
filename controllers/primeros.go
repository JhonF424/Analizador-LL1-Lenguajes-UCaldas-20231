package controllers

import (
	"github.com/JhonF424/LL1/models"
)

func Primeros(grammar []models.Grammar, firsts map[string][]string) map[string][]string {
	for _, p := range grammar {
		if EsTerminal(p.Productions[0]) {
			AnnadirAlConjunto(firsts, p.Symbol, p.Productions[0])
		}
	}
	for changed := true; changed; {
		changed = false
		for _, p := range grammar {
			if len(p.Productions) > 0 && !EsTerminal(p.Productions[0]) {
				oldLen := len(firsts[p.Symbol])
				for _, s := range firsts[p.Productions[0]] {
					AnnadirAlConjunto(firsts, p.Symbol, s)
				}
				if len(firsts[p.Symbol]) > oldLen {
					changed = true
				}
			}
		}
	}

	return firsts
}
