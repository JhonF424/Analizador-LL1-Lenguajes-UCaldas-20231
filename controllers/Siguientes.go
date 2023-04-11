package controllers

import "github.com/JhonF424/LL1/models"

func Siguientes(grammar []models.Grammar, primeros map[string][]string) map[string][]string {
	siguientes := make(map[string][]string)

	// Inicializar el conjunto de siguientes de cada símbolo no terminal como un conjunto vacío.
	for _, prod := range grammar {
		if !EsTerminal(prod.Symbol) {
			siguientes[prod.Symbol] = []string{}
		}
	}

	// Agregar $ al conjunto de siguientes de la producción inicial.
	siguientes[grammar[0].Symbol] = []string{"$"}

	for {
		// Variable para indicar si se han agregado nuevos símbolos a los conjuntos de siguientes.
		agregados := false

		// Para cada producción, agregar los primeros de los símbolos que vienen después de los no terminales
		for _, prod := range grammar {
			for i, sym := range prod.Productions {
				if !EsTerminal(sym) {
					// Si el no terminal es el último símbolo de la producción, agregar el siguiente del símbolo que contiene la producción.
					if i == len(prod.Productions)-1 {
						for _, s := range siguientes[prod.Symbol] {
							if !Contiene(siguientes[sym], s) {
								siguientes[sym] = append(siguientes[sym], s)
								agregados = true
							}
						}
					} else {
						// Si el no terminal no es el último símbolo de la producción, agregar los primeros de los símbolos que vienen después.
						for _, s := range primeros[prod.Productions[i+1]] {
							if !Contiene(siguientes[sym], s) {
								siguientes[sym] = append(siguientes[sym], s)
								agregados = true
							}
						}
						// Si los siguientes de los símbolos que vienen después contienen lambda, agregar los siguientes del no terminal.
						if Contiene(primeros[prod.Productions[i+1]], "lambda") {
							for _, s := range siguientes[prod.Symbol] {
								if !Contiene(siguientes[sym], s) {
									siguientes[sym] = append(siguientes[sym], s)
									agregados = true
								}
							}
						}
					}
				}
			}
		}

		// Si no se agregaron nuevos símbolos a los conjuntos de siguientes, terminar.
		if !agregados {
			break
		}
	}

	return siguientes
}
