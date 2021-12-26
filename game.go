package main

import (
	"io"
)

//structs for game logic

type Field struct {
	Width  int
	Height int
	Grid   [][]string
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

//method that fills the grid with words according to their positions

func (f *Field) SetWords() {

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
		},
	}
	field.SetWords()
	field.Fill()
	field.Game()
}
