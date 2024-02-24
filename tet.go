package main

import (
    "fmt"
    "math/cmplx"
)

// LambertW approximates the Lambert W function using Newton's method.
// It returns the principal branch W(z) for real z >= -1/e and complex z.
func LambertW(z complex128) complex128 {
    const (
        epsilon = 1e-13 // precision
        maxIter = 10000  // maximum iterations
    )

    // Initial guess for Newton's method
    var w complex128
    if real(z) < 0 {
        w = -1
    } else if real(z) == 0 {
        w = 0
    } else {
        w = 1
    }

    // Apply Newton's method
    for i := 0; i < maxIter; i++ {
        wNew := w - (w*cmplx.Exp(w)-z)/(w*(1+w))
        if cmplx.Abs(wNew-w) < epsilon {
            return wNew
        }
        w = wNew
    }

    return w
}

// Tetration computes the tetration of a complex base with a real exponent.
func Tetration(base complex128, z float64) complex128 {
    // Use the Lambert W function to compute tetration
    w := LambertW(cmplx.Log(base))
    result := w
    for n := 2; n <= int(z); n++ {
        result = w * cmplx.Exp(w*(result-cmplx.Log(w)))
    }
    return result
}

func main() {
    // Test the Tetration function
    base := complex(2, 0.01) // Base of the tetration
    exponent := 3.0       // Exponent of the tetration
    result := Tetration(base, exponent)
    fmt.Println("Tetration result:", result)
}
