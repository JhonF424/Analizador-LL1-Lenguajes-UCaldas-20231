package controllers

import "strings"

func esTerminal(s string) bool {
	return s != "λ" && s == strings.ToLower(s)
}
func esNoTerminal(s string) bool {
	return s != "λ" && s == strings.ToUpper(s)
}

func primeros(gramatica map[string][]string, prims map[string][]string) {
	for clave, valor := range gramatica {
		for _, val := range valor {
			// En caso de que sea un terminal lo agregamos al mapa de terminales
			if esTerminal(val) {
				prims[clave] = append(prims[clave], val)
				break
			}
			if esNoTerminal(val) {
				clave = val
			}
		}
	}
}
