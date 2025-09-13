package main

import (
	"testing"
)

func TestSumar(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, tt := range tests {
		got := sum(tt.a, tt.b)
		if got != tt.expect {
			t.Errorf("Sumar(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestRestar(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{5, 2, 3},
		{1, -1, 2},
		{0, 0, 0},
	}

	for _, tt := range tests {
		got := sub(tt.a, tt.b)
		if got != tt.expect {
			t.Errorf("Restar(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestMultiplicar(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{2, 3, 6},
		{-1, 5, -5},
		{0, 10, 0},
	}

	for _, tt := range tests {
		got := mult(tt.a, tt.b)
		if got != tt.expect {
			t.Errorf("Multiplicar(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestDividir(t *testing.T) {
	tests := []struct {
		a, b    float64
		expect  float64
		wantErr bool
	}{
		{10, 2, 5.0, false},
		{5, -1, -5.0, false},
		{1, 0, 0, true}, // expect error
	}

	for _, tt := range tests {
		got, err := div(tt.a, tt.b)
		if tt.wantErr {
			if err == nil {
				t.Errorf("Dividir(%v, %v) expected error, got nil", tt.a, tt.b)
			}
		} else {
			if err != nil {
				t.Errorf("Dividir(%v, %v) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.expect {
				t.Errorf("Dividir(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expect)
			}
		}
	}
}
