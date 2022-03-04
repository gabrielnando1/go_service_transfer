// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/transfer/receive": {
            "post": {
                "description": "endpoint for  transfer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfer"
                ],
                "summary": "transfer",
                "parameters": [
                    {
                        "description": "transfer request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto_transfer_request.TransferReceiveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto_common_response.GenericResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "$ref": "#/definitions/dto_transfer_response.TransferReceiveResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto_common_response.GenericResponse": {
            "type": "object",
            "properties": {
                "content": {},
                "error_message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "dto_transfer_request.TransferReceiveRequest": {
            "type": "object",
            "properties": {
                "account_from": {
                    "type": "string"
                },
                "account_target": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "expected_on": {
                    "type": "string"
                },
                "expiration_date": {
                    "type": "string"
                },
                "external_id": {
                    "type": "string"
                }
            }
        },
        "dto_transfer_response.TransferReceiveResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "internal_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}