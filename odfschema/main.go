package main

import (
	"github.com/ivpusic/golog"
	"github.com/ivpusic/neo"
	"github.com/ivpusic/neo-cors"
	"github.com/ivpusic/neo/middlewares/logger"
	"github.com/skratchdot/open-golang/open"
	"os"
	"rng/loader"
	"rng/odfschema/api"
	"sync"
	"time"
)

const schemaName = "OpenDocument-v1.2-os-schema.rng"

var (
	log = golog.GetLogger("application")
)

func main() {
	if file, err := os.Open(schemaName); err == nil {
		defer file.Close()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			app := neo.App()
			app.Use(logger.Log)
			app.Use(cors.Init())
			root, _ := loader.Load(file)
			log.Info("loaded", root)
			app.Serve("/", "./build/web")
			api.ServletRegister(app, root)
			app.Start()
			wg.Done()
		}(wg)
		go func() {
			time.Sleep(time.Duration(time.Millisecond * 200))
			open.Start("http://localhost:3000")
		}()
		wg.Wait()
	}
}
