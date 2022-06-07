/*
	Define runtime args
*/

package global

import (
	"fmt"
	"os"

	"github.com/dump-time/games-admin-server/log"
	"github.com/jessevdk/go-flags"
	"github.com/xuri/excelize/v2"
)

var CmdOpts struct {
	DaemonMode bool `short:"d" long:"daemon" description:"Running on daemon mode"`
	ConfigPath string `short:"c" long:"config" description:"The config file path" required:"true"`
	TeamAdminExcelPath string `long:"excel" description:"The team admin info excel file path" required:"true"`
}

func importTeamAdmin(excelFilePath string) error {
	f, err := excelize.OpenFile(excelFilePath)
	if err != nil {
		return err
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	// Extract data from excel
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	// Close the spreadsheet.
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func initFlag() {
	// init flags
	_, err := flags.ParseArgs(&CmdOpts, os.Args)
	if err != nil {
		os.Exit(-1)
	}

	if err := importTeamAdmin(CmdOpts.TeamAdminExcelPath); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
