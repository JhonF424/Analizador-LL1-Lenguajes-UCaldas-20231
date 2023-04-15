package main

import (
	"fmt"
	"reflect"
	"strconv"
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
		{Symbol: "S", Productions: []string{"a", "B", "C", "d"}},
		{Symbol: "B", Productions: []string{"C", "B", "b"}},
		{Symbol: "C", Productions: []string{"c", "c", "lambda"}},
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

func Factorizar(grammar []models.Grammar) []models.Grammar {
	for i := 0; i < len(grammar); i++ {
		for j := 0; j < len(grammar[i].Productions); j++ {
			for k := j + 1; k < len(grammar[i].Productions); k++ {
				if grammar[i].Productions[j][0] == grammar[i].Productions[k][0] {
					temp := commonPrefix(grammar[i].Productions[j], grammar[i].Productions[k])
					if temp != "" {
						nonTerminal := fmt.Sprintf("%s'", temp)
						grammar = addGrammar(grammar, nonTerminal, []string{temp + string(nonTerminal[1])})
						grammar[i].Productions[j] = strings.Replace(grammar[i].Productions[j], temp, nonTerminal, 1)
						grammar[i].Productions[k] = strings.Replace(grammar[i].Productions[k], temp, nonTerminal, 1)
						grammar = removeGrammar(grammar, models.Grammar{Symbol: grammar[i].Symbol, Productions: []string{}})
						i = 0
						break
					}
				}
			}
		}
	}
	return grammar
}

// EliminarRecursionIzquierda elimina la recursión izquierda de una gramática
func EliminarRecursionIzquierda(grammar []models.Grammar) []models.Grammar {
	for i := range grammar {
		symbol := grammar[i].Symbol
		productions := grammar[i].Productions
		newProductions := make([]string, 0)
		nonTerminals := make([]string, 0)

		// separar producciones recursivas y no recursivas
		for _, production := range productions {
			if strings.HasPrefix(production, symbol) {
				nonTerminals = append(nonTerminals, production)
			} else {
				newProductions = append(newProductions, production)
			}
		}

		if len(nonTerminals) == 0 {
			continue
		}

		// construir nuevos símbolos no terminales
		newSymbols := make([]string, len(nonTerminals))
		for j := range nonTerminals {
			newSymbol := symbol + "'" + strconv.Itoa(j)
			newSymbols[j] = newSymbol
		}

		// crear nuevas producciones sin recursión
		for _, production := range newProductions {
			grammar = addGrammar(grammar, symbol, []string{production})
		}

		// crear nuevas producciones con los nuevos símbolos no terminales
		for j, nonTerminal := range nonTerminals {
			productions := strings.Split(nonTerminal[len(symbol):], "|")
			newProductions := make([]string, len(productions))
			for k, production := range productions {
				if strings.HasPrefix(production, symbol) {
					production = strings.Replace(production, symbol, newSymbols[j], 1)
				}
				newProductions[k] = production + newSymbols[j]
			}
			grammar = addGrammar(grammar, newSymbols[j], newProductions)
		}

		// crear producciones con los nuevos símbolos no terminales para las producciones restantes
		for _, nonTerminal := range nonTerminals {
			productions := strings.Split(nonTerminal[len(symbol):], "|")
			newProductions := make([]string, len(productions))
			for k, production := range productions {
				newProductions[k] = production + newSymbols[0]
			}
			grammar = addGrammar(grammar, symbol, newProductions)
		}

		// eliminar producción original
		grammar = removeGrammar(grammar, symbol, productions)
	}

	return grammar
}

// Agrega una producción a la gramática
func addGrammar(grammar []models.Grammar, symbol string, production []string) []models.Grammar {
	for i := range grammar {
		if grammar[i].Symbol == symbol {
			grammar[i].Productions = append(grammar[i].Productions, production...)
			return grammar
		}
	}

	grammar = append(grammar, models.Grammar{Symbol: symbol, Productions: production})
	return grammar
}

// Remueve una producción de la gramática
func removeGrammar(grammar []models.Grammar, symbol string, production []string) []models.Grammar {
	for i := range grammar {
		if grammar[i].Symbol == symbol {
			grammar[i].Productions = remove(grammar[i].Productions, production)
			return grammar
		}
	}

	return grammar
}

// Encuentra el prefijo común más largo entre dos strings
func commonPrefix(s1, s2 string) string {
	var prefix string
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			break
		}
		prefix += string(s1[i])
	}
	return prefix
}

func remove(slice []string, s []string) []string {
	for i, v := range slice {
		if reflect.DeepEqual(v, s) {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
