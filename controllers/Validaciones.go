package controllers

func EsTerminal(symbol string) bool {
	return symbol == "lambda" || symbol[0] < 'A' || symbol[0] > 'Z'
}

func AnnadirAlConjunto(m map[string][]string, key string, values ...string) {
	for _, value := range values {
		if !Contiene(m[key], value) {
			m[key] = append(m[key], value)
		}
	}
}

func Contiene(slice []string, elemento string) bool {
	for _, elem := range slice {
		if elem == elemento {
			return true
		}
	}
	return false
}

func ContieneLambda(s []string) bool {
	return Contiene(s, "lambda")
}
