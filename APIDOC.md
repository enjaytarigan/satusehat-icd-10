# API Documentation

The API provides a list of ICD-10 diagnosis codes and their descriptions in JSON format.

## Endpoints

### GET /api/icd10

Returns a list of ICD-10 diagnosis codes and their descriptions in JSON format.

#### Parameters

-   `page`: The page number (0-indexed) to retrieve, defaults to 0.
-   `size`: The number of records to retrieve per page, defaults to 10.
-   `search`: A string to search in the description and code fields.

#### Response

The response will be a JSON object with the following structure:
