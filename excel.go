package json2docs

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/xuri/excelize/v2"
)

func ExcelTabularGenerator(format []byte, data []byte) (error, string) {
	// Parse format and data JSON
	var formatData Format
	err := json.Unmarshal(format, &formatData)
	if err != nil {
		return err, ""
	}

	// Create a new Excel file
	file := excelize.NewFile()

	// Set header values and merge cells
	for _, field := range formatData.Header {
		cell1 := fmt.Sprintf("A%d", field.Line)
		cell2 := fmt.Sprintf("%s%d", GetColumnLetter(len(formatData.BodyHeader)), field.Line)
		file.SetCellValue("Sheet1", cell1, field.Text)
		file.MergeCell("Sheet1", cell1, cell2)
		// fmt.Println(cell1, cell2)
		style, err := file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
			},
		})
		if err != nil {
			fmt.Println(err)
		}
		if err := file.SetCellStyle("Sheet1", cell1, cell2, style); err != nil {
			fmt.Println(err)
		}
	}

	// Set body header values
	for _, field := range formatData.BodyHeader {
		cell := fmt.Sprintf("%s%d", GetColumnLetter(field.Index), len(formatData.Header)+1)
		file.SetCellValue("Sheet1", cell, field.Text)
		file.SetColWidth("Sheet1", GetColumnLetter(field.Index), GetColumnLetter(field.Index), float64(field.Width))
	}

	// Set body fields values
	var items [][]byte
	i := 0
	jsonparser.ArrayEach(data, func(item []byte, dataType jsonparser.ValueType, offset int, err error) {
		// fmt.Println(string(item))
		items = append(items, item)
		jsonparser.ObjectEach(item, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			// fmt.Println(string(key), ">>>", string(value))
			for _, field := range formatData.BodyFields {
				if field.Field == string(key) {
					// fmt.Println("================================")
					// fmt.Println(field.Index, " > ", field.Field, " = ", string(value))
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", GetColumnLetter(field.Index), i+2+len(formatData.Header)), string(value))
					file.SetColWidth("Sheet1", GetColumnLetter(field.Index), GetColumnLetter(field.Index), float64(field.Width))
					if IsDigit(string(value)) {
						style, err := file.NewStyle(&excelize.Style{
							Alignment: &excelize.Alignment{
								Horizontal: "right",
							},
						})
						if err != nil {
							fmt.Println(err)

						}
						c := fmt.Sprintf("%s%d", GetColumnLetter(field.Index), i+2+len(formatData.Header))
						if err := file.SetCellStyle("Sheet1", c, c, style); err != nil {
							fmt.Println(err)
						}
					}
				}

			}
			return nil
		})
		i++
	})

	// Set summary values
	for _, field := range formatData.Summary {
		c := fmt.Sprintf("%s%d", GetColumnLetter(field.Index), len(formatData.Header)+2+len(items)+field.Index)
		file.SetCellValue("Sheet1", c, field.Text)
		file.SetColWidth("Sheet1", GetColumnLetter(field.Index), GetColumnLetter(field.Index), float64(field.Width))
		if IsDigit(string(field.Text)) {
			style, err := file.NewStyle(&excelize.Style{
				Alignment: &excelize.Alignment{
					Horizontal: "right",
				},
			})
			if err != nil {
				fmt.Println(err)

			}

			if err := file.SetCellStyle("Sheet1", c, c, style); err != nil {
				fmt.Println(err)
			}
		}
	}

	// Save the Excel file
	filename := RandStringBytes(8) + ".xlsx"
	err = file.SaveAs(filename)
	if err != nil {
		return err, ""
	}

	return nil, filename
}

func GetColumnLetter(index int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if index <= 26 {
		return string(letters[index-1])
	}
	first := (index-1)/26 - 1
	second := (index - 1) % 26
	return string(letters[first]) + string(letters[second])
}
