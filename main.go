package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/berryhill/ca-xlsx-conv/converters"
)

func main() {

	/*
	PHASE_1
 	*/
	// TODO: Logs
	// TODO: Tests
	// TODO: Cross Compatibility -> Windows

	/*
	PHASE_2
 	*/
	// TODO: Sort out dirs/dirs
		// Where to put things
	// TODO: Daemonize
		// Will scan file periodically to see if change occurred
		// Docker?
	// TODO: Connect to FTP of file upload
	// TODO: Integrate with .xlsx export CRON

	/*
	PHASE_3
 	*/
	// TODO: Slack/Healthcheck
	// TODO: Up on the cloud

	qbEnvFile := flag.String("file", "", "Text to parse.")
	qbEnvFileName := flag.String(
		"fileName", "", "Text to parse.")
	qbEnvFileLocation := flag.String(
		"fileLocation", "", "Text to parse.")

	flag.Parse()
	if *qbEnvFile == "" {
		fmt.Printf(
			"Must provide file, for example: -file=qb.xlsx")
		fmt.Println()
		os.Exit(1)
	} else if *qbEnvFileName == "" {
		fmt.Printf(
			"Must provide a file name for exported .xclx, for " +
				"example: -fileName=bp.xlsx")
		fmt.Println()
		os.Exit(1)
	} else if *qbEnvFileLocation == "" {
		fmt.Printf(
			"Must provide a file location for exported .xclx, for " +
				"example: -fileLocation=$HOME/Desktop/")
		fmt.Println()
		os.Exit(1)
	}

	qc := converters.NewQuickbooksSheet(*qbEnvFile)
	qc.Parse()
	b := converters.NewBanyan(*qbEnvFileName, *qbEnvFileLocation)
	b.Parse(qc.QuickbooksTransactions)
}
