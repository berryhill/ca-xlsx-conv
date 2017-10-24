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

func remove(
	s []*models.QuickbooksTransaction, i int,
	) []*models.QuickbooksTransaction {

	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func (qbs *QuickbooksSheet) Parse() error {

	sheet := qbs.File.Sheets[0]
	var transactions []*models.QuickbooksTransaction

	for _, customer_i := range qbs.CustomerIndex {

		index := 1
		for {
			transaction := new(models.QuickbooksTransaction)
			transaction.Customer = sheet.Rows[customer_i].Cells[1].Value

			cell_array := strings.Split(
				sheet.Rows[customer_i + index].Cells[1].Value, " ")
			if cell_array[0] == "Total" {
				break
			}

			transaction.Date = sheet.Rows[customer_i + index].Cells[4].Value
			transaction.Num = sheet.Rows[customer_i + index].Cells[6].Value
			transaction.ShipToAddress1 =
				sheet.Rows[customer_i + index].Cells[8].Value
			transaction.ShipToAddress2 =
				sheet.Rows[customer_i + index].Cells[10].Value
			transaction.ShipToCity =
				sheet.Rows[customer_i + index].Cells[12].Value
			transaction.ShipToState =
				sheet.Rows[customer_i + index].Cells[14].Value
			transaction.ShipZip =
				sheet.Rows[customer_i + index].Cells[16].Value
			transaction.PO = sheet.Rows[customer_i + index].Cells[18].Value
			transaction.Item = sheet.Rows[customer_i + index].Cells[20].Value
			transaction.Qty = sheet.Rows[customer_i + index].Cells[22].Value
			transaction.UM = sheet.Rows[customer_i + index].Cells[24].Value
			transaction.Class = sheet.Rows[customer_i + index].Cells[26].Value

			if transaction.Item == "Freight Charges (Freight Charge)" ||
				transaction.Item == "Bill of Lading Charge (Bill of Lading)" ||
				transaction.Item == "Loading Charges" ||
				transaction.Item == "Bulk (Mixing & Packaging of)" ||
				transaction.Item == "CA Sales Tax (Sales Tax)" ||
				transaction.Item == "Credit Card Charge - MC" ||
				transaction.Item == "" {
				// do not append
			} else if (transaction.ShipToAddress2 == "") {
				// do not append
			} else if (transaction.UM == "") {
				// do not append
			} else {
				transactions = append(transactions, transaction)
			}

			index++
		}
	}

	fmt.Println(len(transactions))
	for _, trans := range(transactions) {
		fmt.Println(trans)
	}

	qbs.QuickbooksTransactions = transactions

	// TODO: Implement error

	return nil
}
