package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"strings"
)

type LoincRecord struct {
	Display            string `json:"display"`
	Code               string `json:"code"`
	Specimen           string `json:"specimen"`
	Scale              string `json:"scale"`
	Examination_Result string `json:"examination_result"`
	Unit_Measure       string `json:"unit_measure"`
}

//go:embed LoincTest.csv
var loincCSV string

// loadICDData loads the loinc test data from the embedded CSV file into an array of LoincRecord objects.
// The first record is skipped, as it is the header.
func loadLoincData() []LoincRecord {
	reader := csv.NewReader(strings.NewReader(loincCSV))
	reader.Comma = ';' // Mengatur delimiter menjadi semicolon (;)

	records, err := reader.ReadAll()
	if err != nil {
		panic(fmt.Sprintf("Failed to read CSV: %v", err))
	}

	// Jika jumlah record <= 1 (hanya header atau kosong), kembalikan slice kosong
	if len(records) <= 1 {
		return []LoincRecord{}
	}

	loincTestData := make([]LoincRecord, len(records)-1)
	for i := 1; i < len(records); i++ {
		record := records[i]
		display := record[0]
		code := record[1]
		specimen := record[7]
		scale := record[4]
		examination_result := record[3]
		unit_measure := record[5]
		loincTestData[i-1] = LoincRecord{
			Display:            display,
			Code:               code,
			Specimen:           specimen,
			Scale:              scale,
			Examination_Result: examination_result,
			Unit_Measure:       unit_measure,
		}
	}

	return loincTestData
}
