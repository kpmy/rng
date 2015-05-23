package main

import (
	"fmt"
	"github.com/ivpusic/neo"
	"github.com/ivpusic/neo-cors"
	"github.com/skratchdot/open-golang/open"
	"os"
	"rng/loader"
	"rng/odfschema/api"
	"sync"
	"time"
)

const schemaName = "OpenDocument-v1.2-os-schema.rng"

func main() {
	if file, err := os.Open(schemaName); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		fmt.Println("loaded", root)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			app := neo.App()
			app.Use(cors.Init())
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
