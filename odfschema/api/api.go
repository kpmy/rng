package api

import (
	"github.com/ivpusic/neo"
	"log"
	"rng/schema"
)

func ServletRegister(app *neo.Application, root schema.Guide) {
	app.Get("/api", func(ctx *neo.Ctx) {
		log.Println("api request")
		ctx.Res.Text("{}", 200)
	})
}
