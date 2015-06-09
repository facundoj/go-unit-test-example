package useful

import "math"

// Invierte la cadena recibida
func Invertir(input string) string {
    inverted := make([]byte, len(input))

    for i := 0; i < len(input); i++ {
        inverted[i] = input[len(input) - 1 - i]
    }

    return string(inverted)
}

// Determina el mayor entero
func ObtenerMaximo(lista []int) int {
    max := math.MinInt8
    for _, el := range lista {
        if el > max {
            max = el
        }
    }
    return max
}