# Task Management API Documentation

## Overview

The Task Management API is a RESTful service for creating, reading, updating, and deleting tasks...

---

## Table of Contents

- [Getting Started](#getting-started)
- [MongoDB Configuration](#mongodb-configuration)
- [User Model](#user-model)
- [Task Model](#task-model)
- [API Endpoints](#api-endpoints)
- [Authentication & Authorization](#authentication--authorization)
- [Error Handling](#error-handling)
- [Example Usage](#example-usage)
- [Troubleshooting](#troubleshooting)

---

## Getting Started

### Prerequisites

- Go 1.18 or later  
- MongoDB (local or remote)  
- Optional: MongoDB Compass  

### Installation

```bash
git clone https://github.com/your-repo/task-api.git
cd task-api
go mod tidy
go run main.go
```
MongoDB Configuration

Database: taskdb

Collections: users, tasks

Automatically created if not present.

{
  "id": "ObjectID (string)",
  "username": "unique username",
  "password": "hashed password",
  "role": "admin | user"
}
