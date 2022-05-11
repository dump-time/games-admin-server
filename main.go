package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"lixiao189/games-admin-server/global"
	"lixiao189/games-admin-server/log"
	"lixiao189/games-admin-server/router"
	"net"
	"os"
)

func main() {
	// Start server gracefully
	server := endless.NewServer(global.Config.Serv.Addr, router.R)

	// daemon mode
	if *global.DaemonMode {
		server.BeforeBegin = func(add string) {
			// stdout pid
			pid := os.Getpid()
			log.Info(fmt.Sprintf("Deamon started: %v", pid))
		}
	}

	// Start server
	if err := server.ListenAndServe(); err != nil {
		switch err.(type) {
		case *net.OpError:
			log.Warn(err)
		default:
			log.Fatal(err)
		}
	}

}
