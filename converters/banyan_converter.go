package converters

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/berryhill/ca-xlsx-conv/models"

	"github.com/tealeg/xlsx"
)

type Banyan struct {
	FileName 		string
	FileLocation	string
	FirstRow		[]string
	SecondRow		[]string
}

func NewBanyan(file_name string, file_location string) *Banyan {

	b := new(Banyan)
	b.FileName = file_name
	b.FileLocation = file_location

	b.FirstRow = append(b.FirstRow, "Load #")
	b.FirstRow = append(b.FirstRow, "Shipper Name")
	b.FirstRow = append(b.FirstRow, "Shipper Location Name")				// blank
	b.FirstRow = append(b.FirstRow, "Shipper Phone #")					// blank
	b.FirstRow = append(b.FirstRow, "Shipper Phone ext")					// blank
	b.FirstRow = append(b.FirstRow, "Shipper Email Address")				// blank
	b.FirstRow = append(b.FirstRow, "Shipper Address 1")
	b.FirstRow = append(b.FirstRow, "Shipper Address 2")					// blank
	b.FirstRow = append(b.FirstRow, "Shipper City")
	b.FirstRow = append(b.FirstRow, "Shipper State")
	b.FirstRow = append(b.FirstRow, "Shipper Zip")
	b.FirstRow = append(b.FirstRow, "Shipper Country")					// blank
	b.FirstRow = append(b.FirstRow, "Shipper Dock Name")					// blank
	b.FirstRow = append(														// blank
		b.FirstRow, "Shipper Accessorials  (reference accessorial " +
			"code tab)")
	b.FirstRow = append(b.FirstRow, "Pick Up Date")						// blank
	b.FirstRow = append(b.FirstRow, "Consignee Name")
	b.FirstRow = append(b.FirstRow, "Consignee Location Name")			// blank
	b.FirstRow = append(b.FirstRow, "Consignee Phone #")					// blank
	b.FirstRow = append(b.FirstRow, "Consignee Phone Ext")				// blank
	b.FirstRow = append(b.FirstRow, "Consignee Email Address")			// blank
	b.FirstRow = append(b.FirstRow, "Consignee Address 1")
	b.FirstRow = append(b.FirstRow, "Consignee Address 2")				// blank
	b.FirstRow = append(b.FirstRow, "Consignee City")
	b.FirstRow = append(b.FirstRow, "Consignee State")
	b.FirstRow = append(b.FirstRow, "Consignee Zip")
	b.FirstRow = append(b.FirstRow, "Consignee Country")					// blank
	b.FirstRow = append(b.FirstRow, "Consignee Dock Name")				// blank
	b.FirstRow = append(														// blank
		b.FirstRow, "Consignee Accessorials (reference accessorial " +
			"code tab)")
	b.FirstRow = append(b.FirstRow, "Delivery Date")						// blank
	b.FirstRow = append(b.FirstRow, "PO#")
	b.FirstRow = append(b.FirstRow, "BOL#")								// blank
	b.FirstRow = append(b.FirstRow, "BillingID")							// blank
	b.FirstRow = append(b.FirstRow, "Pro#")								// blank
	b.FirstRow = append(b.FirstRow, "Awarded Carrier SCAC")				// blank
	b.FirstRow = append(b.FirstRow, "Customer Charge")					// blank
	b.FirstRow = append(b.FirstRow, "Carrier Quote ID #")					// blank
	b.FirstRow = append(b.FirstRow, "Transit Time (days)")				// blank
	b.FirstRow = append(b.FirstRow, "Qty")
	b.FirstRow = append(b.FirstRow, "Weight")								// blank
	b.FirstRow = append(b.FirstRow, "Class")								// blank
	b.FirstRow = append(b.FirstRow, "Package Type")
	b.FirstRow = append(b.FirstRow, "Shipping Mode")						// blank
	b.FirstRow = append(b.FirstRow, "Shipping Qty")						// blank
	b.FirstRow = append(b.FirstRow, "Shipping Package Type")				// blank
	b.FirstRow = append(b.FirstRow, "Addtional Weight")					// blank
	b.FirstRow = append(b.FirstRow, "Equipment Type")						// blank
	b.FirstRow = append(b.FirstRow, "Special Instructions")				// blank
	b.FirstRow = append(b.FirstRow, "Shipper Contact First Name")			// blank
	b.FirstRow = append(b.FirstRow, "Shipper Contact Last Name")			// blank
	b.FirstRow = append(b.FirstRow, "Shipper Contact Fax")				// blank
	b.FirstRow = append(b.FirstRow, "Shipper Note")						// blank
	b.FirstRow = append(b.FirstRow, "Shipper Dock Note")					// blank
	b.FirstRow = append(b.FirstRow, "Shipper Dock Open Time")				// blank
	b.FirstRow = append(b.FirstRow, "Shipper Dock Pick up Time")			// blank
	b.FirstRow = append(b.FirstRow, "Shipper Dock Close Time")			// blank
	b.FirstRow = append(b.FirstRow, "Shipper Pickup Number")				// blank
	b.FirstRow = append(b.FirstRow, "Consignee Contact First Name")		// blank
	b.FirstRow = append(b.FirstRow, "Consignee Contact Last Name")		// blank
	b.FirstRow = append(b.FirstRow, "Consignee Contact Fax")				// blank
	b.FirstRow = append(b.FirstRow, "Consignee Note")						// blank
	b.FirstRow = append(b.FirstRow, "Consignee Dock Note")				// blank
	b.FirstRow = append(b.FirstRow, "Consignee Dock Open Time")			// blank
	b.FirstRow = append(b.FirstRow, "Consignee Dock Delivery Time")		// blank
	b.FirstRow = append(b.FirstRow, "Consignee Dock Close Time")			// blank
	b.FirstRow = append(b.FirstRow, "Consignee Delivery Number")			// blank
	b.FirstRow = append(														// blank
		b.FirstRow, "I Am The (Shipper, Consignee, or ThirdParty)")
	b.FirstRow = append(b.FirstRow, "Pay Type (Collect or Prepaid)")		// blank
	b.FirstRow = append(b.FirstRow, "Company Name")						// blank
	b.FirstRow = append(b.FirstRow, "Company Data Note")					// blank
	b.FirstRow = append(b.FirstRow, "Billing Address Name")				// blank
	b.FirstRow = append(b.FirstRow, "Billing Address 1")					// blank
	b.FirstRow = append(b.FirstRow, "Billing Address 2")					// blank
	b.FirstRow = append(b.FirstRow, "Billing City")						// blank
	b.FirstRow = append(b.FirstRow, "Billing State")						// blank
	b.FirstRow = append(b.FirstRow, "Billing Zip")						// blank
	b.FirstRow = append(b.FirstRow, "Billing Country")					// blank
	b.FirstRow = append(b.FirstRow, "Billing Contact Phone")				// blank
	b.FirstRow = append(b.FirstRow, "Billing Contact Fax")				// blank
	b.FirstRow = append(b.FirstRow, "Billing Contact Email")				// blank
	b.FirstRow = append(b.FirstRow, "Invoice ID#")						// blank
	b.FirstRow = append(b.FirstRow, "Raw Charge")							// blank
	b.FirstRow = append(b.FirstRow, "Carrier Charge")						// blank
	b.FirstRow = append(b.FirstRow, "NMFC")								// blank
	b.FirstRow = append(b.FirstRow, "SKU")
	b.FirstRow = append(b.FirstRow, "Hazmat Y/N")							// blank
	b.FirstRow = append(b.FirstRow, "Description")						// blank
	b.FirstRow = append(b.FirstRow, "Length")								// blank
	b.FirstRow = append(b.FirstRow, "Width")								// blank
	b.FirstRow = append(b.FirstRow, "Height")								// blank
	b.FirstRow = append(														// blank
		b.FirstRow, "Shipment Accessorials (reference accessorial " +
			"code tab)")
	b.FirstRow = append(b.FirstRow, "Declared Liability $")				// blank
	b.FirstRow = append(b.FirstRow, "COD $")								// blank
	b.FirstRow = append(b.FirstRow, "Account ID")							// blank
	b.FirstRow = append(b.FirstRow, "")

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

	for i := 0; i < len(transactions); i++ {
		row_n := sheet.AddRow()

		package_type, quantity, err := parseQuantityAndPackage(
			transactions[i].Item, transactions[i].Qty, transactions[i].UM)
		if err != nil {
			fmt.Println("FUCKED")
		}

		cell := row_n.AddCell()													// Load #
		cell.Value = transactions[i].Num
		cell = row_n.AddCell()													// Shipper Name
		if transactions[i].Class == "Integrated" {
			cell.Value = "Integrated Engineers"
		} else if transactions[i].Customer == "Full Cycle Nutrients" {
			cell.Value = "Full Cycle Nutrients"
		} else if transactions[i].Customer == "KG Chemical" {
			cell.Value = "KG Chemical"
		} else {
			cell.Value = "Custom Ag"
		}
		cell = row_n.AddCell()													// Shipper Location Name
		if transactions[i].Class == "Integrated" {
			cell.Value = "Integrated Engineers"
		} else if transactions[i].Customer == "Full Cycle Nutrients" {
			cell.Value = "Full Cycle Nutrients"
		} else if transactions[i].Customer == "KG Chemical" {
			cell.Value = "KG Chemical"
		} else {
			cell.Value = "Custom Ag"
		}
		cell = row_n.AddCell()													// Shipper Phone #
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Phone ext
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Email Address
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Address 1
		cell.Value = "3430 S Willow Ave"
		cell = row_n.AddCell()													// Shipper Address 2
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper City
		cell.Value = "Fresno"
		cell = row_n.AddCell()													// Shipper State
		cell.Value = "CA"
		cell = row_n.AddCell()													// Shipper Zip
		cell.Value = "93725"
		cell = row_n.AddCell()													// Shipper Country
		cell.Value = "United States"
		cell = row_n.AddCell()													// Shipper Dock Name
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Accessorials  (reference accessorial code tab)
		cell.Value = ""
		cell = row_n.AddCell()													// Pick Up Date
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Name
		cell.Value = transactions[i].ShipToAddress1
		cell = row_n.AddCell()													// Consignee Location Name
		cell.Value = transactions[i].ShipToAddress1
		cell = row_n.AddCell()													// Consignee Phone #
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Phone Ext
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Email Address
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Address 1
		cell.Value = transactions[i].ShipToAddress2
		cell = row_n.AddCell()													// Consignee Address 2
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee City
		cell.Value = transactions[i].ShipToCity
		cell = row_n.AddCell()													// Consignee State
		cell.Value = transactions[i].ShipToState
		cell = row_n.AddCell()													// Consignee Zip
		cell.Value = transactions[i].ShipZip
		cell = row_n.AddCell()													// Consignee Country
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Dock name
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Accessorials (reference accessorial code tab)
		cell.Value = ""
		cell = row_n.AddCell()													// Delivery Date
		cell.Value = ""
		cell = row_n.AddCell()													// PO#
		cell.Value = transactions[i].PO
		cell = row_n.AddCell()													// BOL#
		cell.Value = ""
		cell = row_n.AddCell()													// BillingID
		cell.Value = ""
		cell = row_n.AddCell()													// Pro#
		cell.Value = ""
		cell = row_n.AddCell()													// Awarded Carrier SCAC
		cell.Value = ""
		cell = row_n.AddCell()													// Customer Charge
		cell.Value = ""
		cell = row_n.AddCell()													// Carrier Quote ID #
		cell.Value = ""
		cell = row_n.AddCell()													// Transit Time (days)
		cell.Value = ""
		// cell = row_n.AddCell()													// Qty
		// cell.Value = transactions[i].Qty
		cell = row_n.AddCell()													// Qty
		cell.Value = quantity
		cell = row_n.AddCell()													// Weight
		cell.Value = ""
		cell = row_n.AddCell()													// Class
		cell.Value = ""
		// cell = row_n.AddCell()													// Package Type
		// cell.Value = transactions[i].UM
		cell = row_n.AddCell()													// Package Type
		cell.Value = package_type
		cell = row_n.AddCell()													// Shipping Mode
		cell.Value = ""
		cell = row_n.AddCell()													// Shipping Qty
		cell.Value = ""
		// cell = row_n.AddCell()													// Shipping Package Type
		// cell.Value = ""
		cell = row_n.AddCell()													// Shipping Package Type
		cell.Value = ""
		cell = row_n.AddCell()													// Additional Weight
		cell.Value = ""
		cell = row_n.AddCell()													// Equipment Type
		cell.Value = ""
		cell = row_n.AddCell()													// Special Instructions
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Contact First Name
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Contact Last Name
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Contact Fax
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Note
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Dock Note
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Dock Open Time
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Dock Pick up Time
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Dock Close Time
		cell.Value = ""
		cell = row_n.AddCell()													// Shipper Pickup Number
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Contact First Name
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Contact Last Name
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Contact Fax
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Note
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Dock Note
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Dock Open Time
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Dock Delivery Time
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Dock Close Time
		cell.Value = ""
		cell = row_n.AddCell()													// Consignee Delivery Number
		cell.Value = ""
		cell = row_n.AddCell()													// I Am The (Shipper, Consignee, or ThirdParty)
		cell.Value = ""
		cell = row_n.AddCell()													// Pay Type (Collect or Prepaid)
		cell.Value = ""
		cell = row_n.AddCell()													// Company Name
		cell.Value = ""
		cell = row_n.AddCell()													// Company Data Note
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Address Name
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Address 1
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Address 2
		cell.Value = ""
		cell = row_n.AddCell()													// Billing City
		cell.Value = ""
		cell = row_n.AddCell()													// Billing State
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Zip
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Country
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Contact Phone
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Contact Fax
		cell.Value = ""
		cell = row_n.AddCell()													// Billing Contact Email
		cell.Value = ""
		cell = row_n.AddCell()													// Invoice ID#
		cell.Value = ""
		cell = row_n.AddCell()													// Raw Charge
		cell.Value = ""
		cell = row_n.AddCell()													// Carrier Charge
		cell.Value = ""
		cell = row_n.AddCell()													// NMFC
		cell.Value = ""
		// cell = row_n.AddCell()													// SKU
		// cell.Value = transactions[i].Item
		cell = row_n.AddCell()													// SKU
		cell.Value = transactions[i].Qty
		cell = row_n.AddCell()													// Hazmat Y/N
		cell.Value = ""
		cell = row_n.AddCell()													// Description
		cell.Value = transactions[i].Item
		cell = row_n.AddCell()													// Length
		cell.Value = ""
		cell = row_n.AddCell()													// Width
		cell.Value = ""
		cell = row_n.AddCell()													// Height
		cell.Value = ""
		cell = row_n.AddCell()													// Shipment Accessorials (reference accessorial code tab)
		cell.Value = ""
		cell = row_n.AddCell()													// Declared Liability $
		cell.Value = ""
		cell = row_n.AddCell()													// COD $
		cell.Value = ""
		cell = row_n.AddCell()													// Account ID
		cell.Value = ""
	}

	err = file.Save(b.FileLocation + b.FileName)
	if err != nil {
		return err
	}

	return nil
}

func parseQuantityAndPackage(
	item string, quantity string, um string) (
		package_type string, qty string, err error) {

	item_array := strings.Split(item, " ")

	if um == "ea" {
		for _, v := range(item_array) {
			if v == "1qt" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 12
				qty = strconv.Itoa(int(q))
				package_type = "cases"
			} else if v == "1pt" {
				quantity_int, _ := strconv.ParseFloat(quantity, 32)
				q :=  quantity_int / 20

				quantity_float, _ := strconv.ParseFloat(quantity, 64)
				if quantity_float % 20.0 != 0 {
					q++
				}

				qty = strconv.Itoa(int(q))
				package_type = "cases"
			}
		}
	} else if um == "gal" {
		for _, v := range(item_array) {
			if v == "1" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 4
				qty = strconv.Itoa(int(q))
				package_type = "cases"
			} else if v == "2.5" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 5
				qty = strconv.Itoa(int(q))
				package_type = "cases"
			} else if v == "5" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 5
				qty = strconv.Itoa(int(q))
				package_type = "pail"
			} else if v == "30" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 30
				qty = strconv.Itoa(int(q))
				package_type = "drums"
			} else if v == "55" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 55
				qty = strconv.Itoa(int(q))
				package_type = "drums"
			// } else if v == "265" { // todo: deprecate this..
			// 	q, _ := strconv.ParseInt(quantity, 10, 32)
			// 	qty = strconv.FormatInt(q / 270, 16)
			// 	package_type = "tote"
			} else if v == "270" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 270
				qty = strconv.Itoa(int(q))
				package_type = "tote"
			}
		}
	} else if um == "lb" {
		for _, v := range(item_array) {
			if v == "25" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 25
				qty = strconv.Itoa(int(q))
				package_type = "bag"
			} else if v == "50" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 50
				qty = strconv.Itoa(int(q))
				package_type = "bag"
			} else if v == "2000" {
				quantity_int, _ :=
					strconv.ParseInt(quantity, 10, 32)
				q :=  quantity_int / 2000
				qty = strconv.Itoa(int(q))
				package_type = "bag"
			}
		}
	} else if um == "ton" {

	} else {
		qty = "-1"
	}
	//fmt.Println(qty, package_type)

	return package_type, qty, err
}
