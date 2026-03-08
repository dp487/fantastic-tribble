# A CRUD Movie API with Golang

## Overview

The Movie API is a RESTful service that allows users to manage a collection of movies. It provides functionalities to create, read, update, and delete movie entries. Each movie includes an ID, ISBN, title, and director information.

## Features

- Retrieve a list of all movies.
- Fetch details of a specific movie by ID.
- Create a new movie entry.
- Update an existing movie entry.
- Delete a movie entry by ID.

## Technologies Used

- Go (Golang)
- Gorilla Mux (for routing)
- UUID (for unique movie IDs)
- JSON (for data interchange)

## Base URL

```
http://localhost:8080
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/movie-api.git
   cd movie-api
   ```

2. Ensure you have Go installed. If not, download it from [golang.org](https://golang.org/dl/).

3. Install the required dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. The server will start at `http://localhost:8080`.

## API Documentation

### Movie Structure

A movie object has the following structure:

```json
{
  "id": "string", // Unique identifier for each movie
  "isbn": "string", // ISBN identifier for the movie
  "title": "string", // Title of the movie
  "director": {
    // Director information
    "firstname": "string", // Director's first name
    "lastname": "string" // Director's last name
  }
}
```

## Endpoints

### Get All Movies

```
GET /movies
```

#### Response

- **200 OK**: Returns a JSON array of all movies.

### Get a Movie by ID

```
GET /movies/{id}
```

#### Parameters

- **id**: The unique identifier of the movie.

#### Response

- **200 OK**: Returns the movie object if found.
- **404 Not Found**: If the movie with the specified ID does not exist.

### Create a New Movie

```
POST /movies
```

#### Request Body

```json
{
  "isbn": "string",
  "title": "string",
  "director": {
    "firstname": "string",
    "lastname": "string"
  }
}
```

#### Response

- **201 Created**: Returns the created movie object.
- **400 Bad Request**: If the request body is invalid.

### Update a Movie by ID

```
PUT /movies/{id}
```

#### Parameters

- **id**: The unique identifier of the movie to be updated.

#### Request Body

```json
{
  "isbn": "string",
  "title": "string",
  "director": {
    "firstname": "string",
    "lastname": "string"
  }
}
```

#### Response

- **200 OK**: Returns the updated movie object.
- **404 Not Found**: If the movie with the specified ID does not exist.
- **400 Bad Request**: If the request body is invalid.

### Delete a Movie by ID

```
DELETE /movies/{id}
```

#### Parameters

- **id**: The unique identifier of the movie to be deleted.

#### Response

- **200 OK**: Returns the remaining movies after deletion.
- **404 Not Found**: If the movie with the specified ID does not exist.

## Example Usage

### Get All Movies Example

```bash
curl -X GET http://localhost:8080/movies
```

### Create a Movie Example

```bash
curl -X POST http://localhost:8080/movies -H "Content-Type: application/json" -d '{
  "isbn": "123456",
  "title": "Inception",
  "director": {
    "firstname": "Christopher",
    "lastname": "Nolan"
  }
}'
```

## Contribution

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Notes

- Ensure the server is running on port 8080 before making requests.
