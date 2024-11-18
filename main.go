package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/rs/cors"
)

type server struct {
	icd10Data        []ICDRecord
	loincData        []LoincRecord
	specimenTypeData []SpecimenTypeRecord
	alphanumericData []AlphanumericsRecord
}

const (
	defaultPageSize = 10
	defaultPageNum  = 0
)

// handleGetICD10 handles GET /icd10, returning the list of ICD-10
// records in JSON format. It supports pagination with the following
// query parameters:
//
//	page: the page number (0-indexed) to retrieve, defaults to 0
//
//	size: the number of records to retrieve per page, defaults to 10
//
// It also supports searching with the following query parameter:
//
//	search: a string to search in the description and code fields
func (s *server) handleGetICD10(w http.ResponseWriter, r *http.Request) {
	pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNum = defaultPageNum
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		pageSize = defaultPageSize
	}

	searchTerm := r.URL.Query().Get("search")

	var records []ICDRecord
	if searchTerm != "" {
		records = searchICD10(s.icd10Data, searchTerm)
	} else {
		records = s.icd10Data
	}

	start := pageNum * pageSize
	end := start + pageSize

	if end > len(records) {
		end = len(records)
	}

	responseWithJSON(w, 200, records[start:end])
}

func searchICD10(records []ICDRecord, searchTerm string) []ICDRecord {
	result := make([]ICDRecord, 0, len(records))

	for _, record := range records {
		if strings.Contains(record.Code, searchTerm) ||
			strings.Contains(record.Description, searchTerm) {
			result = append(result, record)
		}
	}

	return result
}

func (s *server) handleGetLoinc(w http.ResponseWriter, r *http.Request) {
	pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNum = defaultPageNum
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		pageSize = defaultPageSize
	}

	searchTerm := r.URL.Query().Get("search")

	var records []LoincRecord
	if searchTerm != "" {
		records = searchLoinc(s.loincData, searchTerm)
	} else {
		records = s.loincData
	}

	start := pageNum * pageSize
	end := start + pageSize

	if end > len(records) {
		end = len(records)
	}

	responseWithJSON(w, 200, records[start:end])
}

func searchLoinc(records []LoincRecord, searchTerm string) []LoincRecord {
	result := make([]LoincRecord, 0, len(records))

	for _, record := range records {
		if strings.Contains(record.Code, searchTerm) ||
			strings.Contains(record.Display, searchTerm) {
			result = append(result, record)
		}
	}

	return result
}

func (s *server) handleGetSpecimenType(w http.ResponseWriter, r *http.Request) {
	pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNum = defaultPageNum
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		pageSize = defaultPageSize
	}

	searchTerm := r.URL.Query().Get("search")

	var records []SpecimenTypeRecord
	if searchTerm != "" {
		records = searchSpecimenType(s.specimenTypeData, searchTerm)
	} else {
		records = s.specimenTypeData
	}

	start := pageNum * pageSize
	end := start + pageSize

	if end > len(records) {
		end = len(records)
	}

	// Avoid out-of-range slicing
	if start > len(records) {
		start = len(records)
	}

	responseWithJSON(w, http.StatusOK, records[start:end])
}

// searchSpecimenType searches the Specimen Type records based on the search term.
func searchSpecimenType(records []SpecimenTypeRecord, searchTerm string) []SpecimenTypeRecord {
	result := make([]SpecimenTypeRecord, 0, len(records))

	for _, record := range records {
		if strings.Contains(strings.ToLower(record.Display), strings.ToLower(searchTerm)) ||
			strings.Contains(strings.ToLower(record.Code), strings.ToLower(searchTerm)) {
			result = append(result, record)
		}
	}

	return result
}

/** */

func (s *server) handleGetAlphanumeric(w http.ResponseWriter, r *http.Request) {
	pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		pageNum = defaultPageNum
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		pageSize = defaultPageSize
	}

	searchTerm := r.URL.Query().Get("search")

	var records []AlphanumericsRecord
	if searchTerm != "" {
		records = searchAlphanumerics(s.alphanumericData, searchTerm)
	} else {
		records = s.alphanumericData
	}

	start := pageNum * pageSize
	end := start + pageSize

	if end > len(records) {
		end = len(records)
	}

	// Avoid out-of-range slicing
	if start > len(records) {
		start = len(records)
	}

	responseWithJSON(w, http.StatusOK, records[start:end])
}

// searchSpecimenType searches the Specimen Type records based on the search term.
func searchAlphanumerics(records []AlphanumericsRecord, searchTerm string) []AlphanumericsRecord {
	result := make([]AlphanumericsRecord, 0, len(records))

	for _, record := range records {
		if strings.Contains(strings.ToLower(record.TestLoincCode), strings.ToLower(searchTerm)) {
			result = append(result, record)
		}
	}

	return result
}

func main() {
	srv := &server{
		icd10Data:        loadICDData(),
		loincData:        loadLoincData(),
		specimenTypeData: loadSpecimenTypeData(),
		alphanumericData: loadAlphanumericData(),
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/api/icd10", srv.handleGetICD10)
	mux.HandleFunc("/api/loinc", srv.handleGetLoinc) // Add LOINC endpoint
	mux.HandleFunc("/api/specimentype", srv.handleGetSpecimenType)
	mux.HandleFunc("/api/alphanumerics", srv.handleGetAlphanumeric)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	handler := cors.Default().Handler(mux)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(httpServer.ListenAndServe())
}
