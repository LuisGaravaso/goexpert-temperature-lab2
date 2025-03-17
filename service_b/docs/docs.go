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
        "/temperature/{location}": {
            "get": {
                "description": "Returns weather info by location (CEP or lat,lng)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Weather"
                ],
                "summary": "Get Weather by Location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Location (CEP or lat,lng)",
                        "name": "location",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/temperatures_internal_usecase_get_weather.GetWeatherOutputDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_infra_web.Error404"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/internal_infra_web.Error422"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_infra_web.Error500"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal_infra_web.Error404": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "location not found"
                }
            }
        },
        "internal_infra_web.Error422": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "invalid location"
                }
            }
        },
        "internal_infra_web.Error500": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        },
        "temperatures_internal_usecase_get_weather.GetWeatherOutputDTO": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "coordinates": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "humidity_in_percentage": {
                    "type": "integer"
                },
                "precipitation_in_millimeters": {
                    "type": "number"
                },
                "pressure_in_millibars": {
                    "type": "number"
                },
                "region": {
                    "type": "string"
                },
                "temp_C": {
                    "type": "number"
                },
                "temp_F": {
                    "type": "number"
                },
                "temp_K": {
                    "type": "number"
                },
                "wind_direction": {
                    "type": "string"
                },
                "wind_in_kph": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Weather API",
	Description:      "API for retrieving weather information by location",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
