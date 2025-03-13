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
        "/api/auth/logout": {
            "post": {
                "description": "Logout User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout User",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/auth/refresh": {
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
        "/private/books": {
            "get": {
                "description": "Get book by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Get book by id",
                "parameters": [
                    {
                        "description": "Query parameters for books list",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BooksQueryParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Book"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "CreateBook Book",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "CreateBook Book",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.BookCreateForm"
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
        },
        "/private/books/{id}": {
            "get": {
                "description": "GetBooksList",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "GetBooksList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                }
            }
        },
        "/private/goals": {
            "get": {
                "description": "Get goals list by book id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "summary": "Get goals list by book id",
                "parameters": [
                    {
                        "description": "Query parameters for goals list",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GoalsQueryParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Goals"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create goal for book Goals",
                "tags": [
                    "Goals"
                ],
                "summary": "CreateGoal Goals",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.GoalCreateForm"
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
        },
        "/private/goals/{id}": {
            "get": {
                "description": "Get goal by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "summary": "Get by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Goals id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Goals"
                        }
                    }
                }
            }
        },
        "/private/quotes": {
            "get": {
                "description": "Get random quote",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quotes"
                ],
                "summary": "Get random quote",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Quotes"
                        }
                    }
                }
            }
        },
        "/private/reports": {
            "post": {
                "description": "CreateReport Report",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Report"
                ],
                "summary": "CreateReport Report",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.ReportCreateForm"
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
        },
        "/private/statistics/goal/{id}": {
            "get": {
                "description": "get goal's activity statistics",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistics"
                ],
                "summary": "Get statistics by goal id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Goal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GoalStatistics"
                        }
                    }
                }
            }
        },
        "/private/statistics/user": {
            "get": {
                "description": "get user's activity statistics",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistics"
                ],
                "summary": "Get statistics by user id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserStatistics"
                        }
                    }
                }
            }
        },
        "/private/statistics/user/full": {
            "get": {
                "description": "Get profile full chart data user id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistics"
                ],
                "summary": "Get profile full chart data user id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FullProfileChartData"
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
                "description": "CreateUser User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "CreateUser User",
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
        "forms.BookCreateForm": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 2
                },
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                }
            }
        },
        "forms.GoalCreateForm": {
            "type": "object",
            "required": [
                "bookId",
                "endDate",
                "goalWords",
                "startDate",
                "title"
            ],
            "properties": {
                "bookId": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 2
                },
                "endDate": {
                    "type": "string"
                },
                "goalWords": {
                    "type": "number",
                    "minimum": 2
                },
                "startDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 2
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
        "forms.ReportCreateForm": {
            "type": "object",
            "required": [
                "bookId",
                "goalId",
                "wordsAmount"
            ],
            "properties": {
                "bookId": {
                    "type": "string"
                },
                "goalId": {
                    "type": "string"
                },
                "wordsAmount": {
                    "type": "number",
                    "minimum": 2
                }
            }
        },
        "forms.Social": {
            "type": "object",
            "properties": {
                "instagram": {
                    "type": "string",
                    "maxLength": 250
                },
                "telegram": {
                    "type": "string",
                    "maxLength": 250
                },
                "tiktok": {
                    "type": "string",
                    "maxLength": 250
                },
                "vk": {
                    "type": "string",
                    "maxLength": 250
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
        "models.Book": {
            "type": "object",
            "properties": {
                "createdAt": {
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
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.BooksQueryParams": {
            "type": "object",
            "properties": {
                "orderBy": {
                    "type": "string",
                    "default": "createdAt desc"
                },
                "page": {
                    "type": "integer"
                },
                "perPage": {
                    "type": "integer"
                }
            }
        },
        "models.FullProfileChartData": {
            "type": "object",
            "properties": {
                "cumulativeProgress": {
                    "description": "DailyProgress      []DailyProgressPoint ` + "`" + `json:\"daily_progress\"` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ProgressPoint"
                    }
                },
                "goalCompletion": {
                    "description": "MonthlyComparison  []MonthlyStats       ` + "`" + `json:\"monthly_comparison\"` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.GoalsChart"
                    }
                }
            }
        },
        "models.GoalStatistics": {
            "type": "object",
            "properties": {
                "averageWordsPerDay": {
                    "type": "number"
                },
                "bookId": {
                    "type": "string"
                },
                "dailyWordsRequired": {
                    "type": "number"
                },
                "daysElapsed": {
                    "type": "integer"
                },
                "daysRemaining": {
                    "type": "integer"
                },
                "goalId": {
                    "type": "string"
                },
                "percentageComplete": {
                    "type": "number"
                },
                "remainingWords": {
                    "type": "number"
                },
                "reportsCount": {
                    "type": "integer"
                },
                "totalWordsWritten": {
                    "type": "number"
                },
                "trendComparedToTarget": {
                    "type": "number"
                }
            }
        },
        "models.Goals": {
            "type": "object",
            "properties": {
                "bookId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "goalWords": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "isExpired": {
                    "type": "boolean"
                },
                "isFinished": {
                    "type": "boolean"
                },
                "startDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "wordsPerDay": {
                    "type": "number"
                },
                "writtenWords": {
                    "type": "number"
                }
            }
        },
        "models.GoalsChart": {
            "type": "object",
            "properties": {
                "averageWordsPerDay": {
                    "type": "number"
                },
                "bookId": {
                    "type": "string"
                },
                "bookTitle": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "dailyWordsRequired": {
                    "type": "number"
                },
                "daysElapsed": {
                    "type": "integer"
                },
                "daysRemaining": {
                    "type": "integer"
                },
                "goalId": {
                    "type": "string"
                },
                "goalTitle": {
                    "type": "string"
                },
                "isExpired": {
                    "type": "boolean"
                },
                "isFinished": {
                    "type": "boolean"
                },
                "percentageComplete": {
                    "type": "number"
                },
                "remainingWords": {
                    "type": "number"
                },
                "reportsCount": {
                    "type": "integer"
                },
                "totalWordsWritten": {
                    "type": "number"
                },
                "trendComparedToTarget": {
                    "type": "number"
                }
            }
        },
        "models.GoalsQueryParams": {
            "type": "object",
            "properties": {
                "bookId": {
                    "type": "string"
                },
                "orderBy": {
                    "type": "string",
                    "default": "createdAt desc"
                },
                "page": {
                    "type": "integer"
                },
                "perPage": {
                    "type": "integer"
                }
            }
        },
        "models.ProgressPoint": {
            "type": "object",
            "properties": {
                "bookId": {
                    "type": "string"
                },
                "bookName": {
                    "type": "string"
                },
                "completionPercent": {
                    "type": "number"
                },
                "date": {
                    "type": "string"
                },
                "goalId": {
                    "type": "string"
                },
                "goalTitle": {
                    "type": "string"
                },
                "targetTotalWords": {
                    "type": "number"
                },
                "totalWords": {
                    "type": "number"
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
        "models.Quotes": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "text": {
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
        "models.UserStatistics": {
            "type": "object",
            "properties": {
                "averageDaysToComplete": {
                    "type": "number"
                },
                "averageWordsPerDay": {
                    "type": "number"
                },
                "averageWordsPerReport": {
                    "type": "number"
                },
                "completedGoals": {
                    "type": "integer"
                },
                "currentStreak": {
                    "type": "integer"
                },
                "longestStreak": {
                    "type": "integer"
                },
                "maxWordsInDay": {
                    "type": "number"
                },
                "mostProductiveDay": {
                    "type": "string"
                },
                "totalBooks": {
                    "type": "integer"
                },
                "totalDaysWithActivity": {
                    "type": "integer"
                },
                "totalGoals": {
                    "type": "integer"
                },
                "totalReports": {
                    "type": "integer"
                },
                "totalWords": {
                    "type": "number"
                },
                "userId": {
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
	Host:             "http://127.0.0.1:5000",
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
