# Fiber Mongo API

A RESTful API built with Go Fiber and MongoDB that provides CRUD operations for managing users, tasks, and tags.

## Features

- RESTful API endpoints for Users, Tasks, and Tags resources
- MongoDB integration for data persistence
- Modular architecture with separation of concerns
- Support for nested relationships (tasks with tags)

## Prerequisites

- Go 1.24 or higher
- Docker installed on your system

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/lleoserrano/fiber_mongo.git
   cd fiber_mongo
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up MongoDB using Docker:
   ```bash
   docker run --name mongo -d -p 27017:27017 mongo
   ```
   This will start a MongoDB instance on port 27017. You can configure your own MongoDB instance if needed.

## Usage

Start the server:

```bash
go run main.go
```

The API will be available at `http://localhost:9001`

## API Endpoints

### Users

- `POST /v1/users` - Create a new user
- `GET /v1/users` - Get all users
- `GET /v1/users/:id` - Get a user by ID
- `PUT /v1/users/:id` - Update a user
- `DELETE /v1/users/:id` - Delete a user

### Tasks

- `POST /v1/tasks` - Create a new task
- `GET /v1/tasks` - Get all tasks
- `GET /v1/tasks/:id` - Get a task by ID
- `PUT /v1/tasks/:id` - Update a task
- `DELETE /v1/tasks/:id` - Delete a task

### Tags

- `POST /v1/tags` - Create a new tag
- `GET /v1/tags` - Get all tags
- `GET /v1/tags/:id` - Get a tag by ID

## Data Models

### User

```json
{
  "id": "[ObjectID]",
  "name": "string",
  "email": "string",
  "password": "string"
}
```

### Task

```json
{
  "id": "[ObjectID]",
  "title": "string",
  "description": "string",
  "tags": ["string"],
  "assign": "[ObjectID]",
  "done": "boolean"
}
```

### Tag

```json
{
  "id": "[ObjectID]",
  "name": "string"
}
```

## Project Structure

```
fiber_mongo/
├── db/              # Database operations
├── docs/            # JSON schema definitions
├── main.go          # Entry point
├── tags/            # Tag handlers and models
├── tasks/           # Task handlers and models
└── users/           # User handlers and models
```

## Technologies

- [Go Fiber](https://github.com/gofiber/fiber) - Web framework
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver) - MongoDB driver for Go
