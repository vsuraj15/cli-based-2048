package config

type Direction int

const (
	DirectionUP = iota
	DirectionDown
	DirectionLeft
	DirectionRight
	NoDirection
)

func (b *board) move(direction Direction) {
	switch direction {
	case DirectionUP:
		b.moveUP()
	case DirectionDown:
		b.moveDown()
	case DirectionRight:
		b.moveRight()
	case DirectionLeft:
		b.moveLeft()
	}
}

func (b *board) moveLeft() {
	for i := 0; i < b.rows; i++ {
		old := b.matrix[i]
		b.matrix[i] = moveCell(old, b.columns)
	}
}
func (b *board) moveRight() {
	b.reverseRows()
	b.moveLeft()
	b.reverseRows()
}
func (b *board) moveDown() {
	b.transpose()
	b.moveLeft()
	b.transpose()
	b.transpose()
	b.transpose()
}
func (b *board) moveUP() {
	b.swapRows()
	b.moveDown()
	b.swapRows()
}

func (b *board) reverseRows() {
	for i := 0; i < b.rows; i++ {
		b.matrix[i] = reverse(b.matrix[i])
	}
}

func (b *board) transpose() {
	result := make([][]int, 0)
	for i := 0; i < b.rows; i++ {
		result = append(result, make([]int, b.columns))
	}
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.columns; c++ {
			result[r][c] = b.matrix[b.columns-c-1][r]
		}
	}
	b.matrix = result
}

func (b *board) swapRows() {
	result := make([][]int, 0)
	for i := 0; i < b.rows; i++ {
		result = append(result, make([]int, b.columns))
	}
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.columns; c++ {
			result[b.rows-r-1][c] = b.matrix[r][c]
		}
	}
	b.matrix = result
}
