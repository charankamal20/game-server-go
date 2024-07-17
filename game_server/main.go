package main

import (
	"net/http"

	"github.com/anthdm/hollywood/actor"
)


type GameServer struct {}

func newGameServer() actor.Receiver	{
	return &GameServer{}
}

func (s *GameServer) Receive(c *actor.Context) {

}

func (s *HTTPServer) handleWS(w http.ResponseWriter, r *http.Request) {

}

func main() {

	e, _ := actor.NewEngine()
	e.Spawn(newGameServer, "server")

}