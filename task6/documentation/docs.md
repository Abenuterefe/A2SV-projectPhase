# 📝 Task Manager API Documentation

A RESTful API for managing tasks with user authentication, role-based access control (admin/user), and task tracking. Built using **Go (Gin)** and **MongoDB**, this backend supports JWT-based authentication, modular route separation, and clean service layers.

---

## 📚 Table of Contents

* [🧾 About the App](#-about-the-app)
* [🔐 Authentication](#-authentication)

  * [POST /auth/register](#post-authregister)
  * [POST /auth/login](#post-authlogin)
  * [POST /auth/promote](#post-authpromote)
* [🗂 Tasks](#-tasks)

  * [POST /tasks/](#post-tasks)
  * [GET /tasks/](#get-tasks)
  * [PUT /tasks/\:id](#put-tasksid)
  * [DELETE /tasks/\:id](#delete-tasksid)
* [⚠️ Errors](#-errors)

---

## 🧾 About the App

The **Task Manager API** allows users to register, log in, and view their tasks. Admin users have elevated privileges to:

* Create new tasks
* Update and delete any task
* Promote other users to admin

Regular users:

* Can register, log in, and view tasks assigned to them.

MongoDB is used as the database with the official MongoDB Go driver, and JWT tokens are used for authentication. Access is controlled using role-based middleware.

---

## 🔐 Authentication

### POST `/auth/register`

Registers a new user.

#### 📘 Description:

* Accepts a username and password.
* Password is hashed.
* Default role is `user`.

#### 🔸 Request Body:

```json
{
  "user_name": "abenezer123",
  "password": "securePassword"
}
```

#### ✅ Success Response:

```json
{
  "message": "User registered"
}
```

#### ❌ Error Responses:

```json
{
  "error": "Invalid input"
}
```

```json
{
  "error": "user already exists"
}
```

---

### POST `/auth/login`

Logs in a user and returns a JWT token.

#### 📘 Description:

* Accepts username and password.
* If correct, returns a JWT token for authenticated requests.

#### 🔸 Request Body:

```json
{
  "user_name": "abenezer123",
  "password": "securePassword"
}
```

#### ✅ Success Response:

```json
{
  "token": "<JWT token>"
}
```

#### ❌ Error Responses:

```json
{
  "error": "user not found"
}
```

```json
{
  "error": "invalid credentials"
}
```

---

### POST `/auth/promote`

Promotes an existing user to admin. **Only accessible by admins.**

#### 📘 Description:

* Requires admin token.
* Promotes a user by `user_name`.

#### 🔸 Headers:

```
Authorization: Bearer <admin-token>
```

#### 🔸 Request Body:

```json
{
  "user_name": "usernameToPromote"
}
```

#### ✅ Success Response:

```json
{
  "message": "User promoted to admin"
}
```

#### ❌ Error Responses:

```json
{
  "error": "Only admins can promote users"
}
```

```json
{
  "error": "user not found"
}
```

---

## 🗂 Tasks

### POST `/tasks/`

Creates a new task. **Admin only.**

#### 📘 Description:

* Admins can create new tasks.
* Assigned to a user (implicitly or explicitly).

#### 🔸 Headers:

```
Authorization: Bearer <admin-token>
```

#### 🔸 Request Body:

```json
{
  "title": "Build Gin App",
  "description": "Implement task routes"
}
```

#### ✅ Success Response:

```json
{
  "message": "Task created"
}
```

---

### GET `/tasks/`

Gets all tasks for the logged-in user.

#### 📘 Description:

* Regular users only see their own tasks.
* Admins can see all tasks if needed (optional extension).

#### 🔸 Headers:

```
Authorization: Bearer <token>
```

#### ✅ Success Response:

```json
[
  {
    "_id": "...",
    "title": "Build Gin App",
    "description": "Implement task routes",
    "completed": false,
    "created_at": "2025-07-28T12:00:00Z"
  }
]
```

---

### PUT `/tasks/:id`

Updates a task. **Admin only.**

#### 📘 Description:

* Update title, description, or completion status.

#### 🔸 Headers:

```
Authorization: Bearer <admin-token>
```

#### 🔸 Request Body:

```json
{
  "title": "Build Gin App v2",
  "description": "Add update and delete logic",
  "completed": true
}
```

#### ✅ Success Response:

```json
{
  "message": "Task updated"
}
```

#### ❌ Error Response:

```json
{
  "error": "invalid task ID"
}
```

---

### DELETE `/tasks/:id`

Deletes a task. **Admin only.**

#### 📘 Description:

* Permanently removes the task with the given ID.

#### 🔸 Headers:

```
Authorization: Bearer <admin-token>
```

#### ✅ Success Response:

```json
{
  "message": "Task deleted"
}
```

#### ❌ Error Response:

```json
{
  "error": "invalid task ID"
}
```

---

## ⚠️ Errors

Common error responses:

```json
{
  "error": "Missing token"
}
```

```json
{
  "error": "Invalid token"
}
```

```json
{
  "error": "Admin access only"
}
```

```json
{
  "error": "Invalid input"
}
```

---

## 📌 Notes

* All protected routes require the `Authorization: Bearer <JWT>` header.
* Only users with role `admin` can create, update, or delete tasks.
* Users can only view their own tasks via `GET /tasks/`.
* Extendable to support task filters, search, or pagination.

---

**Author**: Task Manager App Team
**Last Updated**: July 28, 2025
