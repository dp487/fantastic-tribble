# Fantastic Tribble

A simple **RESTful Movie API** built with Go and Gorilla Mux. This project demonstrates basic CRUD operations for managing a collection of movies.

## Features

- Create, Read, Update, and Delete movies
- JSON-based API
- In-memory storage (no database)
- UUID-based movie IDs

## Tech Stack

- **Language**: Go (Golang)
- **Router**: Gorilla Mux
- **Data Format**: JSON

## Project Structure

```
fantastic-tribble/
├── main.go     # Main API logic and handlers
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Running the Server

```bash
# Clone the repository
git clone https://github.com/dp487/fantastic-tribble.git
cd fantastic-tribble

# Run the application
go run main.go
```

The server starts on **port 8080**.

## API Endpoints

| Method | Endpoint       | Description              |
|--------|----------------|--------------------------|
| GET    | `/movies`      | Get all movies           |
| GET    | `/movies/{id}` | Get a movie by ID        |
| POST   | `/movies`      | Create a new movie       |
| PUT    | `/movies/{id}` | Update a movie by ID     |
| DELETE | `/movies/{id}` | Delete a movie by ID     |

## Example Usage

### Get All Movies

```bash
curl http://localhost:8080/movies
```

### Create a Movie

```bash
curl -X POST http://localhost:8080/movies \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "123456",
    "title": "Inception",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }'
```

## License

MIT
