package main

import (
	_ "embed"
	"encoding/csv"
	"strings"
)

type ICDRecord struct {
	Code string `json:"code"`
	Description string `json:"description"`
}

//go:embed icd10.csv
var icd10CSV string

// loadICDData loads the ICD-10 data from the embedded CSV file into an array of ICDRecord objects.
// The first record is skipped, as it is the header.
func loadICDData() []ICDRecord {
	reader := csv.NewReader(strings.NewReader(icd10CSV))

	records, _ := reader.ReadAll()


	icd10Data := make([]ICDRecord, len(records) - 1)
	
	for i := 1; i < len(records) - 1; i++ {
		record := records[i]
		code := record[0]
		description := record[1]
		icd10Data[i-1] = ICDRecord{code, description}
	}

	return icd10Data
}