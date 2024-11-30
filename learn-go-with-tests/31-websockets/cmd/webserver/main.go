package main

import (
	"go_academy/learn-go-with-tests/31-websockets"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	//nolint:gomnd
	db, openErr := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if openErr != nil {
		log.Fatalf("problem opening %s %v", dbFileName, openErr)
	}

	store, nfspsErr := poker.NewFileSystemPlayerStore(db)

	if nfspsErr != nil {
		log.Fatalf("problem creating file system player store, %v ", nfspsErr)
	}

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, npsErr := poker.NewPlayerServer(store, game)
	if npsErr != nil {
		panic(openErr)
	}

	//nolint:gosec
	log.Fatal(http.ListenAndServe(":5000", server))
}
