package greet

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Sergio", "Hello, Sergio"},
		{"Fabian", "Hello, Fabian"},
		{"", "Hello, "},
	}

	for _, tt := range tests {
		got := Greet(tt.name)
		if got != tt.want {
			t.Errorf("Greet(%q) = %q; want %q", tt.name, got, tt.want)
		}
	}
}