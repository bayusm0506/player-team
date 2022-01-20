# Player Team

Simple RESTful API to make player team. No database implementation yet

## Clone Project

```bash
https://github.com/bayusm0506/player-team.git
cd /player-team
```

## Structure Folder

```
├── app
│   ├── app.go
│   ├── handler
│   │   ├── common.go
│   │   └── handler.go
│   └── model
│       └── model.go
├── config
│   └── config.go
├── go.mod
├── go.sum
├── main.go
```

## Installation

```bash
# Install mux router
go get -u github.com/gorilla/mux

# Install godotenv
go get -u github.com/joho/godotenv
```

## Usage

```go
# create file .env in root folder and add variable
APPS_PORT=2022

# running project
go run main.go

# or if you wanna run project with .exe file
go build
./player-team
```

## API

#### GET ALL TEAMS
```sh
GET : /api/teams
```
#### GET TEAMS BY ID
```sh
GET : /api/teams/{id}
```
#### GET ALL PLAYERS
```sh
GET : /api/players
```
#### GET PLAYERS BY ID
```sh
GET : /api/players/{id}
```
#### GET PLAYER TEAM BY TEAM ID
```sh
GET : /api/players/teams/{teamid}
```
#### CREATE A TEAM
```sh
POST : /api/teams
```
```json
JSON : 
{
    "name": "Manchester City",
    "city": "Inggris"
}
```
#### CREATE A PLAYER
```sh
POST : /api/players
```
```json
JSON : 
{
    "name": "Scott McTominay",
    "position": "Midfielder",
    "nationality": "Scotland",
    "goals": "11",
    "assists": "2",
    "teamid": 98498081
}
```

```
#### APP INFO

#### Author originally
Bayu Setra Maulana [bayusm.com]
```
