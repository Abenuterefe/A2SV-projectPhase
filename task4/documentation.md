# Task Manager API Documentation

## Overview

The Task Manager API is a simple RESTful service for managing tasks. It is built using Go and the Gin web framework. All data is stored in memory, making it suitable for testing, development, or learning purposes. No authentication or persistent storage is implemented.

---

## Table of Contents

- [Getting Started](#getting-started)
- [Task Model](#task-model)
- [API Endpoints](#api-endpoints)
  - [GET /tasks](#get-tasks)
  - [GET /tasks/:id](#get-tasksid)
  - [POST /tasks](#post-tasks)
  - [PUT /tasks/:id](#put-tasksid)
  - [DELETE /tasks/:id](#delete-tasksid)
- [Example Usage](#example-usage)
- [Limitations](#limitations)
- [Troubleshooting](#troubleshooting)

---

## Getting Started

### Prerequisites

- Go 1.18 or later

### Installation & Run

```bash
git clone https://github.com/your-username/task-manager
cd task-manager
go run main.go
```

## Task Model
The in-memory task data uses the following structure:
```type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}
```

Each task has:

ID: a unique string identifier

Title: short title of the task

Description: optional detailed information


## API END POINTS

## GET /tasks
Retrieve a list of all tasks.

Response:

200 OK with JSON array of tasks

```[
  {
    "id": "1",
    "title": "Buy groceries",
    "description": "Milk, Bread, Eggs"
  }
]
```

## GET /tasks/:id
Retrieve a specific task by its ID.

Parameters:

id (string) – the ID of the task

Response:

`200 OK` with task object

`404 Not Found` if task does not exist


## POST /tasks
Create a new task.

Request Body:

```{
  "title": "New Task",
  "description": "This is a new task"
}
```
## PUT /tasks/:id
Update an existing task by ID.

Request Body:
```{
  "title": "Updated Title",
  "description": "Updated description"
}
```

Response:

`200 OK` with updated task

`404 Not` Found if task does not exist

## DELETE /tasks/:id
Delete a task by ID.

Response:

`204` No Content on success

`404 Not Found` if task does not exist


## Example Usage

Add a Task
```curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Write Docs","description":"Write documentation for Go app"}'
```

## Get All Tasks 

`curl http://localhost:8080/tasks`

## Update a Task

```curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Task","description":"Changed desc"}'
```

## Delete Tasks 
`curl -X DELETE http://localhost:8080/tasks/1`

## Limitations
Tasks are stored in memory only — they will be lost when the app stops.

No authentication or user support.

No pagination or search functionality.

No data validation.

## Troubleshooting
Missing Task ID: Ensure the task you're trying to update/delete exists.

App Restart Loses Data: All data is held in memory; use a database for persistence.

Port Already in Use: Change the default port in `main.go`.