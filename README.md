# goLang Todo Api 🚀⭐

This is a simple Todo API built using Golang.

## Technologies Used

- **Golang**: The primary language used for developing the API.
- **Chi**: Lightweight, fast, and flexible HTTP router for Go (Golang), used to define routes and middleware for web applications.
- **PostgreSQL**: Database used to store todo items.
- **Docker**: Containerization platform used to deploy the database.

## Installation

### Prerequisites

- Docker
- Docker Compose

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/RogerioPiatek/goLang-todoApi.git
   cd goLang-todoApi
   ```
2. Build and run the application using Docker Compose:
   ```bash
   docker-compose up -d
   ```
   - The API should now be accessible at http://localhost:9000.
   - The Database should be running on http://localhost:5432.

## API Endpoints

- **GET /**: Fetch all todo items.
- **GET /{id}**: Fetch a specific todo item.
- **POST /**: Create a new todo item.
- **PUT /{id}**: Update an existing todo item.
- **DELETE /{id}**: Delete a todo item.

## Testing the project
Feel free to use the `requests.http` file provided at the root folder to test the project.

You can interact with the file using the [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) Visual Studio Code extension.

## Configuration

You can configure the database connection and other settings by modifying the `docker-compose.yml` and `config-docker.toml` files.

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues.
