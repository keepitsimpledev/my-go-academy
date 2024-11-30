//go:build !race

package poker

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

//nolint:gochecknoglobals
var (
	dummyGame = &GameSpy{
		StartCalled:      false,
		StartCalledWith:  0,
		BlindAlert:       []byte{},
		FinishedCalled:   false,
		FinishCalledWith: "",
	}
	dummyPlayerStore = &StubPlayerStore{
		Scores:   make(map[string]int),
		WinCalls: []string{},
		League:   []Player{},
	}
)

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server, _ := NewPlayerServer(dummyPlayerStore, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, *response, http.StatusOK)
	})

	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		game := dummyGame
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
	})

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := dummyGame

		out := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		NewCLI(in, out, game).PlayPoker()

		assertMessagesSentToUser(t, out, PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})
}

func mustMakePlayerServer(tb testing.TB, store PlayerStore, game Game) *PlayerServer {
	tb.Helper()

	server, err := NewPlayerServer(store, game)
	if err != nil {
		tb.Fatal("problem creating player server", err)
	}

	return server
}

func mustDialWS(tb testing.TB, url string) *websocket.Conn {
	tb.Helper()

	ws, response, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		tb.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	defer response.Body.Close()

	return ws
}

func writeWSMessage(tb testing.TB, conn *websocket.Conn, message string) {
	tb.Helper()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		tb.Fatalf("could not send message over ws connection %v", err)
	}
}

func newGameRequest() *http.Request {
	request, err := http.NewRequest(http.MethodGet, "/game", nil)
	if err != nil {
		panic(err)
	}

	return request
}

func assertStatus(tb testing.TB, got httptest.ResponseRecorder, want int) {
	tb.Helper()

	if got.Code != want {
		tb.Errorf("got: %d. want %d.", got.Code, want)
	}
}

// AssertPlayerWin allows you to spy on the store's calls to RecordWin.
func AssertPlayerWin(tb testing.TB, store *StubPlayerStore, winner string) {
	tb.Helper()

	if len(store.WinCalls) != 1 {
		tb.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		tb.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}

func assertGameStartedWith(tb testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	tb.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartCalledWith == numberOfPlayersWanted
	})

	if !passed {
		tb.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}

	return false
}

func assertFinishCalledWith(tb testing.TB, game *GameSpy, winner string) {
	tb.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishCalledWith == winner
	})

	if !passed {
		tb.Errorf("expected finish called with %q but got %q", winner, game.FinishCalledWith)
	}
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertMessagesSentToUser(tb testing.TB, stdout *bytes.Buffer, messages ...string) {
	tb.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		tb.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
