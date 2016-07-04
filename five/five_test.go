package five

import (
	"testing"
)

func TestEuler5A(t *testing.T) {
	if Euler5A(10) != 2520 {
		t.Error("Euler5(10) != 2520")
	}

	if Euler5B(10) != 2520 {
		t.Error("Euler5B(10) != 2520")
	}
}