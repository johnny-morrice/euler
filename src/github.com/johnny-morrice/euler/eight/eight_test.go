package main

import (
    "testing"
)

func TestEuler8(t *testing.T) {
    if euler8(4) != 5832 {
        t.Error("euler8(problem, 4) != 5832")
    }
}

// Solve Project Euler problem 8
func euler8(factcount int) int64 {
    m := factorMatrix(problem, factcount)
    return max(m, factcount - 1)
}