package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"strings"
)

type SpecimenTypeRecord struct {
	Display string `json:"display"`
	Code    string `json:"code"`
}

//go:embed specimentype.csv
var specimentypeCSV string

// loadICDData loads the loinc test data from the embedded CSV file into an array of SpecimenTypeRecord objects.
// The first record is skipped, as it is the header.
func loadSpecimenTypeData() []SpecimenTypeRecord {
	reader := csv.NewReader(strings.NewReader(specimentypeCSV))
	reader.Comma = ';' // Mengatur delimiter menjadi semicolon (;)

	records, err := reader.ReadAll()
	if err != nil {
		panic(fmt.Sprintf("Failed to read CSV: %v", err))
	}

	// Jika jumlah record <= 1 (hanya header atau kosong), kembalikan slice kosong
	if len(records) <= 1 {
		return []SpecimenTypeRecord{}
	}

	SpecimenTypeData := make([]SpecimenTypeRecord, len(records)-1)
	for i := 1; i < len(records); i++ {
		record := records[i]
		display := record[1]
		code := record[0]
		SpecimenTypeData[i-1] = SpecimenTypeRecord{
			Display: display,
			Code:    code,
		}
	}

	return SpecimenTypeData
}
