package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	masses, err := convertInputToNumeric(f)
	if err != nil {
		panic(err)
	}

	fuel, err := FuelRequirement(masses)
	if err != nil {
		panic(err)
	}

	fmt.Println(fuel)
}

func convertInputToNumeric(r io.Reader) ([]int, error) {
	var masses []int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		currentMass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("convert to int: %w", err)
		}

		masses = append(masses, currentMass)
	}

	return masses, nil
}

// FuelRequirement returns the total fuel requirement given
// a collection of masses
func FuelRequirement(masses []int) (int, error) {
	var totalFuel int

	for _, mass := range masses {
		for {
			mass = getFuelRequirement(mass)
			if mass <= 0 {
				break
			}

			totalFuel += mass
		}
	}

	return totalFuel, nil
}

func getFuelRequirement(mass int) int {
	return (mass / 3) - 2
}
