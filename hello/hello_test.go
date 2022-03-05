package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions test in action-demo" {
		t.Errorf("Greet() = %s; Expected Hello GitHub Actions test in action-demo", result)
	}

}
