
# Event Booking Application

This is a simple event booking application built using Go as part of an Udemy course on Go basics. The application provides various routes to manage users, events, and registrations.

---

## Features
- User registration and authentication
- Create, update, and delete events
- Register for events
- View a list of events or details of a specific event
- API token-based authentication

---

## How to Install and Run Locally

### Prerequisites
- Go (version 1.16 or higher)

### Steps to Install

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Kamil-Budzik/golang-projects
   cd event-booking
   ```

2. **Install dependencies:**
   The project uses Go modules, so dependencies will be managed automatically.
   ```bash
   go mod tidy
   ```
  
3. **Run the application:**
   ```bash
   go run main.go
   ```

   By default, the application will start on `http://localhost:8080`.

4. **Testing with Postman:**
   - Import the provided Postman collection (included in the project) to test the various API endpoints.
   - Ensure that you have valid authentication tokens for requests that require authorization.
   - Example: Authorization: eyJlbWFpbCI6InRlc3QxQGV4YW1wbGUu

### User Management
- **Create User**  
  **POST** `/signup`  
  **Body:**
  ```json
  {
    "email": "test1@example.com",
    "password": "test"
  }
  ```
  Response: 201 Created or Error Message

- **Login User**  
  **POST** `/login`  
  **Body:**
  ```json
  {
    "email": "test1@example.com",
    "password": "test"
  }
  ```
  Response: JWT token

---

### Event Management
- **Create Event**  
  **POST** `/events`  
  Requires Authorization token in headers.  
  **Body:**
  ```json
  {
    "name": "Test Event",
    "description": "Description of the event",
    "location": "Event location",
    "dateTime": "2025-01-01T15:30:00.000Z"
  }
  ```
  Response: Event ID or Error Message

- **Get Events**  
  **GET** `/events`  
  Returns a list of events.

- **Get Event by ID**  
  **GET** `/events/{id}`  
  Returns details of a specific event.

- **Update Event**  
  **PUT** `/events/{id}`  
  Requires Authorization token in headers.  
  **Body:**
  ```json
  {
    "name": "Updated Event Name",
    "description": "Updated Event Description",
    "location": "Updated Location",
    "dateTime": "2025-01-01T15:30:00Z"
  }
  ```
  Response: Updated event details or Error Message

- **Delete Event**  
  **DELETE** `/events/{id}`  
  Requires Authorization token in headers.  
  Response: 204 No Content or Error Message

---

### Event Registration
- **Register for an Event**  
  **POST** `/events/{id}/register`  
  Requires Authorization token in headers.  
  Response: Registration details or Error Message

- **Cancel Registration**  
  **DELETE** `/events/{id}/register`  
  Requires Authorization token in headers.  
  Response: 204 No Content or Error Message
