## Task Manager 

A simple, modular RESTful API for task management, built using Go, the Gin web framework, and MongoDB as a persistent data store.

 ##  Features
-Create, Read, Update, and Delete (CRUD) operations on tasks

-Persistent storage using MongoDB (not in-memory)

-Clean MVC architecture

-Uses Go contexts for DB calls with timeouts

-Built using idiomatic Go

-Simple and extensible codebase

## Tech Stack

This application uses the following tools:

Go – programming language

Gin – fast, minimalist HTTP web framework

MongoDB – NoSQL document database

MongoDB Go Driver – official driver to connect and interact with MongoDB

JSON – data format for input/output

MVC Architecture – separates concerns cleanly across the app

## API Endpoints


Here are the main REST endpoints:
## post task

POST /tasks — Create a new task

`request`
```{
  "title": "Learn MongoDB Go Driver",
  "completed": false
}
```
`Success Response` – 200 OK

response 

```{
  "id": "64cc3c92f1a3490012d37e0d",
  "title": "Learn MongoDB Go Driver",
  "completed": false,
  "created_at": "2025-07-21T12:30:00Z",
  "updated_at": "2025-07-21T12:30:00Z"
}
```

`Error Response` – 400 Bad Request
```{
  "error": "Invalid request payload"
}
```
## get all task

GET /tasks — Retrieve all tasks

`Request`

No request body required.
Success Response – `200 OK`

```[
  {
    "id": "64cc3c92f1a3490012d37e0d",
    "title": "Learn MongoDB Go Driver",
    "completed": false,
    "created_at": "2025-07-21T12:30:00Z",
    "updated_at": "2025-07-21T12:30:00Z"
  },
  {
    "id": "64cc3c92f1a3490012d37e0e",
    "title": "Write documentation",
    "completed": true,
    "created_at": "2025-07-20T18:45:00Z",
    "updated_at": "2025-07-20T19:00:00Z"
  }
]
```
Error Response – `500 Internal Server Error`

```{
  "error": "Failed to fetch tasks"
}
```

GET /tasks/:id — Get a specific task by ID

Request
No request body — the `:id` is passed in the URL.

Example:
`GET /tasks/64cc3c92f1a3490012d37e0d`

Success Response – `200 OK`
```{
  "id": "64cc3c92f1a3490012d37e0d",
  "title": "Learn MongoDB Go Driver",
  "completed": false,
  "created_at": "2025-07-21T12:30:00Z",
  "updated_at": "2025-07-21T12:30:00Z"
}
```

Error Response – `404 Not Found`

```{
  "error": "Task not found"
}
```
PUT /tasks/:id — Update a task by ID

Request
Example:
`PUT /tasks/64cc3c92f1a3490012d37e0d`

```{
  "title": "Learn Go MongoDB Driver - Updated",
  "completed": true
}
```

Success Response – `200 OK`

```{
  "id": "64cc3c92f1a3490012d37e0d",
  "title": "Learn Go MongoDB Driver - Updated",
  "completed": true,
  "created_at": "2025-07-21T12:30:00Z",
  "updated_at": "2025-07-21T13:45:00Z"
}
```

Error Response – `400 Bad Request`
```{
  "error": "Invalid task ID"
}
```
or 
```{
  "error": "Task not found"
}
```

DELETE /tasks/:id — Delete a task by ID

Request
No request body — just the ID in the URL.
Example:
`DELETE /tasks/64cc3c92f1a3490012d37e0d`

Success Response – `200 OK`
```{
  "message": "Task deleted successfully"
}
```

Error Response – `404 Not Found`

```{
  "error": "Task not found"
}
```

All endpoints use JSON for input and output.