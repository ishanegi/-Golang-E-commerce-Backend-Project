# -Golang-E-commerce-Backend-Project

Crafting dynamic backend with Gin, JWT auth, SQLC DB management, insights via SQLC queries &amp; APIs, concurrent multi-product orders. Showcase skills in structure, auth, data, concurrency.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This is a web application that allows users to manage products, customers, and orders. It provides an intuitive interface for creating, updating, and viewing various aspects of an e-commerce platform. 

## Features

- **Product Management**: Easily add, update, and view product details such as name, price, and quantity.
- **Customer Records**: Manage customer information including name, email, and contact details.
- **Order Processing**: Place and process orders for multiple products concurrently while ensuring consistency.
- **Analytics**: Gain insights into top customers, product ratings, and other relevant metrics.
- **Secure Authentication**: Utilize JWT-based authentication for secure access to the application.
- **Middleware Structure**: Learn about middleware usage for handling authentication, CORS, and more.
- **Database Integration**: Integrate with a PostgreSQL database for persistent storage of data.
- **API Endpoints**: Utilize well-defined API endpoints for seamless interaction with the application.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/) (v1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/) (v13 or higher)
- [Git](https://git-scm.com/)


## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/ishanegi/-Golang-E-commerce-Backend-Project.git
   ```

2. Navigate to the project directory:

   ```sh
   cd -Golang-E-commerce-Backend-Project
   ```

3. Install project dependencies using Go modules:

   ```sh
   go mod download
   ```

4. Configure the database connection:
   
   - Open the `internal/database/database.go` file.
   - Modify the `ConnectDB` function to provide your PostgreSQL connection details (username, password, dbname, etc.).

5. Create the required tables in your PostgreSQL database:

   - Use the provided SQL scripts in the `sql` folder to create the necessary tables.

6. Build and run the application:

   ```sh
   go run main.go
   ```

7. Access the application:

   Open your web browser and go to [http://localhost:8080](http://localhost:8080) to access the application.
