package day9

import (
	"aoc2024/utils"
	"bufio"
	"os"
	"strconv"
)

var buf = bufio.NewReader(os.Stdin)

func UpdateChecksum(input []string) string {
	blocks := SpreadBlocks(input[0])
    blockInts := RearrangeBlocks(blocks)
    checksum := CalculateChecksum(blockInts)
	return strconv.Itoa(checksum)
}

func SpreadBlocks(blocks string) []string {
	result := make([]string, 0)
	fileIndex := 0
	for i, char := range blocks {
		count := int(char - '0')
		isFile := i%2 == 0
		currentBlockChunk := "."

		if isFile {
			currentBlockChunk = strconv.Itoa(fileIndex)
			fileIndex++
		}

        for j := 0; j < count; j++ {
			result = append(result, currentBlockChunk)
		}
	}
    return result
}

func RearrangeBlocks(blocks []string) []int {
    for i, fileId := range blocks {
        if fileId != "." {
            continue
        }

        for j := len(blocks)-1; j >= 0; j-- {
            if blocks[j] == "." {
                continue
            }

            if (j <= i) {
                break
            }
    
            blocks[i], blocks[j] = blocks[j], blocks[i]
            break
        }
    } 

    blocks = utils.Filter(blocks, func(s string) bool {
        return s != "."
    })
    blockInts := utils.Map(blocks, func(s string) int {
        num, _ := strconv.Atoi(s)
        return num
    })
    return blockInts
}

func CalculateChecksum(blockInts []int) (total int) {
    for i, fileId := range blockInts {
        total += i * fileId
    } 
    return
}
