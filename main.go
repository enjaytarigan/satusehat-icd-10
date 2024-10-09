package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)


type server struct {
	icd10Data []ICDRecord	
}


const (
	defaultPageSize = 10
	defaultPageNum  = 0
)

// handleGetICD10 handles GET /icd10, returning the list of ICD-10
// records in JSON format. It supports pagination with the following
// query parameters:
//
//     page: the page number (0-indexed) to retrieve, defaults to 0
//
//     size: the number of records to retrieve per page, defaults to 10
//
// It also supports searching with the following query parameter:
//
//     search: a string to search in the description and code fields
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


func main() {
	srv := &server{
		icd10Data: loadICDData(),
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/icd10", srv.handleGetICD10)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(httpServer.ListenAndServe())
}