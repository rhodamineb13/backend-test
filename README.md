# backend-test
https://github.com/rhodamineb13/backend-test

## 1. Overview

This is a simple REST API project using Go and Gin web framework. The specifications are as follows:
- Go 1.24.6
- MySQL 8.0.45
- Redis 6.0.16

To run this project, type `go run main.go` on the terminal.

## 2. Features

Users can 
- Insert new product
- View all products
- Get products given an ID
- Update product data
- Delete product
- Get list of categories
- Get category by ID, and
- Insert new category

## 3. Details

This project is written using a handler-service-repository pattern, or sometimes called as N-tier architecture. It's chosen because of its clear separation of function of each layer. 

Handler layer works as the first layer that deals with users, as it receives request and sends response.

Service layer is where the core logic of our program is. The calculation, branches, and others are done in this layer.

Repository layer connects the project to outside infrastructure (mainly to store the data) such as database (SQL, Redis) or AWS.

Each layer doesn't directly 'shake hand' against each other, rather they use interface as the 'third party'. This adheres to the fifth SOLID principle; Dependency Inversion Principle which states that two layers must not be directly dependent of each other. This helps developers create unit test (using mocking) and alter the implementation in certain layer without breaking the implementation on another layer. For example, if I decide to use PostgreSQL DB connection on repository layer instead of MySQL, this wouldn't cause a breaking change on the service layer as the service layer is dependent on I...Repository interface which is implemented by ...repository.

In database I use integer as the primary key for both products and categories table. Compared to UUID which isn't good for indexing, using integer is faster, especially when we have a large dataset.

On products table, indexing is done on both `category_id` column and `id` column, with `category_id` is having more priority than `id` so that the search is done by looking up the `category_id` first.

I also implement caching so frequent lookup on data will not result in multiple call to database. Instead it will look up on the cache first on Redis using the given key. This reduces the request/response time.