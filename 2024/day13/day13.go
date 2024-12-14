package day13

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Coordinate struct {
	X, Y int
}

type ClawMachine struct {
	AButton, BButton, Prize Coordinate
}

type PressCombo struct {
	ATimes, BTimes int
}

var buf = bufio.NewReader(os.Stdin)

const PRIZE_OFFSET = 10000000000000
const A_TOKEN_COST = 3
const B_TOKEN_COST = 1

func ClawMachinesFewestTokens(input []string) string {
	totalTokens := 0
	clawMachines := getClawMachines(input)
	for _, clawMachine := range clawMachines {
		validCombos := hitButtons(0, 0, clawMachine, nil, nil)
		if len(validCombos) > 0 {
			leastTokens := math.MaxInt
			for _, combo := range validCombos {
				leastTokens = int(math.Min(float64(leastTokens), float64((combo.ATimes*A_TOKEN_COST)+(combo.BTimes*B_TOKEN_COST))))
			}
			totalTokens += leastTokens
		}
	}
	return strconv.Itoa(totalTokens)
}

func ClawMachineMatrixMath(input []string) string {
	totalTokens := 0
	clawMachines := getClawMachines(input)
	for _, clawMachine := range clawMachines {
		ax, ay := float64(clawMachine.AButton.X), float64(clawMachine.AButton.Y)
		bx, by := float64(clawMachine.BButton.X), float64(clawMachine.BButton.Y)
		px, py := float64(clawMachine.Prize.X + PRIZE_OFFSET), float64(clawMachine.Prize.Y + PRIZE_OFFSET)

        // Equations
        // ax + bx = px
        // ay + by = py
        //
        // Matrix
        // | ax bx | | a | = | px |
        // | ay by | | b | = | py |
        //
        // Inverse Matrix
        //                       | by -bx |
        // 1 / (ax*by - ay*bx) * | -ay ax |
        //
        // Solution
		inv := 1 / (ax*by - ay*bx)
		aHits := inv*by*px + inv*(-bx)*py
		bHits := inv*(-ay)*px + inv*ax*py

		if isWholeNumber(aHits) && isWholeNumber(bHits) {
			totalTokens += int(math.Round(aHits))*A_TOKEN_COST + int(math.Round(bHits))*B_TOKEN_COST
		}
	}
	return strconv.Itoa(totalTokens)
}

func isWholeNumber(num float64) bool {
	return (math.Round(num*1000) / 1000) == math.Round(num)
}

func hitButtons(aTimes, bTimes int, clawMachine ClawMachine, checkedCombos *map[PressCombo]bool, validCombos *[]PressCombo) []PressCombo {
	if checkedCombos == nil {
		checkedCombos = &map[PressCombo]bool{}
	}

	if validCombos == nil {
		validCombos = &[]PressCombo{}
	}

	aButton, bButton, prize := clawMachine.AButton, clawMachine.BButton, clawMachine.Prize
	if _, ok := (*checkedCombos)[PressCombo{aTimes, bTimes}]; ok {
		return *validCombos
	}

	if (aButton.X*aTimes)+(bButton.X*bTimes) == prize.X &&
		(aButton.Y*aTimes)+(bButton.Y*bTimes) == prize.Y {
		(*checkedCombos)[PressCombo{aTimes, bTimes}] = true
		*validCombos = append(*validCombos, PressCombo{aTimes, bTimes})
		return *validCombos
	}

	if (aButton.X*aTimes)+(bButton.X*bTimes) > prize.X ||
		(aButton.Y*aTimes)+(bButton.Y*bTimes) > prize.Y {
		(*checkedCombos)[PressCombo{aTimes, bTimes}] = true
		return *validCombos
	}

	if _, ok := (*checkedCombos)[PressCombo{aTimes, bTimes + 1}]; !ok {
		hitButtons(aTimes, bTimes+1, clawMachine, checkedCombos, validCombos)
	}

	if _, ok := (*checkedCombos)[PressCombo{aTimes + 1, bTimes}]; !ok {
		hitButtons(aTimes+1, bTimes, clawMachine, checkedCombos, validCombos)
	}

	(*checkedCombos)[PressCombo{aTimes, bTimes}] = true
	return *validCombos
}

func getClawMachines(input []string) (clawMachines []ClawMachine) {

	for i := 0; i < len(input); i += 4 {
		clawMachine := ClawMachine{}
		buttonRegex := regexp.MustCompile(`X\+(\d+).*Y\+(\d+)`)
		prizeRegex := regexp.MustCompile(`X=(\d+).*Y=(\d+)`)

		aMatches := buttonRegex.FindAllStringSubmatch(input[i], 2)
		bMatches := buttonRegex.FindAllStringSubmatch(input[i+1], 2)
		prizeMatches := prizeRegex.FindAllStringSubmatch(input[i+2], 2)

		clawMachine.AButton.X, _ = strconv.Atoi(aMatches[0][1])
		clawMachine.AButton.Y, _ = strconv.Atoi(aMatches[0][2])

		clawMachine.BButton.X, _ = strconv.Atoi(bMatches[0][1])
		clawMachine.BButton.Y, _ = strconv.Atoi(bMatches[0][2])

		clawMachine.Prize.X, _ = strconv.Atoi(prizeMatches[0][1])
		clawMachine.Prize.Y, _ = strconv.Atoi(prizeMatches[0][2])

		clawMachines = append(clawMachines, clawMachine)
	}

	return
}
