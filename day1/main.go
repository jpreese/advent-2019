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

	totalFuel, err := calculateTotalFuelRequirement(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(totalFuel)

	return
}

func calculateTotalFuelRequirement(r io.Reader) (int, error) {
	var totalFuel int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		currentMass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("string to int: %w", err)
		}

		for {
			currentMass = getFuelRequirement(currentMass)
			if currentMass <= 0 {
				break
			}

			totalFuel += currentMass
		}
	}

	return totalFuel, nil
}

func getFuelRequirement(mass int) int {
	return (mass / 3) - 2
}
