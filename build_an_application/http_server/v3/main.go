package server

import(
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	return 42, nil
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5050", server))
}
