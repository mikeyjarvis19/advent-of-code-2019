package main

import "testing"

func TestCalculateFuelPart1(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		fuel_required := CalculateFuelPart1(table.mass)
		if fuel_required != table.fuel {
			t.Errorf("Calculated fuel required was incorrect, got: %d, want: %d.", fuel_required, table.fuel)
		}
	}
}

func TestCalculateFuelPart2(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, table := range tables {
		fuel_required := CalculateFuelPart2(table.mass)
		if fuel_required != table.fuel {
			t.Errorf("Calculated fuel required was incorrect, got: %d, want: %d.", fuel_required, table.fuel)
		}
	}
}
