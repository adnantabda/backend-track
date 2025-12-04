# Task Management API Documentation

## Base URL
`http://localhost:8080/api`

## Authentication
This API uses JWT (JSON Web Tokens) for authentication. Include the token in the `Authorization` header for protected routes.

### Register
```
POST /auth/register
```
Request body:
```json
{
    "username": "user1",
    "password": "password123"
}
```

### Login
```
POST /auth/login
```
Request body:
```json
{
    "username": "user1",
    "password": "password123"
}
```
Response:
```json
{
    "token": "jwt.token.here",
    "user": {
        "id": 1,
        "username": "user1",
        "role": "user"
    }
}
```

## Tasks

### Get All Tasks
```
GET /tasks
```
**Permissions**: All authenticated users

### Get Task by ID
```
GET /tasks/:id
```
**Permissions**: All authenticated users

### Create Task
```
POST /tasks
```
Request body:
```json
{
    "title": "Complete assignment",
    "description": "Finish the task management API"
}
```
**Permissions**: All authenticated users

### Update Task
```
PUT /tasks/:id
```
Request body:
```json
{
    "title": "Updated title",
    "description": "Updated description",
    "status": "completed"
}
```
**Permissions**: Task owner or Admin

### Delete Task
```
DELETE /tasks/:id
```
**Permissions**: Task owner or Admin

## Admin Endpoints

### Promote User to Admin
```
POST /users/:id/promote
```
**Permissions**: Admin only

## Error Responses
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Missing or invalid authentication token
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

## Environment Variables
- `JWT_SECRET`: Secret key for JWT signing (required in production)
- `PORT`: Port to run the server on (default: 8080)
