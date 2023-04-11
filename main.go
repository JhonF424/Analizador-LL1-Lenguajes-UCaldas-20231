package main

import (
	"fmt"
	"strings"

	"github.com/JhonF424/LL1/controllers"
	"github.com/JhonF424/LL1/models"
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
		"S": {"a", "B", "C", "d"},
		"B": {"C", "B", "b"},
		"C": {"c", "c", "lambda"},
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

/*
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
}
*/
