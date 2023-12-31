// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "Leung Yan Tung",
            "url": "https://github.com/ushio0107",
            "email": "leungyantung0107@gmail.com"
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
        "/v1/login": {
            "post": {
                "description": "Verifies the provided account credentials.\nIf the password verification fails five times, the user is required to wait for one minute before attempting again.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Verify an account",
                "parameters": [
                    {
                        "description": "Account credentials",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseType.Response"
                        }
                    },
                    "401": {
                        "description": "Incorrect username or password",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseType.BadRequestResponse"
                        }
                    },
                    "429": {
                        "description": "Password attempts exceed ",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseType.BadRequestResponse"
                        }
                    }
                }
            }
        },
        "/v1/signup": {
            "post": {
                "description": "Create an account by the desired username and password.\nEnter the username and password,\n\nThe username must meet the following criteria:\n- Minimum length of 3 characters and a maximum length of 32 characters.\n\nThe password must meet the following criteria:\n- Minimum length of 8 characters and maximum length of 32 characters.\n- Must contain at least 1 uppercase letter, 1 lowercase letter, and 1 number.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Create an account",
                "parameters": [
                    {
                        "description": "Account credentials",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseType.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid username or password",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseType.BadRequestResponse"
                        }
                    },
                    "409": {
                        "description": "Account already exists",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseType.BadRequestResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AccountRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ResponseType.BadRequestResponse": {
            "type": "object",
            "properties": {
                "reason": {
                    "type": "string",
                    "example": "failed reason"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "models.ResponseType.Response": {
            "type": "object",
            "properties": {
                "reason": {
                    "type": "string",
                    "example": ""
                },
                "success": {
                    "type": "boolean",
                    "example": true
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
	Title:            "account_management_api",
	Description:      "This is a REST API which can create an account and verify an account.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
