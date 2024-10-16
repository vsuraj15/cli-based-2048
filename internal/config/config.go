package config

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const (
	probabilitySpace          = 100
	probabilityOccuranceOfTwo = 80
	// this is the sequence which is used to clear the screen :magic
	clearScreenSequence = "\033[H\033[2J" // this works in mac. Might need other string for other OS
)

type Board interface {
	Display()
	AddElement()
	TakeInput()
	IsOver() bool
	CountScore() (int, int)
}

type board struct {
	matrix  [][]int
	newRow  int
	newCol  int
	rows    int
	columns int
	over    bool
}

func CreateBoard(rows, cols int) Board {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		matrix = append(matrix, make([]int, cols))
	}
	return &board{
		matrix:  matrix,
		rows:    rows,
		columns: cols,
	}
}

func (b *board) Display() {
	d := color.New(color.FgBlue, color.Bold)
	fmt.Println(clearScreenSequence)
	for i := 0; i < len(b.matrix); i++ {
		printHorizontal(b.rows)
		fmt.Printf("|")
		for j := 0; j < len(b.matrix[0]); j++ {
			fmt.Printf("%3s", "")
			if b.matrix[i][j] == 0 {
				fmt.Printf("%-6s|", "")
			} else {
				if i == b.newRow && j == b.newCol {
					d.Printf("%-6d|", b.matrix[i][j])
				} else {
					fmt.Printf("%-6d|", b.matrix[i][j])
				}
			}
		}
		fmt.Printf("%4s", "")
		fmt.Println()
	}
	printHorizontal(b.rows)
}

func (b *board) AddElement() {
	source1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source1)
	value := r1.Int() % probabilitySpace
	if value <= probabilityOccuranceOfTwo {
		value = 2
	} else {
		value = 4
	}
	empty := 0
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.columns; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	elementCount := r1.Int()%empty + 1
	index := 0
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.columns; j++ {
			if b.matrix[i][j] == 0 {
				index++
				if index == elementCount {
					b.newRow = i
					b.newCol = j
					b.matrix[i][j] = value
					return
				}
			}
		}
	}
	return
}

func (b *board) TakeInput() {
	var dir Direction
	var err error
	dir, err = GetKeyStroke()
	if err != nil {
		if errors.Is(err, errEndGame) {
			b.over = true
			return
		} else {
			log.Fatalf("Unknown key stroke found. %+v\n", err)
			return
		}
	}
	if dir == NoDirection {
		b.TakeInput()
	}
	b.move(dir)
}

func (b *board) IsOver() bool {
	empty := 0
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.columns; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	return empty == 0 || b.over
}

func (b *board) CountScore() (int, int) {
	total := 0
	maximum := 0
	matrix := b.matrix
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.columns; j++ {
			total += matrix[i][j]
			maximum = getMaximum(maximum, matrix[i][j])
		}
	}
	return maximum, total
}
