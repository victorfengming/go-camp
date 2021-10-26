package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 27},
		{12, 35, 37},
		{30000, 40000, 52000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriang %d %d got %d ; exceped", tt.a, tt.b, tt.c)
		}
	}
}

/**
=== RUN   TestTriangle
    add_test.go:16: calcTriang 8 15 got 27 ; exceped
    add_test.go:16: calcTriang 30000 40000 got 52000 ; exceped
--- FAIL: TestTriangle (0.00s)

FAIL

Process finished with the exit code 1
*/
