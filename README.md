# GoCommerce: E-commerce API

**GoCommerce** is an e-commerce backend API built in **Golang** using **MySQL** as the database. This project provides essential functionalities for an e-commerce platform, including user authentication, product management, and order handling, with a focus on role-based access control for different types of users (e.g., admin, customer).

## Features

- **User Authentication**: Register, log in, and manage users with JWT-based authentication.
- **Product Management**: CRUD operations for products, including listing, adding, editing, and deleting products.
- **Role-Based Access Control**: Admins have special access to protected routes for managing products and users.
- **Order Management**: Create, view, and manage orders for customers.

## Tech Stack

- **Golang**: Main programming language for the API.
- **MySQL**: Relational database for storing user, product, and order data.
- **GORM**: ORM for interacting with MySQL.
- **JWT**: Used for user authentication and session management.
- **Docker & Docker Compose**: Containerization for consistent development and deployment.

## Project Structure

```
/ecommerce-api
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── .env
├── main.go
├── /config                    # Configuration setup
│   └── config.go
├── /db                        # Database setup
│   └── db.go
├── /models                    # Database models
│   ├── user.go
│   └── product.go
├── /controllers               # API handlers/controllers
│   ├── user_controller.go
│   └── product_controller.go
├── /middleware                # Middleware for authentication/authorization
│   ├── auth.go
└── /utils                     # Utility functions
    ├── hash.go
    └── jwt.go
```

Feel free to submit issues or 
