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

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {

	score, ok := s.scores[name]

	if !ok {
		return 0, errors.New("player not found")
	}

	return score, nil
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int {
			"Pepper": 20,
			"Floyd": 10,
		},
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

func testName(name string) string {
	return fmt.Sprintf("Test GET for %s", name)
}
