package helpers

import (
    "os"
    "bufio"
)

func ReadFileLines(path string) ([]string, error) {
    var file, err = os.Open(path)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    var lines []string
    var scanner = bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines, scanner.Err()
}
