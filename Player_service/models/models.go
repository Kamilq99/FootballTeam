package models

type Player struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Position string `json:"position"`
	Age      int    `json:"age"`
}

var Players []Player = []Player{
	{
		ID:       1,
		Name:     "Thibaut",
		LastName: "Courtois",
		Position: "Goalkeeper",
		Age:      33,
	},
	{
		ID:       2,
		Name:     "Andriy",
		LastName: "Lunin",
		Position: "Goalkeeper",
		Age:      30,
	},
}
