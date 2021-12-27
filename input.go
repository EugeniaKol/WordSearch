package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"io"
	"strconv"
	"strings"
)

//method for parsing user input

func (f *Field) Parse(r io.Reader) (Position, bool, error) {
	var position Position
	var err error

	fc.Print("Beginning coordinate is ")
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	coords := strings.Split(scanner.Text(), " ")
	if coords[0] == "finish" {
		fg.Println("Congratulations! Your score is ", f.Score.Points)
		return position, true, nil
	}
	position.Beginning[0], err = strconv.Atoi(coords[1])
	if err != nil {
		fr.Println("Invalid input - correct format is LETTER number")
		return Position{}, false, errors.New("invalid input")
	}
	rs := []rune(coords[0])
	index := int(rs[0])

	if index > 90 || index < 30 {
		fr.Println("Invalid input - correct format is LETTER number")
		return Position{}, false, errors.New("invalid input")
	}
	position.Beginning[1] = index - 65

	fc.Print("End coordinate is ")
	scanner.Scan()

	coords = strings.Split(scanner.Text(), " ")
	if coords[0] == "finish" {
		fg.Println("Congratulations! Your score is ", f.Score.Points)
		return position, true, nil
	}

	position.End[0], err = strconv.Atoi(coords[1])
	if err != nil {
		fr.Println("Invalid input - correct format is LETTER number")
		return Position{}, false, errors.New("invalid input")
	}

	rs = []rune(coords[0])
	index = int(rs[0])
	if index > 90 || index < 30 {
		fr.Println("Invalid input - correct format is LETTER number")
		return Position{}, false, errors.New("invalid input")
	}
	position.End[1] = index - 65

	return position, false, nil
}

//method for formatting and printing output

func (f *Field) Format() {
	ch := 'A'
	for i := 0; i < f.Width; i++ {
		fc.Printf("%c ", ch)
		ch++
	}
	fmt.Printf("   \t Words to find:\n")

	for i := 0; i < f.Height; i++ {
		for j := 0; j < f.Width; j++ {
			if f.Grid[i][j].Found {
				fw.Print(f.Grid[i][j].Char)
			} else {
				fmt.Print(f.Grid[i][j].Char)
			}
			fmt.Print(" ")
		}
		fc.Printf(" %d", i)
		if i < len(f.Items) {
			word := f.Items[i].Word
			if f.Items[i].Found {
				fw.Printf("\t %s ", word)
				fmt.Println()
			} else {
				fmt.Printf("\t %s\n", word)
			}

		} else if i == f.Height-2 {
			fr := color.New(color.FgRed)
			fr.Printf("\t Your score is %d\n", f.Score.Points)
		} else if i == f.Height-1 {
			fr := color.New(color.FgRed)
			fr.Printf("\t Enter \"")
			fmt.Print("finish")
			fr.Printf("\" to give up\n")
			fmt.Println()
		} else {
			fmt.Printf("\n")
		}
	}
}
