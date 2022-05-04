/*
	Define runtime args
*/

package global

import "flag"

// daemonMode run server in daemon mode
var DaemonMode = flag.Bool("d", false, "Run server in daemon mode")

func InitFlag() {
	// init flags
	flag.Parse()
}
