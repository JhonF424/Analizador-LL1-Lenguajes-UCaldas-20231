package main

import (
	"fmt"
	"strings"
)

type Gramatica struct {
	produccionInicial string
	producciones      map[string][][]string
}

func esTerminal(s string) bool {
	return s != "λ" && s == strings.ToLower(s)
}

func esNoTerminal(s string) bool {
	return s != "λ" && s == strings.ToUpper(s)
}

func contiene(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func puedeProducirEpsilon(primeros []string) bool {
	return contiene(primeros, "λ")
}

func (g *Gramatica) calcularPrimeros() map[string][]string {
	primeros := make(map[string][]string)

	// Paso 1
	for noTerminal := range g.producciones {
		primeros[noTerminal] = []string{}
	}

	// Paso 2
	primeros[g.produccionInicial] = append(primeros[g.produccionInicial], "λ")

	// Paso 3
	for cambio := true; cambio; {
		cambio = false
		for noTerminal, producciones := range g.producciones {
			for _, produccion := range producciones {
				for _, simbolo := range produccion {
					// Paso 3.1
					if esTerminal(simbolo) {
						if !contiene(primeros[noTerminal], simbolo) {
							primeros[noTerminal] = append(primeros[noTerminal], simbolo)
							cambio = true
						}
						break
					}
					// Paso 3.2
					if esNoTerminal(simbolo) {
						for _, primer := range primeros[simbolo] {
							if primer != "λ" && !contiene(primeros[noTerminal], primer) {
								primeros[noTerminal] = append(primeros[noTerminal], primer)
								cambio = true
							}
						}
						if !puedeProducirEpsilon(primeros[simbolo]) {
							break
						}
					}
					// Paso 3.3
					if simbolo == "λ" {
						if !contiene(primeros[noTerminal], "λ") {
							primeros[noTerminal] = append(primeros[noTerminal], "λ")
							cambio = true
						}
						if !puedeProducirEpsilon(primeros[simbolo]) {
							break
						}
					}
					// Paso 3.4
					if !puedeProducirEpsilon(primeros[simbolo]) {
						break
					}
				}
			}
		}
	}

	// Paso 4
	return primeros
}

func (g *Gramatica) calcularSiguientes() map[string][]string {
	siguientes := make(map[string][]string)
	for noTerminal := range g.producciones {
		siguientes[noTerminal] = []string{}
	}
	siguientes[g.produccionInicial] = []string{"$"}

	for {
		cambios := false
		for noTerminal, producciones := range g.producciones {
			for _, produccion := range producciones {
				for i, simbolo := range produccion {
					if esNoTerminal(simbolo) {
						siguientePosibles := []string{}
						if i < len(produccion)-1 {
							siguientePosibles = primeros(produccion[i+1:])
							if puedeProducirEpsilon(siguientePosibles) {
								siguientePosibles = append(siguientePosibles, siguientes[noTerminal]...)
							}
						} else {
							siguientePosibles = siguientes[noTerminal]
						}
						nuevosSiguientes := diferencia(siguientes[simbolo], siguientePosibles)
						if len(nuevosSiguientes) > 0 {
							siguientes[simbolo] = append(siguientes[simbolo], nuevosSiguientes...)
							cambios = true
						}
					}
				}
			}
		}
		if !cambios {
			break
		}
	}
	return siguientes
}

func simboloEnMinuscula(simbolo string) string {
	/* Aquí se puede implementar la lógica para convertir un símbolo a su forma en minúscula,
	en función de las reglas de la gramática*/
	return strings.ToLower(simbolo)
}

func diferencia(s1 []string, s2 []string) []string {
	diferencia := []string{}
	for _, e1 := range s1 {
		encontrado := false
		for _, e2 := range s2 {
			if e1 == e2 {
				encontrado = true
				break
			}
		}
		if !encontrado {
			diferencia = append(diferencia, e1)
		}
	}
	return diferencia
}

func main() {
	gramatica := Gramatica{
		produccionInicial: "LE",
		producciones: map[string][][]string{
			"LE": {{"R", "F", "|", "E", "lamda"}},
			"E":  {{"s", "*", "|", "l", "R", "LE", "s"}},
			"F":  {{"4", "|", "6", "R", "|", "t", "E", "lamda"}},
			"R":  {{"i", "|", "E"}},
		},
	}
	primeros := gramatica.calcularPrimeros()
	for noTerminal, simbolos := range primeros {
		fmt.Printf("Primeros(%s) = {%s}\n", noTerminal, strings.Join(simbolos, ", "))
	}
}

/*package main

import (
	"fmt"
	"unicode"
)

func main() {
	gramatica := map[string][]string{
		"LE": {"R", "F", "|", "E", "lambda"},
		"E":  {"s", "*", "|", "l", "R", "LE", "s"},
		"F":  {"4", "|", "6", "R", "|", "t", "E", "lambda"},
		"R":  {"i", "|", "E"},
	}

	fmt.Println(gramatica)

	for clave, valores := range gramatica {

		primerValor := valores[0]

		if unicode.IsUpper(rune(primerValor[0])) {
			fmt.Printf("%s: %s es una letra mayúscula\n", clave, primerValor)
		} else if unicode.IsLower(rune(primerValor[0])) {
			fmt.Printf("%s: %s es una letra minúscula\n", clave, primerValor)
		} else {
			fmt.Printf("%s: %s no es una letra\n", clave, primerValor)
		}
	}


}
*/
