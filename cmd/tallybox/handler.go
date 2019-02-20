package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	client "github.com/confio/credible-chain/client"
	"github.com/iov-one/weave/crypto"
)

type Application struct {
	Router  *chi.Mux
	Client  *client.CredibleClient
	Key     *crypto.PrivateKey
	ChainID string
	Port    int
}

func NewApplication(key *crypto.PrivateKey, remote string, port int) (*Application, error) {
	cc := client.NewRemoteClient(remote)
	chain, err := cc.ChainID()
	if err != nil {
		return nil, err
	}
	app := Application{
		Client:  cc,
		Key:     key,
		ChainID: chain,
		Port:    port,
	}
	app.initRouter()
	return &app, nil
}

func (a *Application) initRouter() {
	// Initialise a new router
	r := chi.NewRouter()
	r.Get("/", a.GetStatus)
	r.Get("/tally", a.ListTally)
	// r.Post("/vote", PostVote)

	// Log and apply to application
	// Log(LogModuleStartup, true, "Router initialised OK", nil)
	a.Router = r
}

func (a *Application) Serve() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", a.Port), a.Router)
}

// func ListCharities(w http.ResponseWriter, r *http.Request) {
// 	respond(w, app.Data.Charities)
// }

// func ListRecentVotes(w http.ResponseWriter, r *http.Request) {
// 	votes, err := getRecentVotes()
// 	if err != nil {
// 		respondWithError(w, errorTypeDatabase, err)
// 		return
// 	}

// 	respond(w, votes)
// }

func (a *Application) GetStatus(w http.ResponseWriter, r *http.Request) {
	status, err := a.Client.Status()
	if err != nil {
		render(w, 500, err.Error())
	}
	respond(w, status)
}

func (a *Application) ListTally(w http.ResponseWriter, r *http.Request) {
	// query tallies for all options
	tallies, err := a.Client.GetAllTallies()
	if err != nil {
		render(w, 500, err.Error())
	}
	// convert to a lookup table
	lookup := make(map[string]int64, len(tallies))
	for _, t := range tallies {
		lookup[t.Option] = t.Total
	}
	respond(w, lookup)
}

/***** helpers ****/

func render(w http.ResponseWriter, code int, toRender interface{}) {
	json, err := json.MarshalIndent(toRender, "", "  ")
	if err != nil {
		render(w, 500, err.Error())
		return
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return
}

func respond(w http.ResponseWriter, toRender interface{}) {
	render(w, http.StatusOK, toRender)
}
