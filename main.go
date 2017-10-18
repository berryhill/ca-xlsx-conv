package main

import (
	"github.com/berryhill/ca-xlsx-conv/converters"
)

func main() {

	quickbooks_converter := converters.NewQuickbooksSheet()
	quickbooks_converter.Parse()
}
