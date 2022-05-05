/*
	Define runtime args
*/

package global

import "flag"

// DaemonMode daemonMode run server in daemon mode
var DaemonMode = flag.Bool("d", false, "Run server in daemon mode")

// ConfigPath custom config file path
var ConfigPath = flag.String("config", "./config.yml", "The config file path")

func initFlag() {
	// init flags
	flag.Parse()
}
