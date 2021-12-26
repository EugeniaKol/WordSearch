package main

var field = &Field{
	Width:  10,
	Height: 10,
	Words:  []string{"garden", "search", "message", "son", "icecream"},
	Items: []*Item{
		{
			Word: "garden",
			Position: Position{
				Beginning: [2]int{0, 0},
				End:       [2]int{0, 5},
			},
		},
		{
			Word: "search",
			Position: Position{
				Beginning: [2]int{6, 5},
				End:       [2]int{1, 0},
			},
		},
		{
			Word: "message",
			Position: Position{
				Beginning: [2]int{6, 4},
				End:       [2]int{0, 4},
			},
		},
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
		{
			Word: "testing",
			Position: Position{
				Beginning: [2]int{6, 6},
				End:       [2]int{0, 0},
			},
		},
	},
}
