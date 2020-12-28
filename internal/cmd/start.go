package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli"

	"github.com/1gkx/taskmanager/internal/conf"
	"github.com/1gkx/taskmanager/internal/router"
	"github.com/1gkx/taskmanager/internal/store"
	"github.com/1gkx/taskmanager/internal/template"
)

var Start = cli.Command{
	Name:        "start",
	Usage:       "Start web server",
	Description: `Description`,
	Action:      runWeb,
}

func runWeb(c *cli.Context) {

	if err := conf.Read(); err != nil {
		panic(err)
	}
	fmt.Printf("Config: %+v\n", conf.Cfg)
	if err := store.Initialize(); err != nil {
		panic(err)
	}

	// session.Start()
	template.InitTemplate()

	fmt.Printf("Server start at %s port\n", ":8000")
	if err := http.ListenAndServeTLS(":443", "conf/cert/cert.pem", "conf/cert/key.pem", router.NewRouter()); err != nil {
		log.Fatal(err)
	}

}
