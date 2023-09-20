// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "arsenzhar@yandex.ru."
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "post": {
                "description": "Получает на вход ссылку и отдаёт в ответе сокращённый вариант",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "Сокращает ссылку",
                "parameters": [
                    {
                        "example": "https://ya.ru",
                        "description": "URL для сокращения",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Сокращённая ссылка",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/shorten": {
            "post": {
                "description": "Получает на вход ссылку и отдаёт в ответе сокращённый вариант",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Сокращает ссылку",
                "parameters": [
                    {
                        "description": "URL для сокращения",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestCreateLink"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCreateLink"
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/shorten/batch": {
            "post": {
                "description": "Получает на вход массив ссылок и отдаёт в ответе сокращённый вариант",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Сокращает массив ссылок",
                "parameters": [
                    {
                        "description": "Массив URL для сокращения",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RequestLinks"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ResponseLinks"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/user/urls": {
            "get": {
                "description": "Отдаёт все ссылки пользователя",
                "produces": [
                    "application/json"
                ],
                "summary": "URL пользователя",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "204": {
                        "description": "У пользователя нет ссылок",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Получает на вход массив ссылок и удаляет их",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "Удаляет ссылки пользователя",
                "parameters": [
                    {
                        "description": "URL для удаления",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Отвечает OK, если работает",
                "produces": [
                    "text/plain"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/{url_id}": {
            "get": {
                "description": "По сокращённой ссылке делает редирект на оригинальную ссылку",
                "summary": "Делает редирект на оригинальную ссылку",
                "parameters": [
                    {
                        "type": "string",
                        "example": "maIJa1",
                        "description": "ID Short URL",
                        "name": "url_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "307": {
                        "description": "Temporary Redirect",
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "http://localhost:8080/maIJa1"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.RequestCreateLink": {
            "type": "object",
            "properties": {
                "url": {
                    "description": "URL – ссылка, которую необходимо сократить.",
                    "type": "string"
                }
            }
        },
        "models.RequestLinks": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "description": "CorrelationID – ID ссылки.",
                    "type": "string",
                    "format": "uuid",
                    "example": "58039b0a-480d-11ee-9ace-0e6250a0eb02"
                },
                "original_url": {
                    "description": "OriginalURL – ссылка, которую необходимо сократить.",
                    "type": "string",
                    "example": "https://ya.ru"
                }
            }
        },
        "models.ResponseCreateLink": {
            "type": "object",
            "properties": {
                "result": {
                    "description": "Result – сокращённая ссылка.",
                    "type": "string",
                    "example": "http://localhost:8080/maIJa1"
                }
            }
        },
        "models.ResponseLinks": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "description": "CorrelationID – ID ссылки.",
                    "type": "string",
                    "format": "uuid",
                    "example": "58039b0a-480d-11ee-9ace-0e6250a0eb02"
                },
                "short_url": {
                    "description": "ShortURL – сокращённая ссылка.",
                    "type": "string",
                    "example": "http://localhost:8080/maIJa1"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.",
	Host:             "localhost:8080.",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "URLShortener API",
	Description:      "Сервис сокращения URL.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}