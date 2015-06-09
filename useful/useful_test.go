package useful

import "testing"

func TestInvertir(t *testing.T) {
    output := Invertir("Hola")
    if output != "aloH" {
        t.Error("Expected /aloH, got ", output)
    }
}

func TestObtenerMaximo(t *testing.T) {
    output := ObtenerMaximo([]int{6, 3, 9, 4})
    if output != 9 {
        t.Error("Expected 9, got ", output)
    }
}