package converters

import (
	"github.com/tealeg/xlsx"
	"github.com/berryhill/ca-xlsx-conv/models"
)

type Banyan struct {
	FirstRow		[]string
	SecondRow		[]string
}

func NewBanyan() *Banyan {
	b := new(Banyan)

	b.FirstRow = append(b.FirstRow, "")
	b.FirstRow = append(b.FirstRow, "Load #")
	b.FirstRow = append(b.FirstRow, "Shipper Name")
	b.FirstRow = append(b.FirstRow, "Shipper Address 1")
	b.FirstRow = append(b.FirstRow, "Shipper City")
	b.FirstRow = append(b.FirstRow, "Shipper State")
	b.FirstRow = append(b.FirstRow, "Shipper Zip")
	b.FirstRow = append(b.FirstRow, "")
	b.FirstRow = append(b.FirstRow, "Consignee Name")
	b.FirstRow = append(b.FirstRow, "Consignee Address 1")
	b.FirstRow = append(b.FirstRow, "Consignee City")
	b.FirstRow = append(b.FirstRow, "Consignee State")
	b.FirstRow = append(b.FirstRow, "Consignee Zip")
	b.FirstRow = append(b.FirstRow, "PO#")
	b.FirstRow = append(b.FirstRow, "Product")
	b.FirstRow = append(b.FirstRow, "QTY")
	b.FirstRow = append(b.FirstRow, "Package Type")
	b.FirstRow = append(b.FirstRow, "")

	b.SecondRow = append(b.SecondRow, "Date")
	b.SecondRow = append(b.SecondRow, "Num")
	b.SecondRow = append(b.SecondRow, "Shipper Name")
	b.SecondRow = append(b.SecondRow, "Shipper Address")
	b.SecondRow = append(b.SecondRow, "Shipper City")
	b.SecondRow = append(b.SecondRow, "Shipper State")
	b.SecondRow = append(b.SecondRow, "Shipper Zip Code")
	b.SecondRow = append(b.SecondRow, "Customer")
	b.SecondRow = append(b.SecondRow, "Ship To Address 1")
	b.SecondRow = append(b.SecondRow, "Ship To Address 2")
	b.SecondRow = append(b.SecondRow, "Ship To City")
	b.SecondRow = append(b.SecondRow, "Ship To State")
	b.SecondRow = append(b.SecondRow, "Ship Zip")
	b.SecondRow = append(b.SecondRow, "P.O.#")
	b.SecondRow = append(b.SecondRow, "Item")
	b.SecondRow = append(b.SecondRow, "Qty")
	b.SecondRow = append(b.SecondRow, "U/M")
	b.SecondRow = append(b.SecondRow, "Class")

	return b
}

func (b *Banyan) Parse(transactions []*models.QuickbooksTransaction) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	row := sheet.AddRow()
	for _, s := range(b.FirstRow) {
		cell := row.AddCell()
		cell.Value = s
	}
	row_2 := sheet.AddRow()
	for _, s := range(b.SecondRow) {
		cell := row_2.AddCell()
		cell.Value = s
	}

	for i := 0; i < len(transactions); i++ {
		row_n := sheet.AddRow()

		cell := row_n.AddCell()
		cell.Value = transactions[i].Date
		cell = row_n.AddCell()
		cell.Value = transactions[i].Num
		cell = row_n.AddCell()
		if transactions[i].Class == "Integrated" {
			cell.Value = "Integrated Engineers"
		} else if transactions[i].Customer == "Full Cycle Nutrients" {
			cell.Value = "Full Cycle Nutrients"
		} else if transactions[i].Customer == "KG Chemical" {
			cell.Value = "KG Chemical"
		} else {
			cell.Value = "Custom Ag"
		}
		cell = row_n.AddCell()
		cell.Value = "3430 S Willow Ave"
		cell = row_n.AddCell()
		cell.Value = "Fresno"
		cell = row_n.AddCell()
		cell.Value = "CA"
		cell = row_n.AddCell()
		cell.Value = "93725"
		cell = row_n.AddCell()
		cell.Value = transactions[i].Customer
		cell = row_n.AddCell()
		cell.Value = transactions[i].ShipToAddress1
		cell = row_n.AddCell()
		cell.Value = transactions[i].ShipToAddress2
		cell = row_n.AddCell()
		cell.Value = transactions[i].ShipToCity
		cell = row_n.AddCell()
		cell.Value = transactions[i].ShipToState
		cell = row_n.AddCell()
		cell.Value = transactions[i].ShipZip
		cell = row_n.AddCell()
		cell.Value = transactions[i].PO
		cell = row_n.AddCell()
		cell.Value = transactions[i].Item
		cell = row_n.AddCell()
		cell.Value = transactions[i].Qty
		cell = row_n.AddCell()
		cell.Value = transactions[i].UM
		cell = row_n.AddCell()
		cell.Value = transactions[i].Class
	}

	err = file.Save("Banyan.xlsx")
	if err != nil {
		return err
	}

	return nil
}
