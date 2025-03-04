// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Barathrum54",
            "url": "linkedin.com/in/barathrum54",
            "email": "tahabdurmus0@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Responds with a simple \"Hello, World!\" message.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Root Endpoint",
                "responses": {
                    "200": {
                        "description": "Hello, World!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/examples": {
            "get": {
                "description": "Retrieves all examples from the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "examples"
                ],
                "summary": "Get All Examples",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Example"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new example in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "examples"
                ],
                "summary": "Create Example",
                "parameters": [
                    {
                        "description": "Example Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.CreateExampleRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/routes.CreateExampleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Example": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "Primary key for the record.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name field, required with a max length of 100 characters.",
                    "type": "string"
                }
            }
        },
        "routes.CreateExampleRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "The name of the example to be created.",
                    "type": "string"
                }
            }
        },
        "routes.CreateExampleResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "routes.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.2",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GoBo - Go Fiber Boilerplate",
	Description:      "A boilerplate application for building web services using Go and Fiber.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
