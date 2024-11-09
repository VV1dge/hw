package main

import (
	"fmt"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want string
	}{
		{1, 1, " n"},
		{2, 2, " #n# n"},
		{3, 3, " # n# #n # n"},
		{4, 4, " # #n# # n # #n# # n"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%dx%d", test.x, test.y), func(t *testing.T) {
			result := CreateBoard(test.x, test.y)
			if result != test.want {
				t.Errorf("got %s, want %s", test.want, result)
			}
		})
	}
}
