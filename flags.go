/*
	Define runtime args
*/

package main

import "flag"

// daemonMode run server in daemon mode
var daemonMode = flag.Bool("d", false, "Run server in daemon mode")

func init() {
	// init flags
	flag.Parse()
}
