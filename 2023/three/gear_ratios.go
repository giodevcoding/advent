package three

import (
	"fmt"
)

func RunPartOne(input []string) {
    IsSymbol("5")
    IsSymbol('.')
}

func IsSymbol[T rune | string](input T) bool {
    inputType := fmt.Sprintf("%T", input)
    
    if inputType == "string" {

    }

    fmt.Println(inputType)

    return true
}
