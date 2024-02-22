# README.md

# Order Management System

This is a simple order management system built with Go and Echo framework. It provides APIs for managing customers and orders.

## Features

- Create new customers
- Create new orders for customers

## APIs

### Create Customer

```
POST /customers
```

Request:

```json
{
  "name": "John Doe",
  "phone": "+123456789" 
}
```

Response:

```json
{
  "id": 1,
  "name": "John Doe",
  "phone": "+123456789"
}
```

### Create Order

```
POST /orders
```

Request: 

```json
{
    "customer_id": 1,
    "item": "bread",
    "amount": 500,
    "time": "0:00"
}
```

Response:

```json
{
    "id": 1,
    "customer_id": 1,
    "item": "bread",
    "amount": 500,
    "time": "0:00"
}
```



## Running Locally

```
go run main.go
```

It will start the server on port 8080.

## Deployment

The app can be easily deployed to any cloud platform like AWS, GCP, Azure etc. 

Some options:

### Docker

Build Docker image:

```
docker build -t orders .
```

Run container:

```
docker run -p 8080:8080 orders
```

### AWS Elastic Beanstalk

- Create an EB application 
- Deploy code from GitHub repo
- EB will handle building docker container and running it

### Google Cloud Run

- Build docker image
- Push to Google Container Registry
- Deploy image to Cloud Run

