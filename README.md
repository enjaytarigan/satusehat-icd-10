# SATUSEHAT ICD-10 API

## What is it?

This is a simple API server that provides a list of ICD-10 diagnosis codes and their descriptions in JSON format.

## How to use it?

The API endpoint is `GET /api/icd10`. You can use your favorite HTTP client to access the API.

### Build binary

To build the binary, run `go build -o bin/satusehat-icd10-api` in the project root directory. The binary will be outputted in the `/bin` directory.

## How to run it?

Build the binary and run it. The API server will listen on port 8080.
