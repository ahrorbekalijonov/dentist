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
        "/v1/appointment": {
            "get": {
                "description": "Api for get a new appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "GetAppointment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Api for update appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "UpdateAppointment",
                "parameters": [
                    {
                        "description": "Appointment",
                        "name": "Appointment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Api for creat a new appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "CreateAppointment",
                "parameters": [
                    {
                        "description": "CreateAppointment",
                        "name": "Appointment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAppointment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Api for delete appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "DeleteAppointment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/appointmentid": {
            "get": {
                "description": "Api for get appointment with client id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "GetAppointmentWithClientId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client_id",
                        "name": "client_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AllAppointments"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/appointmentnew": {
            "post": {
                "description": "Api for creating a new client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "CreateAppointmentWithClient",
                "parameters": [
                    {
                        "description": "CreateAppointmentWithClient",
                        "name": "Client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.New"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/appointments": {
            "get": {
                "description": "Api for get all appointments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "GetAllAppointments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AllAppointments"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/appointmentsdate": {
            "get": {
                "description": "Api for get appointments with date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appointment"
                ],
                "summary": "GetAppointmentsWithDate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "integer",
                        "name": "integer",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AllAppointments"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/client": {
            "get": {
                "description": "Api for get client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "GetClient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Api for update client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "UpdateClient",
                "parameters": [
                    {
                        "description": "UpdateClient",
                        "name": "Client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Api for creating a new client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "CreateClient",
                "parameters": [
                    {
                        "description": "CreateClient",
                        "name": "Client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqClient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Api for delete client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "DeleteClient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/clients": {
            "get": {
                "description": "Api for get all clients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "GetAllClients",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AllClients"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/count": {
            "get": {
                "description": "Api for get all clients count",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "GetAllClientsCount",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/search": {
            "get": {
                "description": "Api for searching clients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "SearchingClients",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SearchClients",
                        "name": "str",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AllClients"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AllAppointments": {
            "type": "object",
            "properties": {
                "appointments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Appointment"
                    }
                }
            }
        },
        "models.AllClients": {
            "type": "object",
            "properties": {
                "clients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Client"
                    }
                }
            }
        },
        "models.Appointment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "clientId": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "diagnostics": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "treatment": {
                    "type": "string"
                }
            }
        },
        "models.Client": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birthDate": {
                    "type": "string"
                },
                "fatherName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.StandartError"
                }
            }
        },
        "models.New": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "appointmentId": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "clientName": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "diagnostics": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "treatment": {
                    "type": "string"
                }
            }
        },
        "models.ReqAppointment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "clientId": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "diagnostics": {
                    "type": "string"
                },
                "treatment": {
                    "type": "string"
                }
            }
        },
        "models.ReqClient": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birthDate": {
                    "type": "string"
                },
                "fatherName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "models.ReqNew": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "clientName": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "diagnostics": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "treatment": {
                    "type": "string"
                }
            }
        },
        "models.StandartError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Dentist",
	Description:      "Dentist-backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}