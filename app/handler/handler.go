package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/bayusm0506/player-team/app/model"
	"github.com/gorilla/mux"
)

// Init teams var as a slice model.Team struct
var teams []model.Team

// Init players var as a slice model.Player struct
var players []model.Player

// Init playersTeam var as a slice model.Player struct
var playersTeam []model.Player

// Create a Team
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	team := model.Team{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&team); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	team.ID = rand.Intn(100000000)
	teams = append(teams, team)
	defer r.Body.Close()

	respondJSON(w, http.StatusCreated, teams)
}

// Add a Player
func AddPlayer(w http.ResponseWriter, r *http.Request) {
	player := model.Player{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&player); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}

	player.ID = rand.Intn(100000000)
	players = append(players, player)
	defer r.Body.Close()

	respondJSON(w, http.StatusCreated, players)
}

// Get a Teams
func GetTeams(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, teams)
}

// Get a Team By ID
func GetTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	id, _ := strconv.Atoi(params["id"])

	for _, item := range teams {
		if item.ID == id {
			respondJSON(w, http.StatusOK, item)
			return
		}
	}

	respondJSON(w, http.StatusOK, teams)
}

// Get a Players
func GetPlayers(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, players)
}

// Get a Player By ID
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	id, _ := strconv.Atoi(params["id"])

	for _, item := range players {
		if item.ID == id {
			respondJSON(w, http.StatusOK, item)
			return
		}
	}

	respondJSON(w, http.StatusOK, players)
}

// Get a Player By Team ID
func GetPlayerByTeamID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	teamid, _ := strconv.Atoi(params["teamid"])

	for _, item := range players {
		if item.TeamID == teamid {
			playersTeam = append(playersTeam, item)
		}
	}

	respondJSON(w, http.StatusOK, playersTeam)
}
