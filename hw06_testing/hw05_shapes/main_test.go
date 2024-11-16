package main

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		shape any
		want  float64
	}{
		{ // Circle
			shape: Circle{5},
			want:  78.53981633974483,
		},
		{ // Rectangle
			shape: Rectangle{5, 10},
			want:  50,
		},
		{ // Triangle
			shape: Triangle{8, 6},
			want:  24,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result, err := calculateArea(tt.shape)
			if err != nil {
				t.Error(err)
			}
			if result != tt.want {
				t.Errorf(" want %v, got %v", tt.want, result)
			}
		})
	}
}
