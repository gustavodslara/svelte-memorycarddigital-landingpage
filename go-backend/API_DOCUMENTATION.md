# Memory Card Digital - Backend API Documentation

Base URL: `http://localhost:8000`

## Authentication

### Register
Create a new user account.

- **URL**: `/api/register`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**: User object (excluding password).

### Login
Authenticate a user and receive a JWT token.

- **URL**: `/api/login`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1Ni..."
  }
  ```

### Get Current User
Get details of the currently authenticated user.

- **URL**: `/api/user`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <token>`
- **Response**: User object.

---

## Subscription

### Get Subscription Status
Check the current subscription status of the authenticated user.

- **URL**: `/api/subscription`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <token>`
- **Response**:
  ```json
  {
    "subscription_status": "active",       // "active" or "inactive"
    "subscription_plan": "monthly",        // "monthly", "yearly", or "none"
    "subscription_expires_at": "2024-12-31T23:59:59Z" // ISO 8601 date or null
  }
  ```

### Subscribe
Subscribe to a plan (Mock implementation for now).

- **URL**: `/api/subscription`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer <token>`
- **Content-Type**: `application/json`
- **Body**:
  ```json
  {
    "plan": "monthly" // "monthly" or "yearly"
  }
  ```
- **Response**:
  ```json
  {
    "message": "Subscription successful",
    "subscription_status": "active",
    "subscription_plan": "monthly",
    "subscription_expires_at": "2024-01-02T..."
  }
  ```

---

## Game Saves

### Upload Save
Upload a game save file.

- **URL**: `/api/saves`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer <token>`
- **Body**: Multipart form data (file, game_id, etc.) - *See handlers/save.go for details*

### Get Saves
List all saves for the authenticated user.

- **URL**: `/api/saves`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <token>`

### Download Save
Download a specific save file.

- **URL**: `/api/saves/:id/download`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <token>`
