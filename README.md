# Post Management API

A RESTful API for user registration, authentication, post creation, and commenting, built with Go, Gin, GORM, and PostgreSQL. The project supports Docker-based deployment for easy setup.

## Features

- User signup and login with JWT authentication
- Create, retrieve, and delete posts
- Add, retrieve, and delete comments on posts
- A post can have contain multiple comments
- Input validation for email, password, and contact number
- PostgreSQL database integration

## Project Structure

```
post_management/
  .env
  docker-compose.yml
  Dockerfile
  go.mod
  go.sum
  main.go
  database/
    database.go
  handlers/
    comment.go
    login.go
    post.go
    signup.go
  middlewares/
    checkAuth.go
  models/
    comment.go
    post.go
    user.go
  utils/
    checkInput.go
```

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/)

## Environment Variables

The `.env` file contains database connection details:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=task
```

> **Note:** When running with Docker Compose, the `DB_HOST` is set to `db` in the compose file.

## Running with Docker Compose

1. **Build and Start the Services**

   Open a terminal in the `post_management` directory and run:

   ```sh
   docker-compose up --build
   ```

   This will:
   - Start a PostgreSQL database container (`db`)
   - Build and run the Go API container (`app`)
   - Expose the API on [http://localhost:8080](http://localhost:8080)

2. **API Endpoints**

    > Postman Documentation - https://documenter.getpostman.com/view/37244551/2sB2x2HtkU 

   - `POST   /signup` — Register a new user
   - `POST   /login` — Login and receive a JWT token
   - `POST   /post` — Create a new post (requires JWT)
   - `GET    /posts` — Get all posts (requires JWT)
   - `DELETE /post/:id` — Delete a post (requires JWT)
   - `POST   /comment/:post_id` — Add a comment to a post (requires JWT)
   - `GET    /comments/:post_id` — Get comments for a post (requires JWT)
   - `DELETE /comment/:comment_id` — Delete a comment (requires JWT)

   > For protected endpoints, include the JWT token in the `Authorization` header as `Bearer <token>`.

3. **Stopping the Services**

   Press `Ctrl+C` in the terminal, then run:

   ```sh
   docker-compose down
   ```

## Development (without Docker)

1. Install Go and PostgreSQL locally.
2. Update `.env` with your local DB credentials.
3. Run the database and ensure it is accessible.
4. Install dependencies:

   ```sh
   go mod download
   ```

5. Run the application:

   ```sh
   go run main.go
   ```

