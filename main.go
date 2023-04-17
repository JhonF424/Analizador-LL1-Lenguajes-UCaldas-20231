package main

import (
	"fmt"
	"strings"

	"LL1/controllers"
	"LL1/models"
)

func main() {

	/*gramatica := []models.Grammar{
		"E":  {"T", "E'"},
		"E'": {"+", "T", "E'", "|", "λ"},
		"T":  {"F", "T'"},
		"T'": {"*", "F", "T'", "|", "λ"},
		"F":  {"(", "E", ")", "|", "id"},
	}
	*/

	/*gramatica := []models.Grammar{
		"A": {"B", "C"},
		"B": {"λ", "|", "m"},
		"C": {"λ", "|", "s"},
	}*/

	/*gramatica := []models.Grammar{
		"E":  {"T", "EP"},
		"EP": {"+", "T", "EP", "λ"},
		"T":  {"F", "TP"},
		"TP": {"*", "F", "TP", "λ"},
		"F":  {"(", "E", ")", "ident"},
	}*/

	/*gramatica := []models.Grammar{
		{Symbol: "S", Productions: []string{"a", "B", "C", "d"}},
		{Symbol: "B", Productions: []string{"B", "C", "B", "b"}},
		{Symbol: "C", Productions: []string{"c", "c", "lambda"}},
	}*/

	/*gramatica := []models.Grammar{
	{Symbol: "T", Productions: []string{"P", "m", "R"}},
	{Symbol: "T", Productions: []string{"P", "m", "D"}},
	{Symbol: "P", Productions: []string{"a", "m", "b"}},
	{Symbol: "P", Productions: []string{"a", "m", "d"}},
	{Symbol: "D", Productions: []string{"d"}},
	{Symbol: "R", Productions: []string{"r"}}}*/

	/*gramatica := []models.Grammar{
		{Symbol: "E", Productions: []string{"T", "T", "+", "E"}},
		{Symbol: "T", Productions: []string{"F", "F", "*", "T"}},
		{Symbol: "F", Productions: []string{"(", "E", ")", "id"}},
	}*/

	gramatica := []models.Grammar{

		{Symbol: "E", Productions: []string{"T", "EP"}},
		{Symbol: "EP", Productions: []string{"+", "T", "EP"}},
		{Symbol: "EP", Productions: []string{"-", "T", "EP"}},
		{Symbol: "EP", Productions: []string{"lambda"}},
		{Symbol: "T", Productions: []string{"F", "TP"}},
		{Symbol: "TP", Productions: []string{"*", "F", "TP"}},
		{Symbol: "TP", Productions: []string{"/", "F", "TP"}},
		{Symbol: "TP", Productions: []string{"lambda"}},
		{Symbol: "F", Productions: []string{"(", "E", ")"}},
		{Symbol: "F", Productions: []string{"num"}},
		{Symbol: "F", Productions: []string{"id"}}}

	//gsr := EliminarRecursionIzquierda(gramatica) 

	// Imprimir gramática factorizada
	fmt.Println("\nGramática sin recursión izquierda:")
	for _, g := range gramatica {
		fmt.Printf("%s -> %s\n", g.Symbol, strings.Join(g.Productions, " | "))
	}

	firsts := make(map[string][]string)
	firsts = controllers.Primeros(gramatica, firsts)

	fmt.Println("\nPrimeros:")
	for nt, s := range firsts {
		fmt.Printf("%s: {%s}\n", nt, strings.Join(s, ", "))
	}

	follows := controllers.Siguientes(gramatica, firsts)

	fmt.Println("\nSiguientes:")
	for nt, s := range follows {
		fmt.Printf("%s: {%s}\n", nt, strings.Join(s, ", "))
	}

	solution := controllers.CalculateSolutionSet(gramatica, firsts, follows)
	fmt.Println("\nConjunto Solución:")
	for k1, v1 := range solution {
		fmt.Print("\n" + k1 + ": {")
		for k2 := range v1 {
			fmt.Print(k2, ",")
		}
		fmt.Print("}")
	}
}

// EliminarRecursionIzquierda elimina la recursión izquierda de una gramática
func EliminarRecursionIzquierda(grammar []models.Grammar) []models.Grammar {
	newGrammar := make([]models.Grammar, len(grammar))

	// Copiar la gramática original en la nueva gramática
	for i, rule := range grammar {
		newGrammar[i] = models.Grammar{Symbol: rule.Symbol, Productions: make([]string, len(rule.Productions))}
		copy(newGrammar[i].Productions, rule.Productions)
	}

	// Para cada regla en la gramática
	for i, rule := range newGrammar {
		// Para cada símbolo en la producción de la regla
		for j, symbol := range rule.Productions {
			// Si hay una recursión izquierda
			if symbol == rule.Symbol {
				// Crear un nuevo símbolo para las producciones recursivas
				newSymbol := symbol + "'"

				// Crear una nueva regla para las producciones recursivas
				newRule := models.Grammar{Symbol: newSymbol, Productions: []string{}}

				// Añadir una producción lambda a la nueva regla
				newRule.Productions = append(newRule.Productions, "lambda")

				// Para cada producción en la regla original que tenga el símbolo recursivo
				for _, production := range rule.Productions {
					if strings.HasPrefix(production, symbol) {
						// Añadir la producción sin el símbolo recursivo y con el nuevo símbolo
						newProduction := strings.TrimPrefix(production, symbol)
						newProduction += newSymbol
						newRule.Productions = append(newRule.Productions, newProduction)
					}
				}

				// Eliminar las producciones recursivas de la regla original
				newGrammar[i].Productions = append(newGrammar[i].Productions[:j], newGrammar[i].Productions[j+1:]...)

				// Añadir la nueva regla a la gramática
				newGrammar = append(newGrammar, newRule)

				// Añadir el nuevo símbolo a la regla original y las otras reglas que lo contienen
				for k := range newGrammar {
					for l, production := range newGrammar[k].Productions {
						if strings.Contains(production, symbol) {
							newProduction := strings.Replace(production, symbol, newSymbol, -1)
							newGrammar[k].Productions[l] = newProduction
						}
					}
					if newGrammar[k].Symbol == symbol {
						newGrammar[k].Symbol = newSymbol
					}
				}

				// Salir del bucle para evitar procesar producciones adicionales de la regla
				break
			}
		}
	}

	return newGrammar
}
