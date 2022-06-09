/*
	Define runtime args
*/

package global

import (
	"os"

	"github.com/jessevdk/go-flags"
)

var CmdOpts struct {
	DaemonMode bool `short:"d" long:"daemon" description:"Running on daemon mode"`
	ConfigPath string `short:"c" long:"config" description:"The config file path" required:"true"`
}

func initFlag() {
	// init flags
	_, err := flags.ParseArgs(&CmdOpts, os.Args)
	if err != nil {
		os.Exit(-1)
	}
}
