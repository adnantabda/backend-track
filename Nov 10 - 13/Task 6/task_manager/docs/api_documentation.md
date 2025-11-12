# Task Management API Documentation

## Base URL

`http://localhost:8080`

## Endpoints

### GET /tasks
Get a list of all tasks.

- Response: 200 OK
- Response Body: JSON array of task objects

### GET /tasks/:id
Get details of a specific task by ID.

- Parameters:
  - `id`: Task ID
- Response: 200 OK
- Response Body: JSON task object
- Errors:
  - 404 Not Found if task does not exist

### POST /tasks
Create a new task.

- Request Body (JSON):
  - `title` (string, required): Title of the task
  - `description` (string, optional): Description of the task
  - `due_date` (string, optional, ISO 8601 format): Due date of the task
  - `status` (string, required): Status of the task

- Response: 201 Created
- Response Body: JSON task object with generated ID
- Errors:
  - 400 Bad Request if input is invalid or required fields are missing

### PUT /tasks/:id
Update a specific task by ID.

- Parameters:
  - `id`: Task ID
- Request Body (JSON):
  - `title` (string, required): Updated title of the task
  - `description` (string, optional): Updated description
  - `due_date` (string, optional, ISO 8601 format): Updated due date
  - `status` (string, required): Updated status

- Response: 200 OK
- Response Body: JSON updated task object
- Errors:
  - 400 Bad Request if input is invalid
  - 404 Not Found if task does not exist

### DELETE /tasks/:id
Delete a specific task by ID.

- Parameters:
  - `id`: Task ID
- Response: 204 No Content
- Errors:
  - 404 Not Found if task does not exist

## MongoDB Integration

- The API now uses MongoDB for persistent data storage.
- MongoDB connection URI is configured in `main.go`.
- Task IDs are MongoDB ObjectIDs represented as hex strings.
- Ensure MongoDB is running locally or update the URI accordingly.

## Notes
- Dates should be in ISO 8601 format (e.g., `2025-12-08T20:00:00Z`).
- Status can be any string representing the task state (e.g., "pending", "completed").
