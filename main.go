package main

import (
	"fmt"
	"strings"
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
		"B": {"C", "B", "|", "b"},
		"C": {"c", "c", "|", "λ"},
	}*/

	gramatica := map[string]string{
		"S": "aBCd",
		"B": "CB|b",
		"C": "cc|λ",
	}

	p := map[string][]string{}
	probando(gramatica, p, "S")
	for k, v := range p {
		fmt.Println(k, v)
	}

}

func esTerminal(s string) bool {
	return s != "λ" && s != "|" && s == strings.ToLower(s)
}
func esNoTerminal(s string) bool {
	return s != "λ" && s == strings.ToUpper(s)
}

func contiene(slice []string, elemento string) bool {
	for _, elem := range slice {
		if elem == elemento {
			return true
		}
	}
	return false
}

func probando(gramatica map[string]string, prims map[string][]string, inicio string) {
	for k, v := range gramatica {
		if k == inicio {
			listaProds := strings.Split(v, "|")
			for _, palabra := range listaProds {
				for _, letra := range palabra {
					if esTerminal(string(letra)) {
						prims[k] = append(prims[k], string(letra))
						break
					}

					if esNoTerminal(string(letra)) {
						probando(gramatica, prims, string(letra))
					}
				}

			}
		}
	}
}

/*
func encontrarPrimeros(gramatica map[string]string, prims map[string][]string, inicio string) {
	if esNoTerminal(inicio) {
		// Agregar la clave de inicio a los primeros
		prims[inicio] = []string{}

		// Empezar desde la clave de inicio
		encontrarPrimerosRecursivo(gramatica, prims, inicio, inicio)
	} else {
		panic("El símbolo de inicio debe ser un no terminal")
	}
}

func encontrarPrimerosRecursivo(gramatica map[string]string, prims map[string][]string, simboloActual string, inicio string) {
	listaProds := strings.Split(gramatica[simboloActual], "|")

	for _, prod := range listaProds {
		for _, letra := range prod {
			if esTerminal(string(letra)) {
				prims[inicio] = append(prims[inicio], string(letra))
				break
			}

			if esNoTerminal(string(letra)) {
				encontrarPrimerosRecursivo(gramatica, prims, string(letra), inicio)
			}
		}
	}
}
*/

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
