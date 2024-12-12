package day9

import (
	"aoc2024/utils"
	"bufio"
	"os"
	"strconv"
)

var buf = bufio.NewReader(os.Stdin)

func UpdateChecksum(input []string) string {
	blocks := spreadBlocks(input[0])
    blocks = rearrangeBlocks(blocks)
    checksum := calculateChecksum(blocks)
	return strconv.Itoa(checksum)
}

func UpdateChecksumWholeFiles(input []string) string {
	blocks := spreadBlocks(input[0])
    blocks = rearrangeWholeBlocks(blocks)
    checksum := calculateChecksum(blocks)
	return strconv.Itoa(checksum)
}

func spreadBlocks(blocks string) []string {
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

func rearrangeBlocks(blocks []string) []string {
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
    return blocks
}

func rearrangeWholeBlocks(blocks []string) []string {
    for i := len(blocks)-1; i >= 0; i-- {
        if blocks[i] == "." {
            continue
        }
        startIndex, fileBlocksLength := getFileIdBlocks(i, blocks)
        i = startIndex

        spaceCount := 0
        for j := 0; j < len(blocks); j++ {
            if blocks[j] != "." {
                spaceCount = 0
                continue
            }

            if j >= i {
                break
            }
            
            spaceCount++
            if spaceCount == fileBlocksLength {
                spaceStartIndex := (j - spaceCount) + 1
                for k := 0; k < fileBlocksLength; k++ {
                    blocks[startIndex+k], blocks[spaceStartIndex+k] = blocks[spaceStartIndex+k], blocks[startIndex+k]
                }
                break
            }
        }
    }
    return blocks
}

func getFileIdBlocks(endIndex int, blocks []string) (startIndex, length int) {
    currentFileId := blocks[endIndex]
    i := endIndex
    for i >= 0 {
        if blocks[i] != currentFileId {
            break
        }
        startIndex = i
        i--
    }
    length = (endIndex - startIndex)+1

    return
}


func calculateChecksum(blocks []string) (total int) {
    for i, fileId := range blocks {
        fileIdInt, _ := strconv.Atoi(fileId)
        total += i * fileIdInt
    } 
    return
}
