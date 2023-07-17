package json2docs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/buger/jsonparser"
)

func HtmlTabularGenerator(format []byte, data []byte) (error, string) {
	// Parse the format and data JSON arrays
	var formatData Format
	err := json.Unmarshal(format, &formatData)
	if err != nil {
		return err, ""
	}

	// Create the HTML table string
	tableHTML := `<style>
	table, th, td{
		border: 1px solid black;
		border-collapse: collapse;
	}
	</style>`
	// Add header lines
	tableHTML += `<p style="font-size:25px;text-align: center;font-family:'arial';">`
	for _, header := range formatData.Header {
		tableHTML += fmt.Sprintf(`<span>%s</span><br>`, header.Text)
	}
	tableHTML += "</p>"
	tableHTML += `<table width='100%' style="font-family:'arial';">`
	// Add table header
	tableHTML += "<tr style='background-color:rgb(220,220,220);color:#000'>"
	for _, header := range formatData.BodyHeader {
		tableHTML += fmt.Sprintf("<th style='font-size:18px;'>%s</th>", header.Text)
	}
	tableHTML += "</tr>"

	// Add table rows

	var items [][]byte
	i := 0
	jsonparser.ArrayEach(data, func(item []byte, dataType jsonparser.ValueType, offset int, err error) {
		// fmt.Println(string(item))
		items = append(items, item)
		var tds []string
		for td, _ := range formatData.BodyFields {
			tds = append(tds, fmt.Sprintf("%d", td))
		}
		tableHTML += "<tr>"

		jsonparser.ObjectEach(item, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			// fmt.Println(string(key), ">>>", string(value))
			for j, field := range formatData.BodyFields {
				if field.Field == string(key) {
					// fmt.Println("================================")

					if IsDigit(string(value)) {
						tds[j] = fmt.Sprintf("<td style='text-align: right' >%s</td>", value)
					} else {
						tds[j] = fmt.Sprintf("<td>%s</td>", value)

					}
				}

			}

			return nil
		})

		tableHTML += strings.Join(tds, "")

		tableHTML += "</tr>"
		i++
	})

	// Add summary row
	tableHTML += "<tr style='background-color:rgb(220,220,220);color:#000'>"
	for _, summary := range formatData.Summary {
		if IsDigit(string(summary.Text)) {
			tableHTML += fmt.Sprintf("<td style='text-align: right' >%s</td>", summary.Text)
		} else {
			tableHTML += fmt.Sprintf("<td>%s</td>", summary.Text)

		}
	}
	tableHTML += "</tr>"

	tableHTML += "</table>"

	// Save the HTML file
	filename := RandStringBytes(8) + ".html"

	err = ioutil.WriteFile(filename, []byte(tableHTML), 0644)
	if err != nil {
		return err, ""
	}

	return nil, filename
}
