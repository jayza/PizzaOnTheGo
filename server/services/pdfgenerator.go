package services

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jayza/pizzaonthego/models"
	"github.com/jung-kurt/gofpdf/v2"
)

// GeneratePdfReceiptAndOutput generates our pdf by adding text and images to the page
// then saving it to a file (name specified in params).
func GeneratePdfReceiptAndOutput(order *models.Order, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 7, "Pizza on the Go", "0", 0, "TL", false, 0, "")
	pdf.Ln(7)
	pdf.CellFormat(190, 7, "The following is a receipt for your order", "0", 0, "TL", false, 0, "")
	pdf.Ln(20)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(60, 7, "Product", "B", 0, "L", false, 0, "")
	pdf.CellFormat(25, 7, "Size", "B", 0, "L", false, 0, "")
	pdf.CellFormat(25, 7, "Quantity", "B", 0, "L", false, 0, "")
	pdf.CellFormat(40, 7, "Unit Price", "B", 0, "L", false, 0, "")
	pdf.CellFormat(40, 7, "Total Price", "B", 0, "R", false, 0, "")
	pdf.Ln(15)
	pdf.SetFont("Arial", "", 8)
	// Product Name, Quantity, Unit Price, Total
	// Every column gets aligned according to its contents.
	var totalOrderPrice float64 = 0
	for _, lineItem := range order.LineItems {
		// Again, we need the `CellFormat()` method to create a visible
		// border around the cell. We also use the `alignStr` parameter
		// here to print the cell content either left-aligned or
		// right-aligned.
		totalOrderPrice += lineItem.UnitPrice * float64(lineItem.Quantity)
		pdf.CellFormat(60, 7, lineItem.Item.Name, "0", 0, "L", false, 0, "")
		pdf.CellFormat(25, 7, lineItem.Size.Name, "0", 0, "L", false, 0, "")
		pdf.CellFormat(25, 7, strconv.Itoa(lineItem.Quantity), "0", 0, "L", false, 0, "")
		pdf.CellFormat(40, 7, strconv.FormatFloat(lineItem.UnitPrice, 'f', 2, 64)+" SEK", "0", 0, "L", false, 0, "")
		pdf.CellFormat(40, 7, strconv.FormatFloat(lineItem.UnitPrice*float64(lineItem.Quantity), 'f', 2, 64)+" SEK", "0", 0, "R", false, 0, "")
		pdf.Ln(-1)
		if len(lineItem.Ingredients) > 0 {
			pdf.CellFormat(190, 7, "+ "+strconv.Itoa(len(lineItem.Ingredients))+" ingredients and "+lineItem.Variation.Name+" crust", "0", 0, "L", false, 0, "")
		} else {
			pdf.CellFormat(190, 7, "", "0", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
		if len(lineItem.Ingredients) > 0 {
			pdf.CellFormat(190, 7, "Special instructions: \""+lineItem.SpecialInstruction+"\"", "B", 0, "L", false, 0, "")
		} else {
			pdf.CellFormat(190, 7, "", "B", 0, "L", false, 0, "")
		}
		pdf.Ln(10)
	}
	pdf.CellFormat(190, 7, "Total: "+strconv.FormatFloat(totalOrderPrice, 'f', 2, 64)+" SEK", "0", 0, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(190, 7, "These items will be shipped to your address at:", "0", 0, "L", false, 0, "")
	pdf.Ln(10)
	pdf.Cell(190, 7, order.ShippingInformation.FirstName+" "+order.ShippingInformation.LastName)
	pdf.Ln(-1)
	pdf.Cell(190, 7, order.ShippingInformation.StreetAddress)
	pdf.Ln(-1)
	pdf.Cell(190, 7, order.ShippingInformation.ZipCode+" "+order.ShippingInformation.City)
	pdf.Ln(10)
	pdf.Cell(190, 7, "If there are any problems with your delivery, we will be in touch to your cell phone at: "+order.ShippingInformation.PhoneNumber)

	defer fmt.Println("Receipt was generated..", os.Getenv("RECEIPT_FILE_DIRECTORY")+filename)
	return pdf.OutputFileAndClose(os.Getenv("RECEIPT_FILE_DIRECTORY") + filename)
}
