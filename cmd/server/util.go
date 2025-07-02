package main

import (
	"math"
	"strconv"
	"strings"
)

// parseDollar quita “$” y coma de miles, convierte a float, o devuelve NaN.
func parseDollar(s string) float64 {
	clean := strings.ReplaceAll(strings.ReplaceAll(s, "$", ""), ",", "")
	v, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return math.NaN()
	}
	return v
}
