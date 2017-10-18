package converters

import (
	"fmt"
	"strings"

	"github.com/berryhill/ca-xlsx-conv/models"

	"github.com/tealeg/xlsx"
)

type QuickbooksSheet struct {
	File 					*xlsx.File
	CustomerIndex 			[]int
	QuickbooksTransactions 	[]*models.QuickbooksTransaction
}

func NewQuickbooksSheet() *QuickbooksSheet {
	qbs := new(QuickbooksSheet)
	qbs.File = getQuickbooksFile()
	qbs.getCustomerIndexes()

	return qbs
}

func getQuickbooksFile() *xlsx.File {

	excelFileName := "qb.xlsx"
	xlsx_file, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	return xlsx_file
}

func (qbs *QuickbooksSheet) getCustomerIndexes() error {

	sheet := qbs.File.Sheets[0]
	for index, rows := range sheet.Rows {
		if rows.Cells[1].Value != "" {
			cell_array := strings.Split(rows.Cells[1].Value, " ")
			if cell_array[0] != "Total" {
				qbs.CustomerIndex = append(qbs.CustomerIndex, index)
			}
		}
	}

	// TODO: Implement error

	return nil
}

func (qbs *QuickbooksSheet) Parse() error {

	sheet := qbs.File.Sheets[0]

	for _, customer_i := range qbs.CustomerIndex {

		transaction := new(models.QuickbooksTransaction)
		transaction.Customer = sheet.Rows[customer_i].Cells[1].Value

		transaction.Date = sheet.Rows[customer_i + 1].Cells[4].Value
		transaction.Num = sheet.Rows[customer_i + 1].Cells[6].Value
		transaction.ShipToAddress1 = sheet.Rows[customer_i + 1].Cells[8].Value
		transaction.ShipToAddress2 = sheet.Rows[customer_i + 1].Cells[10].Value
		transaction.ShipToCity = sheet.Rows[customer_i + 1].Cells[12].Value
		transaction.ShipToState = sheet.Rows[customer_i + 1].Cells[14].Value
		transaction.ShipZip = sheet.Rows[customer_i + 1].Cells[16].Value
		transaction.PO = sheet.Rows[customer_i + 1].Cells[18].Value
		transaction.Item = sheet.Rows[customer_i + 1].Cells[20].Value
		transaction.Qty = sheet.Rows[customer_i + 1].Cells[22].Value
		transaction.UM = sheet.Rows[customer_i + 1].Cells[24].Value
		transaction.Class = sheet.Rows[customer_i + 1].Cells[26].Value

		fmt.Println()
		fmt.Println(transaction)

		qbs.QuickbooksTransactions = append(
			qbs.QuickbooksTransactions, transaction)
	}

	fmt.Println()
	fmt.Println(qbs.QuickbooksTransactions)

	// TODO: Impelement error

	return nil
}
