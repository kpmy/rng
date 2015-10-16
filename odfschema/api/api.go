package api

import (
	"github.com/ivpusic/neo"
	"github.com/kpmy/rng/schema"
	"log"
)

func ServletRegister(app *neo.Application, root schema.Guide) {
	app.Get("/api", func(ctx *neo.Ctx) (int, error) {
		log.Println("api request")
		return 200, nil
	})
}
