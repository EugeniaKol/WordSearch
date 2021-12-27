package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
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
	Found    bool
}

type Position struct {
	Beginning [2]int
	End       [2]int
}

type Score struct {
	Discovered int
	Points     int
	Streak     int
}

//method that fills a grid with random letters

func (f *Field) Fill() {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	for i := 0; i < f.Height; i++ {
		for j := 0; j < f.Width; j++ {
			if f.Grid[i][j].Char == "" {
				f.Grid[i][j].Char = string(letters[rand.Intn(len(letters))])
			}
		}
	}
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

//method that updates user score after entering a word

func (f *Field) UpdateScore(word string, success bool) {
	if success {
		f.Score.Discovered++
		f.Score.Streak++
		f.Score.Points += len(word)*100 + (f.Score.Streak-1)*100
	} else {
		f.Score.Streak = 0
		f.Score.Points -= 100
	}
}

//method that takes coordinates of a grid and returns a bool
//meaning 'true' if the word is there

func (f *Field) Check(input Position) (bool, string) {
	for _, p := range f.Items {
		if input.Beginning == p.Position.Beginning {
			if input.End == p.Position.End {
				p.Found = true
				f.UpdateScore(p.Word, true)

				chars := strings.Split(p.Word, "")
				start, end := p.Position.Beginning, p.Position.End
				i, j := start[0], start[1]

				var shift [2]int
				shift[0] = sgn(end[0] - start[0])
				shift[1] = sgn(end[1] - start[1])

				for _, _ = range chars {
					f.Grid[i][j].Found = true
					i += shift[0]
					j += shift[1]
				}
				return true, p.Word
			}
		}
	}
	f.UpdateScore("", false)
	return false, ""
}

//method that ties together game's logic and user input/output

func (f *Field) Game() {
	f.Score = &Score{}
	for {
		f.Format()

		position, finish, err := f.Parse(os.Stdin)
		if err != nil {
			continue
		}
		if finish {
			break
		}
		found, word := f.Check(position)

		if found {
			fg.Print("A Word found - ")
			fc.Print(word)
			fg.Println("! Well done!")
		} else {
			fr.Println("You were wrong, try again :(")
		}

		time.Sleep(1 * time.Second)
		if f.Score.Discovered == len(f.Items) {
			f.Format()
			fg.Println("Congratulations! Your score is", f.Score.Points)
			break
		}
	}
}
