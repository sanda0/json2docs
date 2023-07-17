package json2docs

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/jung-kurt/gofpdf"
)

func PdfTabularGenerator(format []byte, data []byte) (error, string) {
	// Parse the format and data JSON arrays
	var formatData Format
	err := json.Unmarshal(format, &formatData)
	if err != nil {
		return err, ""
	}

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page
	pdf.AddPage()

	// Set font and font size for the header
	pdf.SetFont("Arial", "B", 16)

	// Print header lines
	for _, header := range formatData.Header {
		pdf.SetXY(10, float64(header.Line*10))
		pdf.CellFormat(0, 10, header.Text, "", 0, "C", false, 0, "")
	}

	// Set font and font size for the table
	pdf.SetFont("Arial", "", 12)

	// Set table header background color
	pdf.SetFillColor(240, 240, 240)

	// Calculate total table width
	totalWidth := 0
	for _, header := range formatData.BodyHeader {
		totalWidth += header.Width
	}
	for _, field := range formatData.BodyFields {
		totalWidth += field.Width
	}

	// Calculate available width on the page
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, _, rightMargin, _ := pdf.GetMargins()
	availableWidth := pageWidth - leftMargin - rightMargin

	// Adjust table width to fit the page
	if totalWidth > int(availableWidth) {
		scaleFactor := float64(availableWidth) / float64(totalWidth)
		for i := range formatData.BodyHeader {
			formatData.BodyHeader[i].Width = int(float64(formatData.BodyHeader[i].Width) * scaleFactor)
		}
		for i := range formatData.BodyFields {
			formatData.BodyFields[i].Width = int(float64(formatData.BodyFields[i].Width) * scaleFactor)
		}
	}

	tableCellWidth := availableWidth / float64(len(formatData.BodyFields))
	fmt.Println(tableCellWidth)
	// Print table header
	x := 10.0
	y := float64(len(formatData.Header)*10 + 20)
	for _, header := range formatData.BodyHeader {
		pdf.Rect(x, y, float64(tableCellWidth), 10, "F")
		pdf.SetXY(x, y)
		pdf.CellFormat(float64(tableCellWidth), 10, header.Text, "", 0, "C", false, 0, "")
		x += float64(tableCellWidth)
	}

	// Set body fields values
	var items [][]byte
	i := 0
	jsonparser.ArrayEach(data, func(item []byte, dataType jsonparser.ValueType, offset int, err error) {

		y += 10
		items = append(items, item)
		jsonparser.ObjectEach(item, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			x = 10
			for _, field := range formatData.BodyFields {
				if field.Field == string(key) {
					// fmt.Println("================================")
					// fmt.Println(field.Index, " > ", field.Field, " = ", string(value))
					pdf.Rect(x, y, float64(tableCellWidth), 10, "")
					pdf.SetXY(x, y)
					txt := string(value)
					if IsDigit(txt) {
						pdf.CellFormat(float64(tableCellWidth), 10, txt, "", 0, "R", false, 0, "")
					} else {
						pdf.CellFormat(float64(tableCellWidth), 10, txt, "", 0, "C", false, 0, "")
					}
				}
				x += float64(tableCellWidth)
			}

			return nil
		})

		i++
	})

	// Print summary
	x = 10.0
	for _, summary := range formatData.Summary {
		pdf.Rect(x, y, tableCellWidth, 10, "F")
		pdf.SetXY(x, y)
		if IsDigit(summary.Text) {
			pdf.CellFormat(tableCellWidth, 10, summary.Text, "", 0, "R", false, 0, "")
		} else {
			pdf.CellFormat(tableCellWidth, 10, summary.Text, "", 0, "C", false, 0, "")
		}
		x += tableCellWidth
	}

	// Save the PDF file
	filename := RandStringBytes(8) + ".pdf"
	err = pdf.OutputFileAndClose(filename)
	if err != nil {
		return err, ""
	}

	return nil, filename
}
