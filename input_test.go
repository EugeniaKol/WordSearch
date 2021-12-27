package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestParse(c *C) {
	field := Field{}
	for _, testCase := range []struct {
		input    string
		expected Position
		finish   bool
		err      error
	}{
		{
			input: "A 4\nD 0\n",
			expected: Position{
				Beginning: [2]int{4, 0},
				End:       [2]int{0, 3},
			},
		},
		{
			input: "a 4\nD 0\n",
			err:   errors.New("invalid input"),
		},
		{
			input: "K 1\nF 4\n",
			expected: Position{
				Beginning: [2]int{1, 10},
				End:       [2]int{4, 5},
			},
		},
	} {
		reader := strings.NewReader(testCase.input)
		res, fin, err := field.Parse(reader)
		matches := reflect.DeepEqual(res, testCase.expected)
		fmt.Println(res)
		c.Assert(matches, Equals, true)
		c.Assert(fin, Equals, testCase.finish)
		if err != nil {
			c.Assert(err, ErrorMatches, testCase.err.Error())
		}
	}
}
