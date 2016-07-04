package main

import (
	"testing"
)

func TestEuler5A(t *testing.T) {
	actual := Euler5(10)
	if actual != 2520 {
		t.Error("Euler5(10) != 2520 was", actual)
	}
}