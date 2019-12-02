package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	startingInstructions := "1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,9,23,1,5,23,27,1,27,9,31,1,6,31,35,2,35,9,39,1,39,6,43,2,9,43,47,1,47,6,51,2,51,9,55,1,5,55,59,2,59,6,63,1,9,63,67,1,67,10,71,1,71,13,75,2,13,75,79,1,6,79,83,2,9,83,87,1,87,6,91,2,10,91,95,2,13,95,99,1,9,99,103,1,5,103,107,2,9,107,111,1,111,5,115,1,115,5,119,1,10,119,123,1,13,123,127,1,2,127,131,1,131,13,0,99,2,14,0,0"
	instructions, err := convertInstructionsToInt(startingInstructions)
	if err != nil {
		panic(err)
	}

	const answer = 19690720
	const maxSize = 0
	for noun := 0; noun < maxSize; noun++ {
		for verb := 0; verb < maxSize; verb++ {
			tempInstructions := make([]int, len(instructions))
			copy(tempInstructions, instructions)

			tempInstructions[1] = noun
			tempInstructions[2] = verb

			if getProgramOutput(tempInstructions) == answer {
				result := 100*noun + verb
				fmt.Println(result)
			}
		}
	}
}

func convertInstructionsToInt(instructions string) ([]int, error) {
	positions := strings.Split(instructions, ",")
	var result []int

	for _, v := range positions {
		position, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("convert to int: %w", err)
		}

		result = append(result, position)
	}

	return result, nil
}

func getProgramOutput(instructions []int) int {
	var index int
	var posOne int
	var posTwo int
	var posThree int

	const addCode = 1
	const multiplyCode = 2
	const haltCode = 99

	for instructions[index] != haltCode {
		posOne = instructions[index+1]
		posTwo = instructions[index+2]
		posThree = instructions[index+3]

		if instructions[index] == addCode {
			instructions[posThree] = instructions[posOne] + instructions[posTwo]
		} else if instructions[index] == multiplyCode {
			instructions[posThree] = instructions[posOne] * instructions[posTwo]
		}

		index += 4
	}

	return instructions[0]
}
