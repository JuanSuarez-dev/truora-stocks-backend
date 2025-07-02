package main

import (
	"math"
	"testing"
)

func TestParseDollar(t *testing.T) {
  cases := []struct {
    in   string
    want float64
  }{
    {"$4.20", 4.20},
    {"$1,234.56", 1234.56},
    {"$0.00", 0},
    {"invalid", math.NaN()},
  }

  for _, c := range cases {
    got := parseDollar(c.in)
    if math.IsNaN(c.want) {
      if !math.IsNaN(got) {
        t.Errorf("parseDollar(%q) = %v; want NaN", c.in, got)
      }
    } else if got != c.want {
      t.Errorf("parseDollar(%q) = %v; want %v", c.in, got, c.want)
    }
  }
}
