package model

// Team struct
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

// Player struct
type Player struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

// TeamPlayer struct
type TeamPlayer struct {
	PlayerId *Player `json:"playerid"`
	TeamId   *Team   `json:"teamid"`
}
