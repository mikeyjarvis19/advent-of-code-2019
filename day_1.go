package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFileToArray(filepath string) []string {
	var mass_values []string
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		mass_values = append(mass_values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return mass_values
}

func CalculateFuelPart1(mass int) int {
	var fuel_required int = (mass / 3) - 2
	return fuel_required
}

func CalculateFuelPart2(mass int) int {
	var total_fuel_required = 0
	var fuel_required int = CalculateFuelPart1(mass)
	total_fuel_required += fuel_required
	for fuel_required > 0 {
		fuel_required = CalculateFuelPart1(fuel_required)
		// fmt.Println("Fuel required for fuel: ", fuel_required)
		if fuel_required > 0 {
			total_fuel_required += fuel_required
		}
	}
	return total_fuel_required
}

func part_1() {
	fmt.Println("Part 1 answer:")
	var mass_values = ReadFileToArray("inputs/day_1.txt")
	var total_fuel_required int = 0
	for _, element := range mass_values {
		i, _ := strconv.Atoi(element)
		var fuel_for_trip = CalculateFuelPart1(i)
		total_fuel_required += fuel_for_trip
	}
	fmt.Println(total_fuel_required)
}

func part_2() {
	fmt.Println("Part 2 answer:")
	var mass_values = ReadFileToArray("inputs/day_1.txt")
	var total_fuel_required int = 0
	for _, element := range mass_values {
		i, _ := strconv.Atoi(element)
		var fuel_for_trip = CalculateFuelPart2(i)
		total_fuel_required += fuel_for_trip
	}
	fmt.Println(total_fuel_required)
}
func main() {
	part_1()
	part_2()
}
