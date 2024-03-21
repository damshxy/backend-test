# Management Product And Order

## :toolbox: Installation

1. Clone the repository

```
    git clone https://github.com/damshxy/backend-test.git
```

2. Adding Dependencies

```
    go mod tidy
```

3. Run the Applicaton

```
    go run main.go
```

## :open_book: Documentation

## How to access endpoint

### User & Auth

##### Endpoint for Get All Users

```
    localhost:3000/api/v1/users
```

#### Response :

```
    {
        "data": [
            {
                "id": 1,
                "username": "john doe",
                "password": "$2a$10$RXjTHPOYVz11sgPUb1BPVerx6Hhaz1UuzZm1/h/A7uWW5C0tr6rji"
            },
            {
                "id": 2,
                "username": "damshxy",
                "password": "$2a$10$cy1/WJMvPwgES5zOKsON8Od7C9OAuujbrKnS7SogHJ2xksyPyNA6q"
            }
        ],
        "message": "success",
        "status": 200
    }
```

#### Endpoint For Register

```
    localhost:3000/api/v1/register
```

#### Response :

```
    {
        "code": 200,
        "data": {
            "id": 2,
            "username": "damshxy",
            "password": "$2a$10$NYu.7LCDZoFLnlr4wuYYYOeCgQcoKSU9.HgsfX.ozt1oLiIBrPCrm"
        }
    }
```

#### Endpoint for Login

```
    localhost:3000/api/v1/login
```

#### Response :

```
    {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4gZG9lIn0.HxraOUUpkfqFgw6xEYeNwGwheOZiZkrCDGPlgofVz_g",
        "code": 200,
        "data": {
            "id": 1,
            "username": "john doe",
            "password": "$2a$10$RXjTHPOYVz11sgPUb1BPVerx6Hhaz1UuzZm1/h/A7uWW5C0tr6rji"
        }
    }
```

#### Endpoint For Reset Password

```
    localhost:3000/api/v1/resetPassword
```

#### Request :

```
    {
        "username": "john doe",
        "new_password": "john"
    }
```

#### Response :

```
    {
        "code": 200,
        "message": "Password reset successfully"
    }
```

#### Endpoint for Profile

```
    localhost:3000/api/v1/profile
```

#### Request Header :

```
    access_token : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4gZG9lIn0.HxraOUUpkfqFgw6xEYeNwGwheOZiZkrCDGPlgofVz_g
```

#### Response :

```
    {
        "data": {
            "id": 1,
            "username": "john doe",
            "password": "$2a$10$RXjTHPOYVz11sgPUb1BPVerx6Hhaz1UuzZm1/h/A7uWW5C0tr6rji"
        },
        "message": "User profile fetched successfully"
    }
```

#### Endpoint for Logout

```
    localhost:3000/api/v1/logout
```

#### Request Header :

```
    access_token : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4gZG9lIn0.HxraOUUpkfqFgw6xEYeNwGwheOZiZkrCDGPlgofVz_g
```

#### Response :

```
    {
        "message": "Logged out successfully",
        "user": {
            "username": "john doe"
        }
    }
```

### Product

#### Endpoint for Get All Product

```
    localhost:3000/api/v1/products
```

#### Response :

```
    {
        "data": [
            {
                "id": 3,
                "name": "Sepatu adidass",
                "price": 2000000,
                "quantity": 5
            },
            {
                "id": 4,
                "name": "Tas adidass",
                "price": 500000,
                "quantity": 5
            }
        ],
        "message": "Products fetched successfully",
        "status": 200
    }
```

#### Endpoint for Get Product By ID

```
    localhost:3000/api/v1/product/3
```

#### Response :

```
    {
        "data": {
            "id": 3,
            "name": "Sepatu adidass",
            "price": 2000000,
            "quantity": 5
        },
        "message": "Product fetched successfully",
        "status": 200
    }
```

#### Endpoint for Create Product

```
    localhost:3000/api/v1/product
```

#### Request :

```
    {
        "name": "Tas adidass",
        "price": 500000,
        "quantity": 5
    }
```

#### Response :

```
    {
        "data": {
            "id": 4,
            "name": "Tas adidass",
            "price": 500000,
            "quantity": 5
        },
        "message": "Product created successfully",
        "status": 200
    }
```

#### Endpoint for Update Product

```
    localhost:3000/api/v1/product/3
```

#### Request :

```
    {
       "name": "Baju Adidass",
        "price": 1000000,
        "quantity": 1
    }
```

#### Response :

```
    {
       "data": {
            "id": 0,
            "name": "Baju Adidass",
            "price": 1000000,
            "quantity": 1
        },
        "message": "Product updated successfully",
        "status": 200
    }
```

#### Endpoint for Delete Product

```
    localhost:3000/api/v1/product/3
```

#### Response :

```
    {
        "message": "Product deleted successfully",
        "status": 200
    }
```

### Order

#### Endpoint for Get All Order

```
    localhost:3000/api/v1/orders
```

#### Response :

```
    {
        "data": [
            {
                "id": 2,
                "quantity": 5,
                "total_price": 6000000,
                "status": "processed"
            }
        ],
        "message": "Orders fetched successfully",
        "status": 200
    }
```

#### Endpoint for Update process order

```
    localhost:3000/api/v1/order/2/process
```

#### Response :

```
    {
        "data": {
            "id": 2,
            "quantity": 5,
            "total_price": 6000000,
            "status": "processed"
        },
        "message": "Order processed successfully",
        "status": 200
    }
```

#### Endpoint for Update Complete Order

```
    localhost:3000/api/v1/order/2/complete
```

#### Response :

```
    {
        "data": {
            "id": 2,
            "quantity": 5,
            "total_price": 6000000,
            "status": "Completed"
        },
        "message": "Order completed successfully"
    }
```
