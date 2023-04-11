package controllers

import "github.com/JhonF424/LL1/models"

func CalculateSolutionSet(grammar []models.Grammar, firstSets map[string][]string, followSets map[string][]string) map[string]map[string]bool {
	// Declaración del conjunto solución
	solutionSet := make(map[string]map[string]bool)

	// Recorremos todas las producciones
	for _, prod := range grammar {
		// Si la producción no está en el conjunto solución
		if _, ok := solutionSet[prod.Symbol]; !ok {
			// Lo agregamos al conjunto solución
			solutionSet[prod.Symbol] = make(map[string]bool)
		}

		// Si la producción es una cadena vacía
		if len(prod.Productions) == 1 && prod.Productions[0] == "lambda" {
			// Agregamos lambda al conjunto solución
			solutionSet[prod.Symbol]["lambda"] = true
			continue
		}

		// Recorremos las producciones de la producción
		for _, prodSymbol := range prod.Productions {
			// Si el símbolo es un terminal, lo agregamos al conjunto solución
			if EsTerminal(prodSymbol) {
				solutionSet[prod.Symbol][prodSymbol] = true
				break
			}

			// Agregamos el primer conjunto de la producción al conjunto solución
			for _, first := range firstSets[prodSymbol] {
				solutionSet[prod.Symbol][first] = true
			}

			// Si el primer conjunto contiene lambda, agregamos el siguiente conjunto al conjunto solución
			if ContieneLambda(firstSets[prodSymbol]) {
				for _, follow := range followSets[prod.Symbol] {
					solutionSet[prod.Symbol][follow] = true
				}
			} else {
				break
			}
		}
	}

	// Devolvemos el conjunto solución
	return solutionSet
}
