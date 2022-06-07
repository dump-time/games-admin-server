package main

import (
	"fmt"
	"net"
	"os"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/log"
	_ "github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/router"
	"github.com/fvbock/endless"
)

func main() {
	// Start server gracefully
	server := endless.NewServer(global.Config.Serv.Addr, router.R)

	// daemon mode
	if global.CmdOpts.DaemonMode {
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
