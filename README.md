# Task Management System Backend

Welcome to the backend of our Task Management System. Designed for efficiency and clarity, this system serves as the core of a robust task tracking solution. Leveraging the speed of Go and the flexibility of MongoDB, our backend efficiently handles task operations, ensuring reliability and high performance.

## Key Features

- **Efficient Task Management**: Manages tasks with streamlined backend functionalities.
- **Full CRUD Operations**: Supports Create, Read, Update, and Delete (CRUD) operations for tasks.
- **RESTful API Design**: Follows standard REST principles for predictable and easy-to-use endpoints.
- **Input Validation**: Ensures data integrity with declarative struct tags and automatic validation.
- **Scalable Data Storage with MongoDB**: Utilizes MongoDB for flexible and scalable document storage.
- **Centralized Error Handling**: Provides consistent error responses across the API.

## Core Technologies

The backend is built using a range of powerful technologies:

- **Language**: [Go (Golang)](https://go.dev/doc/install) - Renowned for its efficiency and scalability in backend development.
- **Framework**: [Gin](https://gin-gonic.com/) - A HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance.
- **Database**: [MongoDB](https://www.mongodb.com/) - A source-available cross-platform document-oriented database program.
- **Driver**: [mongo-driver](https://github.com/mongodb/mongo-go-driver) - The official MongoDB driver for Go.

## Setup Instructions

### 1. Install Go

Download and install Go from [here](https://go.dev/doc/install).

### 2. Update Environment Variables

- Create a `.env.development` file in the root directory.
- Add the following variables (modify as needed):

```env
PORT=8082
MONGO_URI=mongodb://localhost:27017
DB_NAME=taskdb
```

### 3. Setting up Database

Ensure you have a MongoDB instance running. You can run it locally or use a cloud provider like MongoDB Atlas.


### 4. Running the Server

Start the server with:

```bash
go run main.go
```

The server will start on port `8082` (or the port specified in your `.env.development`).

### 5. Testing the API

You can test the API using `curl` or Postman.

**Create a Task:**
```bash
curl -X POST http://localhost:8082/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Gin", "status": "pending"}'
```

**Get All Tasks:**
```bash
curl http://localhost:8082/api/v1/tasks
```

---

Our Task Management System Backend is designed to be a lightweight yet powerful solution for tracking tasks. It demonstrates the use of modern Go patterns and MongoDB integration.
