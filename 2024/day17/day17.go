package day17

import (
	. "aoc2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var A, B, C int
var program []int

var idxPtr int

var literals = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
var combos = [8](*int){&literals[0], &literals[1], &literals[2], &literals[3], &A, &B, &C, nil}
var instructions = [8]func(operand int){adv, bxl, bst, jnz, bxc, out, bdv, cdv}
var jumped = false
var output []int

var buf = bufio.NewReader(os.Stdin)

func ChronospatialComputer(input []string) string {
	A, B, C, program = parseInput(input)
    output = []int{}
    idxPtr = 0

    printState()

	for idxPtr < len(program) {
		opcode, operand := program[idxPtr], program[idxPtr+1]
		instruction := instructions[opcode]
		instruction(operand)

		if !jumped {
			idxPtr += 2
		} else {
			jumped = false
		}

        printState()
	}

	return strings.Join(Map(output, strconv.Itoa), ",")
}

func parseInput(input []string) (a int, b int, c int, program []int) {
	a, _ = strconv.Atoi(strings.Split(input[0], " ")[2])
	b, _ = strconv.Atoi(strings.Split(input[1], " ")[2])
	c, _ = strconv.Atoi(strings.Split(input[2], " ")[2])

	programStr := strings.Split(input[4], " ")[1]
	return a, b, c, Map(strings.Split(programStr, ","), func(instruction string) int {
		value, _ := strconv.Atoi(instruction)
		return value
	})
}

func xdv(operand int, register *int) {
	numerator := float64(A)
	denominator := math.Pow(2, float64(*combos[operand]))
	result := int(math.Floor(numerator / denominator))
	*register = result
}

func adv(operand int) {
    fmt.Println("adv")
	xdv(operand, &A)
}

func bxl(operand int) {
    fmt.Println("bxl")
	B = B | literals[operand]
}

func bst(operand int) {
    fmt.Println("bst")
	B = *combos[operand] % 8
}

func jnz(operand int) {
    fmt.Println("jnz")
	if A == 0 {
		return
	}
	idxPtr = literals[operand]
    fmt.Println("Jumped to", idxPtr)
	jumped = true
}

func bxc(operand int) {
    fmt.Println("bxc")
	B = B | C
}

func out(operand int) {
    fmt.Println("out")
	output = append(output, *combos[operand] % 8)
}

func bdv(operand int) {
    fmt.Println("bdv")
	xdv(operand, &B)
}

func cdv(operand int) {
    fmt.Println("cdv")
	xdv(operand, &C)
}

func printState() {
    fmt.Println("STATE")
    fmt.Println("Program:", program)
    fmt.Println("idxPtr:", idxPtr)
    fmt.Println("A:", A)
    fmt.Println("B:", B)
    fmt.Println("C:", C)
    fmt.Println("output:", output)
    fmt.Println()
}
