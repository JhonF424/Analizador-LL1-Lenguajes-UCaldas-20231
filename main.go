package main

import (
	"fmt"
	"strings"

	"github.com/JhonF424/LL1/controllers"
	"github.com/JhonF424/LL1/models"
)

func main() {

	/*gramatica := map[string][]string{
		"E":  {"T", "E'"},
		"E'": {"+", "T", "E'", "|", "λ"},
		"T":  {"F", "T'"},
		"T'": {"*", "F", "T'", "|", "λ"},
		"F":  {"(", "E", ")", "|", "id"},
	}
	*/

	/*gramatica := map[string][]string{ // Estos funcionan
		"A": {"B", "C"},
		"B": {"λ", "|", "m"},
		"C": {"λ", "|", "s"},
	}*/

	/*gramatica := map[string][]string{
		"E":  {"T", "EP"},
		"EP": {"+", "T", "EP", "|", "λ"},
		"T":  {"F", "TP"},
		"TP": {"*", "F", "TP", "|", "λ"},
		"F":  {"(", "E", ")", "|", "ident"},
	}*/

	/*gramatica := map[string][]string{
		"S": {"a", "B", "C", "d"},
		"B": {"C", "B", "b"},
		"C": {"c", "c", "lambda"},
	}*/

	/* gramatica := map[string]string{
		"S": "aBCd",
		"B": "CB|b",
		"C": "cc|λ",
	} */

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

	firsts := make(map[string][]string)
	firsts = controllers.Primeros(gramatica, firsts)

	fmt.Println("Primeros:")
	for nt, s := range firsts {
		fmt.Printf("%s: {%s}\n", nt, strings.Join(s, ", "))
	}

	follows := controllers.Siguientes(gramatica, firsts)

	fmt.Println("\nSiguientes:")
	for nt, s := range follows {
		fmt.Printf("%s: {%s}\n", nt, strings.Join(s, ", "))
	}

	solution := calculateSolutionSet(gramatica, firsts, follows)
	fmt.Println("\nConjunto Solución:")
	for k1, v1 := range solution {
		fmt.Print("\n" + k1 + ": {")
		for k2 := range v1 {
			fmt.Print(", ", k2)
		}
		fmt.Print("}")
	}

}

// Función para calcular el conjunto solución
func calculateSolutionSet(grammar []models.Grammar, firstSets map[string][]string, followSets map[string][]string) map[string]map[string]bool {
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
			if controllers.EsTerminal(prodSymbol) {
				solutionSet[prod.Symbol][prodSymbol] = true
				break
			}

			// Agregamos el primer conjunto de la producción al conjunto solución
			for _, first := range firstSets[prodSymbol] {
				solutionSet[prod.Symbol][first] = true
			}

			// Si el primer conjunto contiene lambda, agregamos el siguiente conjunto al conjunto solución
			if controllers.ContieneLambda(firstSets[prodSymbol]) {
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

/*
func CalculateSolutionSet(grammar []models.Grammar, firstSets map[string][]string, followSets map[string][]string) map[string][]string {
	solutionSets := make(map[string][]string)

	// Calcular el conjunto solución para cada producción
	for _, g := range grammar {
		solution := make([]string, 0)
		for _, p := range g.Productions {
			if controllers.ContieneLambda(firstSets[p]) {
				// Si la producción puede derivar epsilon, agregar los siguientes del símbolo no terminal
				follow := followSets[g.Symbol]
				solution = append(solution, follow...)
			} else {
				// Si la producción no puede derivar epsilon, agregar los primeros de la producción
				solution = append(solution, firstSets[p]...)
			}
		}
		solutionSets[g.Symbol] = solution
	}

	return solutionSets
}


func resolverRecursionIzquierda(gramatica map[string]string) map[string]string {
	nuevoNoTerminal := "X"                          // Nuevo símbolo no terminal para representar la recursión izquierda
	produccionesNuevas := make(map[string][]string) // Mapa para almacenar las nuevas producciones

	// Iterar sobre cada producción de la gramática
	for nt, prod := range gramatica {
		producciones := strings.Split(prod, "|")
		produccionesNuevas[nt] = make([]string, 0)
		recursiones := make([]string, 0)

		// Revisar si hay una recursión izquierda en la producción
		for _, p := range producciones {
			if strings.HasPrefix(p, nt) { // Si hay una recursión izquierda
				recursiones = append(recursiones, p)
			} else {
				produccionesNuevas[nt] = append(produccionesNuevas[nt], p)
			}
		}

		// Si hay recursiones izquierdas, generar nuevas producciones
		if len(recursiones) > 0 {
			nuevoNT := nt + nuevoNoTerminal // Nuevo símbolo no terminal para la recursión
			produccionesNuevas[nuevoNT] = make([]string, 0)

			for _, p := range producciones {
				if !strings.HasPrefix(p, nt) { // Si no hay recursión izquierda
					produccionesNuevas[nt] = append(produccionesNuevas[nt], p+nuevoNT)
				} else { // Si hay recursión izquierda
					produccionesNuevas[nuevoNT] = append(produccionesNuevas[nuevoNT], strings.TrimPrefix(p, nt)+nuevoNT)
				}
			}

			// Agregar producción epsilon al nuevo símbolo no terminal
			produccionesNuevas[nuevoNT] = append(produccionesNuevas[nuevoNT], "epsilon")
		}
	}

	// Actualizar la gramática con las nuevas producciones
	for nt, prod := range produccionesNuevas {
		gramatica[nt] = strings.Join(prod, "|")
	}

	return gramatica
}*/
