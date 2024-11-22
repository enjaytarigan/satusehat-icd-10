package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"strings"
)

type AlphanumericsRecord struct {
	Value         string `json:"display"`
	Code          string `json:"code"`
	TestLoincCode string `json:"loinc_test"`
	System        string `json:"system"`
}

//go:embed alphanumeric_satusehat.csv
var AlphanumericsCSV string

func loadAlphanumericData() []AlphanumericsRecord {
	reader := csv.NewReader(strings.NewReader(AlphanumericsCSV))
	reader.Comma = ','       // Mengatur delimiter
	reader.LazyQuotes = true // Mengabaikan error pada tanda kutip

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error parsing CSV: %v\n", err)
		panic(err)
	}

	if len(records) <= 1 {
		return []AlphanumericsRecord{}
	}

	var result []AlphanumericsRecord
	currentTestLoincCode := ""

	for i := 1; i < len(records); i++ {
		row := records[i]

		// Periksa apakah jumlah kolom mencukupi
		if len(row) < 8 {
			fmt.Printf("Skipping row %d: insufficient columns (%d found, 8 expected)\n", i+1, len(row))
			continue
		}

		// Ambil nilai dari kolom CSV
		value := row[5]
		code := row[6]
		system := row[7]
		testLoincCode := row[3]

		if testLoincCode == "" {
			testLoincCode = currentTestLoincCode
		} else {
			currentTestLoincCode = testLoincCode
		}

		if value != "" && code != "" && system != "" && testLoincCode != "" {
			result = append(result, AlphanumericsRecord{
				TestLoincCode: testLoincCode,
				Code:          code,
				Value:         value,
				System:        system,
			})
		}
	}

	return result
}
