package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"
)

const upgradedBufferSize = 1024

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// League stores a collection of players.
type League []Player

// Player stores a name with a number of wins.
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	game     Game
}

// NewLeague creates a league from JSON.
func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

// Find tries to return a player from a league.
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", jsonContentType)

	err := json.NewEncoder(w).Encode(p.store.GetLeague())
	if err != nil {
		panic(err)
	}
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) playGame(w http.ResponseWriter, _ *http.Request) {
	tmpl, parseErr := template.ParseFiles("game.html")

	if parseErr != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", parseErr.Error()), http.StatusInternalServerError)
		return
	}

	execErr := tmpl.Execute(w, nil)
	if execErr != nil {
		panic(execErr)
	}
}

//nolint:exhaustruct,gochecknoglobals
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  upgradedBufferSize,
	WriteBufferSize: upgradedBufferSize,
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)

	_, numberOfPlayersMsg, _ := conn.ReadMessage()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.game.Start(numberOfPlayers, io.Discard) //to do: Don't discard the blinds messages!

	_, winner, _ := conn.ReadMessage()
	p.game.Finish(string(winner))
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
