package data

type Question struct {
	Text       string   `json:"text"`
	Options    []string `json:"options"`
	CorrectIdx int      `json:"correctIdx"`
}

var Quiz = []Question{
	{
		Text:       "What is the capital of France?",
		Options:    []string{"London", "Berlin", "Paris", "Madrid"},
		CorrectIdx: 2,
	},
	{
		Text:       "Which famous playwright wrote 'Romeo and Juliet'?",
		Options:    []string{"William Shakespeare", "Jane Austen", "George Orwell", "Mark Twain"},
		CorrectIdx: 0,
	},
	{
		Text:       "Which planet is known as the Red Planet?",
		Options:    []string{"Mars", "Venus", "Jupiter", "Saturn"},
		CorrectIdx: 0,
	},
}
