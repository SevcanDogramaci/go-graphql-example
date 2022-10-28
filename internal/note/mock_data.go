package note

var NOTES = []Note{
	{
		ID:          "1",
		Title:       "Meeting Note",
		Description: "Prioritize refactoring task",
		Author: Author{
			ID:   "1",
			Name: "Jasmine",
			Age:  24,
		},
	},
	{
		ID:          "2",
		Title:       "Self Note",
		Description: "Love yourself !",
		Author: Author{
			ID:   "1",
			Name: "Jasmine",
			Age:  24,
		},
	},
	{
		ID:          "3",
		Title:       "Mom's Note",
		Description: "Call me :)",
		Author: Author{
			ID:   "1",
			Name: "Jane",
			Age:  58,
		},
	},
}
