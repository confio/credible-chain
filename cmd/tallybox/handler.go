package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confio/credible-chain/x/votes"

	"github.com/go-chi/chi"

	"github.com/iov-one/weave/crypto"

	client "github.com/confio/credible-chain/client"
	"github.com/confio/credible-chain/queue"
)

type Application struct {
	Router  *chi.Mux
	Client  *client.CredibleClient
	Key     *crypto.PrivateKey
	Queue   *queue.Queue
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
	r.Post("/vote", a.PostVote)

	// Log and apply to application
	// Log(LogModuleStartup, true, "Router initialised OK", nil)
	a.Router = r
}

func (a *Application) Serve() error {
	// start queue and pipeline here
	p, err := queue.NewPipeline(a.Client, a.Key)
	if err != nil {
		return err
	}
	a.Queue = queue.NewQueue(500)
	// start the queue in the backround, and serve
	go a.Queue.Run(p)
	return http.ListenAndServe(fmt.Sprintf(":%d", a.Port), a.Router)
}

func (a *Application) PostVote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vr votes.VoteRecord
	err := decoder.Decode(&vr)
	if err != nil {
		render(w, 500, err.Error())
		return
	}

	// Make take
	task := queue.Task{Vote: &vr}
	err = a.Queue.Push(&task)
	if err != nil {
		render(w, 500, err.Error())
		return
	}
	respond(w, "Accepted")
}

func (a *Application) GetStatus(w http.ResponseWriter, r *http.Request) {
	stats := a.Queue.Stats()
	respond(w, stats)
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
