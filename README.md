# SATUSEHAT ICD-10 API

## What is it?

This is a simple API server that provides a list of ICD-10 diagnosis codes and their descriptions in JSON format.

## How to use it?

The API endpoint is `GET /api/icd10`. You can use your favorite HTTP client to access the API.

### Build binary

To build the binary, run `go build -o bin/satusehat-icd10-api` in the project root directory. The binary will be outputted in the `/bin` directory.

### Build docker image

To build the docker image, run `docker build -t satusehat-icd10-api .` in the project root directory.

### Run docker container

To run the docker container, run `docker run -p 8080:8080 satusehat-icd10-api` in the project root directory. The API server will listen on port 8080 on the host machine.

### Pull from docker hub

You can pull the docker image from docker hub by visiting this link: https://hub.docker.com/r/enjaytarigan/satusehat-icd10-api
