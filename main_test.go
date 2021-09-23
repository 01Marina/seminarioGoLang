package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
		{"TX022B", false, "", "", 0},
		{"NN0%2B", false, "", "", 0},
		{"TG022B", false, "", "", 0},
		{"TX02%B", false, "", "", 0},
		{"T)02%B", false, "", "", 0},
		{"TX00", true, "TX", "", 0},
		{"NN00", true, "NN", "", 0},
		{"NN0", false, "", "", 0},
		{"TX0", false, "", "", 0},
		{"NN02", false, "", "", 0},
		{"TX02", false, "", "", 0},
		{"Ts02", false, "", "", 0},
		{"NN2", false, "", "", 0},
		{"NN", false, "", "", 0},
		{"TX", false, "", "", 0},
		{"x", false, "", "", 0},
		{"", false, "", "", 0},
	}
	for _, testData := range cases {
		r, err := ParseString([]byte(testData.Input))
		// acá agregar chequeos propios del test por ejemplo:
		assert.Equal(t, err == nil, testData.Success)
		assert.Equal(t, r.Length, testData.Length)
		assert.Equal(t, r.Type, testData.Type)
		assert.Equal(t, r.Value, testData.Value)
	}
}
