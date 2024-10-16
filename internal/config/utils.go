package config

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

var (
	errEndGame   = fmt.Errorf("** GAME Over **")
	unknownInput = fmt.Errorf("Unknown input found. GameOver")
)

var (
	aCharKey      = 97
	dCharKey      = 100
	hCharKey      = 104
	jCharKey      = 106
	kCharKey      = 107
	lCharKey      = 108
	sCharKey      = 115
	wCharKey      = 119
	rightArrowKey = 65514
	leftArrowKey  = 65515
	downArrowKey  = 65516
	upArrowKey    = 65517
)

func HaltIfEmpty(paramValue int) bool {
	return paramValue < 4
}

func GetKeyStroke() (Direction, error) {
	if err := keyboard.Open(); err != nil {
		log.Fatalf("Failed to open keyboard stroke. Err: %+v\n", err)
	}
	defer keyboard.Close()
	char, key, getErr := keyboard.GetKey()
	if getErr != nil {
		return NoDirection, getErr
	}
	answr := int(char)
	if answr == 0 {
		answr = int(key)
	}
	switch answr {
	case wCharKey, upArrowKey, kCharKey:
		return DirectionUP, nil
	case aCharKey, leftArrowKey, hCharKey:
		return DirectionLeft, nil
	case sCharKey, downArrowKey, jCharKey:
		return DirectionDown, nil
	case dCharKey, rightArrowKey, lCharKey:
		return DirectionRight, nil
	case 3:
		return NoDirection, errEndGame
	}
	return NoDirection, nil
}

func getMaximum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func moveCell(elements []int, column int) []int {
	nonEmpty := make([]int, 0)
	for i := 0; i < column; i++ {
		if elements[i] != 0 {
			nonEmpty = append(nonEmpty, elements[i])
		}
	}
	remaining := column - len(nonEmpty)
	for pos := 0; pos < remaining; pos++ {
		nonEmpty = append(nonEmpty, 0)
	}
	return mergeElements(nonEmpty)
}

func mergeElements(elements []int) []int {
	newArr := make([]int, len(elements))
	newArr[0] = elements[0]
	index := 0
	for i := 1; i < len(elements); i++ {
		if elements[i] == newArr[index] {
			newArr[index] += elements[i]
		} else {
			index++
			newArr[index] = elements[i]
		}
	}
	return newArr
}

func reverse(arr []int) []int {
	result := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

// printHorizontal prints a grid row
func printHorizontal(rows int) {
	for i := 0; i < 10*rows; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
