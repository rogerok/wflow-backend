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
        "/api/auth": {
            "post": {
                "description": "Auth User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Auth User",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.AuthForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TokenResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Refresh User token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh User token",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TokenResponse"
                        }
                    }
                }
            }
        },
        "/private/users": {
            "get": {
                "description": "Get users list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get users list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        },
        "/private/users/{id}": {
            "get": {
                "description": "Get user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.UserCreateForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "forms.AuthForm": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 8
                }
            }
        },
        "forms.Pseudonym": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                }
            }
        },
        "forms.Social": {
            "type": "object",
            "properties": {
                "instagram": {
                    "type": "string"
                },
                "telegram": {
                    "type": "string"
                },
                "tiktok": {
                    "type": "string"
                },
                "vk": {
                    "type": "string"
                }
            }
        },
        "forms.UserCreateForm": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "password",
                "passwordConfirm",
                "pseudonym",
                "socialLinks"
            ],
            "properties": {
                "bornDate": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 255
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "middleName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 8
                },
                "passwordConfirm": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 8
                },
                "pseudonym": {
                    "$ref": "#/definitions/forms.Pseudonym"
                },
                "socialLinks": {
                    "$ref": "#/definitions/forms.Social"
                }
            }
        },
        "models.Pseudonym": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "models.Social": {
            "type": "object",
            "properties": {
                "instagram": {
                    "type": "string"
                },
                "telegram": {
                    "type": "string"
                },
                "tiktok": {
                    "type": "string"
                },
                "vk": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                },
                "pseudonym": {
                    "$ref": "#/definitions/models.Pseudonym"
                },
                "socialLinks": {
                    "$ref": "#/definitions/models.Social"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "responses.CreateResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "responses.TokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:5000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Word-Flow app API",
	Description:      "Word-Flow API docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
