package model

// Team struct
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

// Player struct
type Player struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Nationality string `json:"nationality"`
	Goals       string `json:"goals"`
	Assists     string `json:"assists"`
	TeamID      int    `json:"teamid"`
}
