package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Action int

const (
	Invalid Action = iota
	North
	South
	East
	West
	Left
	Right
	Forward
)

func NewAction(actionCode byte) Action {
	lookup := map[byte]Action{
		'N': North,
		'S': South,
		'E': East,
		'W': West,
		'L': Left,
		'R': Right,
		'F': Forward,
	}

	action, ok := lookup[actionCode]
	if ok {
		return action
	}

	return Invalid
}

func (a Action) String() string {
	return [...]string{"I", "N", "S", "E", "W", "L", "R", "F"}[a]
}

type Instruction struct {
	action Action
	value int
}

func NewInstruction(actionCode byte, value int) *Instruction {
	action := NewAction(actionCode)
	if action == Invalid {
		return nil
	}

	return &Instruction{
		action: action,
		value: value,
	}
}

func (i Instruction) String() string {
	return i.action.String() + strconv.Itoa(i.value)
}

func loadInstructionsFromFile(path string) (instructions []Instruction, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var actionCode byte
		var value int
		_, err := fmt.Sscanf(scanner.Text(), "%c%d", &actionCode, &value)
		if err != nil {
			return nil, err
		}

		instruction := NewInstruction(actionCode, value)
		if instruction == nil {
			return  nil, fmt.Errorf("failed to create instruction %c %d", actionCode, value)
		}

		instructions = append(instructions, *instruction)
	}
	return instructions, scanner.Err()
}

type Vector2d struct {
	x int
	y int
}

func (t Vector2d) ManhattanDistance() int {
	return absInt(t.x) + absInt(t.y)
}

func (t Vector2d) Add(v Vector2d) Vector2d {
	return Vector2d{
		x: t.x + v.x,
		y: t.y + v.y,
	}
}

func (t Vector2d) Multiply(v int) Vector2d {
	return Vector2d{
		x: t.x * v,
		y: t.y * v,
	}
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}



type Direction int

const (
	N Direction = iota
	E
	S
	W
)

var(
	directionVectors = [...]Vector2d{
		{x: 0, y: 1},   // North
		{ x: 1, y: 0},  // East
		{ x: 0, y: -1}, // South
		{ x: -1, y: 0}, // West
	}
)

type Ship struct {
	position Vector2d
	facing Direction
}

func NewShip() *Ship {
	return &Ship{
		position: Vector2d{0, 0},
		facing: E,
	}
}

func (s *Ship) Move(dir Direction, distance int) {
	travel := directionVectors[int(dir) % len(directionVectors)]
	s.position = s.position.Add(travel.Multiply(distance))
}

func (s *Ship) Turn(angle int) {
	numPoints := angle / 90
	dir := (len(directionVectors) + int(s.facing) + numPoints) % len(directionVectors)
	s.facing = Direction(absInt(dir))
}

func (s *Ship) HandleInstruction(i Instruction) {
	switch i.action {
	case North:
		s.Move(N, i.value)
		break
	case South:
		s.Move(S, i.value)
		break
	case East:
		s.Move(E, i.value)
		break
	case West:
		s.Move(W, i.value)
		break
	case Left:
		s.Turn(-i.value)
		break
	case Right:
		s.Turn(i.value)
		break
	case Forward:
		s.Move(s.facing, i.value)
		break
	}
}

func main() {
	// Load the file
	instructions, err := loadInstructionsFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	ship := NewShip()
	for _, inst := range instructions {
		ship.HandleInstruction(inst)
	}

	fmt.Printf("Ship Position: (%d, %d) Facing: %d ManhattanDistance: %d\n", ship.position.x, ship.position.y, ship.facing, ship.position.ManhattanDistance())
}