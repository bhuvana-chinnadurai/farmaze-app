// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/b2bclients": {
            "get": {
                "description": "Retrieve a list of all B2B clients' summaries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "B2B"
                ],
                "summary": "Get client summaries",
                "responses": {
                    "200": {
                        "description": "List of B2B client summaries",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.B2BClient"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/b2bclients/{client_id}": {
            "get": {
                "description": "Retrieve details of a specific B2B client by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "B2B"
                ],
                "summary": "Get client details by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "client_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "B2B client details",
                        "schema": {
                            "$ref": "#/definitions/model.B2BClient"
                        }
                    },
                    "400": {
                        "description": "Invalid client ID format",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Client not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logs in a user with the provided username and password, returning a JWT token upon success",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success: Token generated and returned",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request: Cannot parse request"
                    },
                    "401": {
                        "description": "Unauthorized: Invalid username or password"
                    },
                    "500": {
                        "description": "Internal Server Error: Failed to generate token"
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "Lists all orders, optionally filtered by client ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "List all orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID to filter orders",
                        "name": "client_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of orders",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.CreateOrderResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Creates a new order with the specified details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "Create Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order successfully created",
                        "schema": {
                            "$ref": "#/definitions/api.CreateOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/orders/{client_id}": {
            "get": {
                "description": "Gets orders specific to a given client ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get orders by client ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "client_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of orders for the client",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.CreateOrderResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid client ID format"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/procurements": {
            "get": {
                "description": "Lists all procurements for a given date",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "procurements"
                ],
                "summary": "List procurements by date",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Date for procurement listing",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of procurements for the specified date",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.ProcurementResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid date parameter"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Retrieves a list of all products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "List all products",
                "responses": {
                    "200": {
                        "description": "List of products",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Product"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "description": "Adds a new product to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Add a new product",
                "parameters": [
                    {
                        "description": "Product to add",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created product",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid request format"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/products/{id}": {
            "put": {
                "description": "Edits an existing product identified by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Edit a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product data to update",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated product",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or product ID"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "delete": {
                "description": "Deletes a product identified by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Successfully deleted product"
                    },
                    "400": {
                        "description": "Invalid product ID format"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/products/{product_id}": {
            "get": {
                "description": "Retrieves a product's details by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product details",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid product ID format"
                    },
                    "500": {
                        "description": "Failed to retrieve product by ID"
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers a new user with a username, password, and role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Registration Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.registerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success: User successfully registered",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request: Cannot parse request"
                    },
                    "500": {
                        "description": "Internal Server Error: Error while creating a new user"
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CreateOrderRequest": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ProductRequest"
                    }
                },
                "total_price": {
                    "type": "number"
                }
            }
        },
        "api.CreateOrderResponse": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ProductResponse"
                    }
                },
                "status": {
                    "$ref": "#/definitions/api.OrderStatus"
                },
                "total_price": {
                    "type": "number"
                }
            }
        },
        "api.OrderStatus": {
            "type": "string",
            "enum": [
                "ordered"
            ],
            "x-enum-varnames": [
                "Ordered"
            ]
        },
        "api.ProcurementResponse": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "product_name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "api.ProductRequest": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "api.ProductResponse": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "api.loginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.registerRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.B2BClient": {
            "type": "object",
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "contact_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "available_quantity": {
                    "type": "integer"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "unit": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
