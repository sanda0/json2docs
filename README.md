# json2docs

### install
```
go get github.com/sanda0/json2docs
```

### example
```go
package main

import (
	"fmt"
	"log"

	"github.com/sanda0/json2docs"
)

func main() {
	formatJSON := []byte(`{
		"Header": [
			{"Line": 1, "Text": "Abc Company"},
			{"Line": 2, "Text": "123 Main Street"},
			{"Line": 3, "Text": "Special report"},
			{"Line": 4, "Text": "Reporting period January to December 2018"}
		],
		"BodyHeader": [
			{"Index": 1, "Text": "Code", "Width": 10},
			{"Index": 2, "Text": "Name", "Width": 30},
			{"Index": 3, "Text": "Quantity", "Width": 10},
			{"Index": 4, "Text": "Price", "Width": 10},
			{"Index": 5, "Text": "Total", "Width": 10},
			{"Index": 6, "Text": "AVG", "Width": 10}
		],
		"BodyFields": [
			{"Index": 1, "Field": "prodId", "Width": 10},
			{"Index": 2, "Field": "description", "Width": 30},
			{"Index": 3, "Field": "qty", "Width": 10},
			{"Index": 4, "Field": "price", "Width": 10},
			{"Index": 5, "Field": "amount", "Width": 10},
			{"Index": 6, "Field": "avg", "Width": 10}

		],
		"Summary": [
			{"Index": 1, "Text": "", "Width": 10},
			{"Index": 2, "Text": "", "Width": 30},
			{"Index": 3, "Text": "", "Width": 10},
			{"Index": 4, "Text": "", "Width": 10},
			{"Index": 5, "Text": "150.00", "Width": 10},
			{"Index": 6, "Text": "10.00", "Width": 10}
		]
	}`)

	dataJSON := []byte(`[
		{
			"prodId": "1",
			"description": "Product 1",
			"qty": 1,
			"price": 10.00,
			"amount": 10.00,
			"avg": 5.00
		},
		{
			"prodId": "2",
			"description": "Product 2",
			"qty": 1,
			"price": 20.00,
			"amount": 20.00,
			"avg": 5.00

		},
		{
			"prodId": "3",
			"description": "Product 3",
			"qty": 1,
			"price": 30.00,
			"amount": 30.00,
			"avg": 5.00

		},
		{
			"prodId": "4",
			"description": "Product 4",
			"qty": 1,
			"price": 40.00,
			"amount": 40.00,
			"avg": 5.00

		},
		{
			"prodId": "5",
			"description": "Product 5",
			"qty": 1,
			"price": 50.00,
			"amount": 50.00,
			"avg": 5.00

		},
		{
			"prodId": "3",
			"description": "Product 3",
			"qty": 1,
			"price": 30.00,
			"amount": 30.00,
			"avg": 5.00

		},
		{
			"prodId": "4",
			"description": "Product 4",
			"qty": 1,
			"price": 40.00,
			"amount": 40.00,
			"avg": 5.00

		},
		{
			"prodId": "5",
			
			"qty": 1,
			"price": 50.00,
			"amount": 50.00,
			"description": "Product 5",
			"avg": 5.00

		}
	]`)

	err, excelFilename := json2docs.ExcelTabularGenerator(formatJSON, dataJSON)
	if err != nil {
		log.Fatal(err)
	}
	err, pdfFilename := json2docs.PdfTabularGenerator(formatJSON, dataJSON)
	if err != nil {
		log.Fatal(err)
	}
	err, htmlFilename := json2docs.HtmlTabularGenerator(formatJSON, dataJSON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Excel file generated:", excelFilename)
	fmt.Println("pdf file generated:", pdfFilename)
	fmt.Println("html file generated:", htmlFilename)

	// /================================================================

}
```
