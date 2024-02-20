package server

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {

	score, ok := s.scores[name]

	if !ok {
		return 0, errors.New("player not found")
	}

	return score, nil
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int {
			"Pepper": 20,
			"Floyd": 10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	for name, score := range store.scores {
		t.Run(testName(name), func (t *testing.T) {
			request := newGetScoreRequest(name)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertResponseBody(t, response.Body.String(), fmt.Sprintf("%d", score))
		})
	}

	t.Run("return 404 on missing player", func (t *testing.T) {
		request := newGetScoreRequest("pujaz")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it records wins on POST", func (t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}

	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func testName(name string) string {
	return fmt.Sprintf("Test GET for %s", name)
}
