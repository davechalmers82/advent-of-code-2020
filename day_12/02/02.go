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

func (t Vector2d) RotateClockwise90() Vector2d {
	return Vector2d{
		x: t.y,
		y: -t.x,
	}
}

func (t Vector2d) RotateAntiClockwise90() Vector2d {
	return Vector2d{
		x: -t.y,
		y: t.x,
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
	waypoint Vector2d
}

func NewShip() *Ship {
	return &Ship{
		position: Vector2d{0, 0},
		waypoint: Vector2d{10, 1},
	}
}

func (s *Ship) MoveToWaypoint(times int) {
	s.position = s.position.Add(s.waypoint.Multiply(times))
}

func (s *Ship) MoveWaypoint(dir Direction, distance int) {
	travel := directionVectors[int(dir) % len(directionVectors)]
	s.waypoint = s.waypoint.Add(travel.Multiply(distance))
}

func (s *Ship) RotateWaypoint(angle int) {
	count := absInt(angle / 90)

	for i := 0; i < count; i++ {
		if angle > 0 {
			s.waypoint = s.waypoint.RotateClockwise90()
		} else if angle < 0 {
			s.waypoint = s.waypoint.RotateAntiClockwise90()
		}
	}
}

func (s *Ship) HandleInstruction(i Instruction) {
	switch i.action {
	case North:
		s.MoveWaypoint(N, i.value)
		break
	case South:
		s.MoveWaypoint(S, i.value)
		break
	case East:
		s.MoveWaypoint(E, i.value)
		break
	case West:
		s.MoveWaypoint(W, i.value)
		break
	case Left:
		s.RotateWaypoint(-i.value)
		break
	case Right:
		s.RotateWaypoint(i.value)
		break
	case Forward:
		s.MoveToWaypoint(i.value)
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

	fmt.Printf("Ship Position: (%d, %d) Waypoint: (%d, %d) ManhattanDistance: %d\n", ship.position.x, ship.position.y, ship.waypoint.x, ship.waypoint.y, ship.position.ManhattanDistance())
}