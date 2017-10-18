package main

import (
	"github.com/berryhill/ca-xlsx-conv/converters"
)

func main() {
	/*
	PHASE_1
 	*/
	// TODO: Logs
	// TODO: Tests
	// TODO: Converters
	// TODO: Cross Compatibility -> Windows

	/*
	PHASE_2
 	*/
	// TODO: Sort out dirs/dirs
		// Where to put things
	// TODO: Daemonize
		// Will scan file periodically to see if change occurred

	/*
	PHASE_3
 	*/
	// TODO: Slack/Healthcheck

	quickbooks_converter := converters.NewQuickbooksSheet()
	quickbooks_converter.Parse()
}
