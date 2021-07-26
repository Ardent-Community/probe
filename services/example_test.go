package services

import "testing"

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Errorf("Expected 3, got %d", add(1, 2))
	}
	if add(5, 6) != 11 {
		t.Errorf("Expected 11, got %d", add(5, 6))
	}
	if add(4, 98) != 102 {
		t.Errorf("Expected 102, got %d", add(4, 98))
	}
}
