// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/account/list": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "get account list. Only administrator can get list of accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "List Accounts",
                "parameters": [
                    {
                        "description": "account list request filter",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.FindAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Account"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/account/{acc_id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "get account by id. Creator can get only his account. Administrator can get any account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "acc_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Account"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "deletes an account. Only administrator can delete accounts.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Delete Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "acc_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/account/{acc_id}/email": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "updates email for current user. Admin can update email for any user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Update Email",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "acc_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "account new email",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateAccountEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/account/{acc_id}/password": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "updates password for current account. Admin can change password without provieding an old password. Admin can update password for any user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Update Password",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "acc_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "account old and new password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateAccountPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/account/{acc_id}/user": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "updates telegram user conected to this account. Admin, can update email for any user. Either of one of the fields must be provided",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Update TgUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "acc_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "account new email",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateAccountUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "logins in to account with user provided credantials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "login credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.loginAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.loginAccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "registers new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "register credantials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.registerAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/sub": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "creates subscribtion with given users telegram_id and account_id. Only administrator and bot can create subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "subscription info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createSubscriptionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "updates subscription. Admin and bot can update subscription. Can be used to change subscription status, or price.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "subscription new status and price",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateSubscriptionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/sub/list": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "get subscription list. Only administrator and bot can get list of any accounts. Ordenery user can get list of subscriptions whitch belongs to his account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "List Subscriptions",
                "parameters": [
                    {
                        "description": "subscription list request filter",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.FindSubRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Sub"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/user": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "creates user with given users telegram_id and username. Only administrator and bot can create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.createUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "get user list. Only administrator and bot can get list of accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "List Users",
                "parameters": [
                    {
                        "description": "user list request filter",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.FindUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Tguser"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/user/{usr_id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "returns user object by given id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "usr_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.getUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "deletes user object by given id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "usr_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/user/{usr_id}/password": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "updateds users username and/or status, provided by id. Admin and bot can update any user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "usr_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user new status or username",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.updateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Account": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "chan_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "role": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.FindAccountRequest": {
            "type": "object",
            "properties": {
                "page_settings": {
                    "description": "Username *struct {\n\tLike string ` + "`" + `json:\"LIKE\"` + "`" + `\n\tEq   string ` + "`" + `json:\"EQ\"` + "`" + `\n} ` + "`" + `json:\"username\"` + "`" + `\nEmail *struct {\n\tLike string ` + "`" + `json:\"LIKE\"` + "`" + `\n\tEq   string ` + "`" + `json:\"EQ\"` + "`" + `\n} ` + "`" + `json:\"email\"` + "`" + `",
                    "type": "object",
                    "properties": {
                        "page_number": {
                            "type": "integer"
                        },
                        "page_size": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "domain.FindSubRequest": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "object",
                    "properties": {
                        "eq": {
                            "type": "integer"
                        }
                    }
                },
                "page_settings": {
                    "type": "object",
                    "properties": {
                        "page_number": {
                            "type": "integer"
                        },
                        "page_size": {
                            "type": "integer"
                        }
                    }
                },
                "price": {
                    "type": "object",
                    "properties": {
                        "eq": {
                            "type": "integer"
                        },
                        "range": {
                            "type": "object",
                            "properties": {
                                "from": {
                                    "type": "integer"
                                },
                                "to": {
                                    "type": "integer"
                                }
                            }
                        }
                    }
                },
                "status": {
                    "type": "object",
                    "properties": {
                        "eq": {
                            "type": "string"
                        }
                    }
                },
                "tguser_id": {
                    "type": "object",
                    "properties": {
                        "eq": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "domain.FindUserRequest": {
            "type": "object",
            "properties": {
                "page_settings": {
                    "description": "Username *struct {\n\tLike string ` + "`" + `json:\"LIKE\"` + "`" + `\n\tEq   string ` + "`" + `json:\"EQ\"` + "`" + `\n} ` + "`" + `json:\"username\"` + "`" + `\nEmail *struct {\n\tLike string ` + "`" + `json:\"LIKE\"` + "`" + `\n\tEq   string ` + "`" + `json:\"EQ\"` + "`" + `\n} ` + "`" + `json:\"email\"` + "`" + `",
                    "type": "object",
                    "properties": {
                        "page_number": {
                            "type": "integer"
                        },
                        "page_size": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "domain.Sub": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "activated_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.Tguser": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "telegram_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "http.accountResponse": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "http.createSubscriptionRequest": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "expires_at": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "http.createUserRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "telegram_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "http.createUserResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "telegram_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "http.getUserResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "telegram_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "http.loginAccountRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "http.loginAccountResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_expires_at": {
                    "type": "string"
                },
                "account": {
                    "$ref": "#/definitions/http.accountResponse"
                },
                "refresh_token": {
                    "type": "string"
                },
                "refresh_token_expires_at": {
                    "type": "string"
                }
            }
        },
        "http.registerAccountRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "http.updateAccountEmailRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "http.updateAccountPasswordRequest": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "http.updateAccountUserRequest": {
            "type": "object",
            "properties": {
                "telegram_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "http.updateSubscriptionRequest": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "http.updateUserRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "http.updateUserResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "telegram_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Telegram Subs API",
	Description:      "This service provide functionality for storing and managing privat telegram channels with subscription based payments for acessing content.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}