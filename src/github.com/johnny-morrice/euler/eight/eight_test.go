package main

import (
    "testing"
)

func TestEuler8(t *testing.T) {
    if euler8(problem, 4) != 5832 {
        t.Error("euler8(problem, 4) != 5832")
    }
}