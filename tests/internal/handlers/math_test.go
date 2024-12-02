package handlers

import (
	"testing"

	"boilerplate/internal/handlers"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "Add 1 + 1", a: 1, b: 1, expected: 2},
		{name: "Add 0 + 0", a: 0, b: 0, expected: 0},
		{name: "Add -1 + -1", a: -1, b: -1, expected: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		handlers.Add(1, 1)
	}
}
