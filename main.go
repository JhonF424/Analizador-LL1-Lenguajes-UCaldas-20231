package main

import (
	"fmt"
)

func main() {
	/*gramatica := map[string][]string{
		"E":  {"T", "E'"},
		"E'": {"+", "T", "E'", "|", "λ"},
		"T":  {"F", "T'"},
		"T'": {"*", "F", "T'", "|", "λ"},
		"F":  {"(", "E", ")", "|", "id"},
	}

	gramatica := map[string][]string{
		"A": {"B", "C"},
		"B": {"λ", "|", "m"},
		"C": {"λ", "|", "s"},
	}*/

	gramatica := map[string][]string{
		"E":  {"T", "EP"},
		"EP": {"+", "T", "EP", "|", "λ"},
		"T":  {"F", "TP"},
		"TP": {"*", "F", "TP", "|", "λ"},
		"F":  {"(", "E", ")", "|", "ident"},
	}

	// Se calculan los primeros de la gramática
	primeros := calcularPrimeros(gramatica)

	// Se imprimen los primeros de cada no terminal
	for noTerminal, conjuntoPrimeros := range primeros {
		fmt.Printf("Primeros de %s: %v\n", noTerminal, conjuntoPrimeros)
	}
}

func calcularPrimeros(gramatica map[string][]string) map[string][]string {
	primeros := make(map[string][]string)
	visitados := make(map[string]bool)

	// Función recursiva para calcular los primeros de cada no terminal
	var calcularPrimerosRec func(string) []string
	calcularPrimerosRec = func(noTerminal string) []string {
		// Si ya se visitó este no terminal, se retorna su conjunto de primeros
		if visitados[noTerminal] {
			return primeros[noTerminal]
		}
		visitados[noTerminal] = true

		// Se recorre cada producción del no terminal
		for _, produccion := range gramatica[noTerminal] {
			// Si la producción empieza con un terminal, se agrega a los primeros del no terminal actual
			if esTerminal(produccion[0]) {
				primeros[noTerminal] = append(primeros[noTerminal], string(produccion[0]))
			}
			// Si la producción empieza con un no terminal, se calculan sus primeros y se agregan a los del no terminal actual
			if esNoTerminal(produccion[0]) {
				primeros[noTerminal] = append(primeros[noTerminal], calcularPrimerosRec(string(produccion[0]))...)
			}
			// Si la producción empieza con lambda, se agrega a los primeros del no terminal actual
			if produccion == "λ" {
				primeros[noTerminal] = append(primeros[noTerminal], "λ")
			}
		}
		primeros[noTerminal] = eliminarDuplicados(primeros[noTerminal])
		return primeros[noTerminal]
	}

	// Se calculan los primeros de cada no terminal
	for noTerminal := range gramatica {
		calcularPrimerosRec(noTerminal)
	}

	return primeros
}

func eliminarDuplicados(s []string) []string {
	set := make(map[string]bool)
	var result []string
	for _, item := range s {
		if !set[item] {
			set[item] = true
			result = append(result, item)
		}
	}
	return result
}

func esTerminal(simbolo byte) bool {
	return simbolo >= 'a' && simbolo <= 'z'
}

func esNoTerminal(simbolo byte) bool {
	return simbolo >= 'A' && simbolo <= 'Z'
}

/*
func esTerminal(s string) bool {
	return s != "λ" && s == strings.ToLower(s)
}

func esNoTerminal(s string) bool {
	return s != "λ" && s == strings.ToUpper(s)
}

func main() {
	/*gramatica1 := map[string][]string{
		"LE": {"R", "F", "|", "E", "λ"},
		"E":  {"s", "*", "|", "l", "R", "LE", "s"},
		"F":  {"4", "|", "6", "R", "|", "t", "E", "λ"},
		"R":  {"i", "|", "E"},
	}/

	gramatica2 := map[string][]string{
		"S": {"a", "B", "C", "d"},
		"B": {"C", "B", "|", "b"},
		"C": {"c", "c", "|", "λ"},
	}

	primeros := make(map[string][]string)

	fmt.Println(gramatica2)

	for key, value := range primeros {
		fmt.Println(key, value)
	}

}

func esTerminal(simbolo byte) bool {
	return simbolo >= 'a' && simbolo <= 'z'
}

func esNoTerminal(simbolo byte) bool {
	return simbolo >= 'A' && simbolo <= 'Z'
}

func generaLambda(primeros []string) bool {
	for _, primer := range primeros {
		if primer != "λ" {
			return false
		}
	}
	return true
}

/*
func primeros(gramatica map[string][]string, primeros map[string][]string) {
	for clave, valor := range gramatica {

		for _, val := range valor {
			// En caso de que sea un terminal lo agregamos al mapa de terminales
			if esTerminal(val) {
				primeros[clave] = append(primeros[clave], val)
				break
			}

			if esNoTerminal(val) {
				primeros()
			}
		}
	}
}
*/
