package day11

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var buf = bufio.NewReader(os.Stdin)

type ChangedStoneInfo struct {
    Value string
    BlinkNumber int
}

var changedStoneMemo = make(map[ChangedStoneInfo]int)

func MagicStonesBlinking(input []string) string {
    stones := strings.Split(input[0], " ")
    stonesAfterBlinking := blink(stones, 0, 25)
    return strconv.Itoa(len(stonesAfterBlinking))
}

func MagicStonesBlinkingFast(input []string) string {
    stones := strings.Split(input[0], " ")
    count := 0
    for _, stone := range stones {
        count += blinkFast(stone, 75)
    }
    return strconv.Itoa(count)
}

func blink(stones []string, timesBlinked, timesToBlink int) []string {
    if (timesBlinked == timesToBlink) {
        return stones
    }

    changedStones := []string{}
    for _, stone := range stones {
        changedStones = append(changedStones, changeStone(stone)...)
    }

    return blink(changedStones, timesBlinked+1, timesToBlink)
}

func blinkFast(stone string, timesLeftToBlink int) int {
    if count, ok := changedStoneMemo[ChangedStoneInfo{stone, timesLeftToBlink}]; ok {
        return count
    }

    if timesLeftToBlink == 0 {
        return 1
    }

    if stone == "0" {
        result := blinkFast("1", timesLeftToBlink-1)
        changedStoneMemo[ChangedStoneInfo{stone, timesLeftToBlink}] = result
        return result
    }

    if len(stone) % 2 == 0 {
        leftStoneValue, _ := strconv.Atoi(stone[:len(stone)/2])
        rightStoneValue, _ := strconv.Atoi(stone[len(stone)/2:])
        leftStone := strconv.Itoa(leftStoneValue)
        rightStone := strconv.Itoa(rightStoneValue)

        result := blinkFast(leftStone, timesLeftToBlink-1) + blinkFast(rightStone, timesLeftToBlink-1)
        changedStoneMemo[ChangedStoneInfo{stone, timesLeftToBlink}] = result
        return result
    }

    stoneValue, _ := strconv.Atoi(stone)
    result := blinkFast(strconv.Itoa(stoneValue*2024), timesLeftToBlink-1)
    changedStoneMemo[ChangedStoneInfo{stone, timesLeftToBlink}] = result
    return result
}

func changeStone(stone string) []string {
    if stone == "0" {
        return []string{"1"}
    }

    if len(stone) % 2 == 0 {
        leftStoneValue, _ := strconv.Atoi(stone[:len(stone)/2])
        rightStoneValue, _ := strconv.Atoi(stone[len(stone)/2:])
        leftStone := strconv.Itoa(leftStoneValue)
        rightStone := strconv.Itoa(rightStoneValue)
        return []string{leftStone, rightStone}
    }

    stoneValue, _ := strconv.Atoi(stone)
    return []string{strconv.Itoa(stoneValue*2024)}
    
}
