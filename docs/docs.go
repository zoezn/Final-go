// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://blog.counter-strike.net/index.php/server_guidelines/",
        "contact": {
            "name": "Soporte ZT",
            "url": "https://help.steampowered.com/es/wizard/HelpWithGame/?appid=730"
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
        "/dentistas/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Crear nuevo dentista.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/dentistas/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Buscar dentista por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Actualizar dentista por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Eliminar dentista por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Actualizar campos deseados de dentista por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/paciente/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Crear nuevo paciente.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/paciente/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Actualizar paciente por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/pacientes/:id": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Actualizar paciente por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Eliminar paciente por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Actualizar campos deseados de paciente por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/turno/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Crear nuevo turno.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/turno/:id": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Actualizar campos deseados de turno por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/turnos": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Obtener turno por dni de paciente.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "dni del paciente",
                        "name": "dni",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/turnos/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Buscar turno por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Actualizar turno por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Eliminar turno por Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/turnos/noids": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Crear turno por matricula dentista y id paciente.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "web.Response": {
            "type": "object",
            "properties": {
                "data": {}
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
	Title:            "Proyecto Final Esp. Backend IV",
	Description:      "Proyecto integrador final, API odontologica en GoLang.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
