package main

import (
	"bufio"
	"fmt"
	"os"
)

type OpCode string

const (
	Nop OpCode = "nop"
	Acc        = "acc"
	Jmp        = "jmp"
)

type Instruction struct {
	operation OpCode
	argument  int
	processed bool
}

func loadInstructionsFromFile(path string) (instructions []Instruction, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var instruction Instruction
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &instruction.operation, &instruction.argument)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction)
	}
	return instructions, scanner.Err()
}

type Registers struct {
	ip  int
	acc int
}

func main() {
	// Load the file
	program, err := loadInstructionsFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	registers := Registers{}
	for !program[registers.ip].processed {

		ipDelta := 1

		switch program[registers.ip].operation {
		case Nop:
			break
		case Acc:
			registers.acc += program[registers.ip].argument
			break
		case Jmp:
			ipDelta = program[registers.ip].argument
			break
		}

		program[registers.ip].processed = true
		registers.ip += ipDelta
	}

	fmt.Printf("IP: %d ACC: %d\n", registers.ip, registers.acc)
}
