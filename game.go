package main

import (
	"io"
	"strings"
)

//structs for game logic

type Cell struct {
	Char  string
	Found bool
}

type Field struct {
	Width  int
	Height int
	Grid   [][]Cell
	Words  []string
	Items  []*Item
	Score  *Score
}

type Item struct {
	Word     string
	Position Position
}

type Position struct {
	Beginning [2]int
	End       [2]int
}

type Score struct {
	Discovered int
	Points     int
}

//method that fills a grid with random letters

func (f *Field) Fill() {

}

func sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

//method that fills the grid with words according to their positions

func (f *Field) SetWords() {
	for i := 0; i < f.Height; i++ {
		f.Grid = append(f.Grid, make([]Cell, f.Width))
	}

	for _, item := range f.Items {
		chars := strings.Split(item.Word, "")
		start, end := item.Position.Beginning, item.Position.End
		i, j := start[0], start[1]

		var shift [2]int
		shift[0] = sgn(end[0] - start[0])
		shift[1] = sgn(end[1] - start[1])

		for _, char := range chars {
			f.Grid[i][j].Char = char
			i += shift[0]
			j += shift[1]
		}
	}
}

//method that takes coordinates of a grid and returns a bool
//meaning 'true' if the word is there

func (f *Field) Check(input Position) (bool, string) {

}

//method that ties together game's logic and user input/output

func (f *Field) Game() {

}

//method for parsing user input

func (f *Field) Parse(r io.Reader) (Position, error) {

}

func main() {
	var field = &Field{
		Width:  10,
		Height: 10,
		Words:  []string{"garden", "search", "message", "son", "icecream"},
		Items: []*Item{
			{
				Word: "search",
				Position: Position{
					Beginning: [2]int{6, 5},
					End:       [2]int{1, 0},
				},
			},
			{
				Word: "garden",
				Position: Position{
					Beginning: [2]int{0, 0},
					End:       [2]int{0, 5},
				},
			},
		},
	}
	field.SetWords()
	field.Fill()
	field.Game()
}
