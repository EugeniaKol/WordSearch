package main

import (
	"reflect"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

//test for a method that fills all empty cells in a Grid with random letters

func (s *MySuite) TestFill(c *C) {
	field := &Field{
		Width:  5,
		Height: 5,
		Words:  []string{"test"},
		Items: []*Item{
			{
				Word: "test",
				Position: Position{
					Beginning: [2]int{0, 1},
					End:       [2]int{0, 4},
				},
			},
		},
	}
	field.SetWords()

	for _, testCase := range []struct {
		input *Field
	}{
		{
			input: field,
		},
	} {
		testCase.input.Fill()

		for i := 0; i < testCase.input.Height; i++ {
			for j := 0; j < testCase.input.Width; j++ {
				c.Assert(testCase.input.Grid[i][j].Char, Not(Equals), "")
			}
		}
	}
}

//test for a method that fills a Grid of a Field struct
//with words passed as an Item struct

func (s *MySuite) TestSetWords(c *C) {
	field := Field{
		Width:  10,
		Height: 10,
	}

	for _, testCase := range []struct {
		input Field
		items []*Item
	}{
		//you can add test cases by adding structs with a Field,
		//and Items that are expected to be placed into the Grid
		{
			input: field,
			items: []*Item{
				{
					Word: "message",
					Position: Position{
						Beginning: [2]int{6, 4},
						End:       [2]int{0, 4},
					},
				},
			},
		},

		{
			input: field,
			items: []*Item{
				{
					Word: "message",
					Position: Position{
						Beginning: [2]int{6, 4},
						End:       [2]int{0, 4},
					},
				},
				{
					Word: "testing",
					Position: Position{
						Beginning: [2]int{6, 6},
						End:       [2]int{0, 0},
					},
				},
			},
		},
		{
			input: field,
			items: []*Item{
				{
					Word: "icecream",
					Position: Position{
						Beginning: [2]int{8, 9},
						End:       [2]int{8, 2},
					},
				},
			},
		},
	} {
		testCase.input.Items = testCase.items
		testCase.input.SetWords()

		for _, item := range testCase.input.Items {
			chars := strings.Split(item.Word, "")
			start, end := item.Position.Beginning, item.Position.End
			i, j := start[0], start[1]

			var shift [2]int
			shift[0] = sgn(end[0] - start[0])
			shift[1] = sgn(end[1] - start[1])

			for _, char := range chars {
				actual := testCase.input.Grid[i][j].Char
				c.Assert(actual, Equals, char)
				i += shift[0]
				j += shift[1]
			}
		}
	}
}

//test for the method that updates user score
//increases total score in case of success and decreases in case of fail

func (s *MySuite) TestUpdateScore(c *C) {
	var field Field
	for _, testCase := range []struct {
		before   Score
		word     string
		success  bool
		expected Score
	}{
		{
			before: Score{
				Discovered: 0,
				Points:     0,
				Streak:     0,
			},

			word:    "word",
			success: true,

			expected: Score{
				Discovered: 1,
				Points:     400,
				Streak:     1,
			},
		},
		{
			before: Score{
				Discovered: 4,
				Points:     1000,
				Streak:     3,
			},

			word:    "word",
			success: true,

			expected: Score{
				Discovered: 5,
				Points:     1700,
				Streak:     4,
			},
		},
		{
			before: Score{
				Discovered: 5,
				Points:     1700,
				Streak:     4,
			},

			word:    "word",
			success: false,

			expected: Score{
				Discovered: 5,
				Points:     1600,
				Streak:     0,
			},
		},
	} {
		field.Score = &testCase.before
		field.UpdateScore(testCase.word, testCase.success)
		equal := reflect.DeepEqual(testCase.expected, *field.Score)
		c.Assert(equal, Equals, true)
	}
}

//test for a method that takes a Position struct and returns
//a bool indicating whether there is a word set at this position
//and the word itself

func (s *MySuite) TestCheck(c *C) {
	field := Field{
		Width:  10,
		Height: 10,
		Items: []*Item{
			{
				Word: "son",
				Position: Position{
					Beginning: [2]int{0, 8},
					End:       [2]int{2, 6},
				},
			},
			{
				Word: "icecream",
				Position: Position{
					Beginning: [2]int{8, 9},
					End:       [2]int{8, 2},
				},
			},
		},
		Score: &Score{},
	}
	for _, testCase := range []struct {
		field    *Field
		position Position
		expected bool
	}{
		//you can add test cases by adding structs with a Field,
		//a Position input and expected input
		{
			field: &field,
			position: Position{
				Beginning: [2]int{0, 0},
				End:       [2]int{5, 5},
			},
			expected: false,
		},
		{
			field: &field,
			position: Position{
				Beginning: [2]int{0, 8},
				End:       [2]int{2, 6},
			},
			expected: true,
		},
		{
			field: &field,
			position: Position{
				Beginning: [2]int{5, 5},
				End:       [2]int{0, 0},
			},
			expected: false,
		},
		{
			field: &field,
			position: Position{
				Beginning: [2]int{8, 9},
				End:       [2]int{8, 2},
			},
			expected: true,
		},
	} {
		testCase.field.SetWords()
		testCase.field.Fill()
		res, _ := testCase.field.Check(testCase.position)
		c.Assert(res, Equals, testCase.expected)
	}
}
