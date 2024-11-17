package main

import "testing"

func TestBook_GetID(t *testing.T) {
	tests := []struct {
		mode  Select
		book1 Book
		book2 Book
		want  bool
	}{
		{
			// By Year - book1
			mode:  ByYear,
			book1: Book{year: 1859},
			book2: Book{year: 1847},
			want:  true,
		},
		{
			// By Year - book2
			mode:  ByYear,
			book1: Book{year: 1847},
			book2: Book{year: 1859},
			want:  false,
		},
		{
			// By Size - book1
			mode:  BySize,
			book1: Book{size: 333},
			book2: Book{size: 222},
			want:  true,
		},
		{
			// By Size - book2
			mode:  BySize,
			book1: Book{size: 222},
			book2: Book{size: 333},
			want:  false,
		},
		{
			// By Rate - book1
			mode:  ByRate,
			book1: Book{rate: 3.6},
			book2: Book{rate: 3.5},
			want:  true,
		},
		{
			// By Rate - book2
			mode:  ByRate,
			book1: Book{rate: 3.5},
			book2: Book{rate: 3.6},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			sravn := NewSravn(tt.mode)
			result := sravn.Compare(&tt.book1, &tt.book2)
			if result != tt.want {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}
