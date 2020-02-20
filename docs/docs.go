// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-20 11:38:36.824433104 +0100 CET m=+0.044809706

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/node/add": {
            "post": {
                "description": "This will verify if a node meets the requirements and will join to the cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new bootstrapped, ready node to the cluster.",
                "operationId": "add-node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "FQDN of the cluster node for adding it",
                        "name": "fqdn",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/node/list": {
            "get": {
                "description": "List all nodes in the current cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List nodes in the cluster.",
                "operationId": "list-nodes",
                "responses": {}
            }
        },
        "/api/v1/node/stage": {
            "post": {
                "description": "This will install a client binary over SSH and will run nanostate, required to setup everything in place",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Stage (bootstrap) a new cluster node.",
                "operationId": "stage-node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "FQDN of the hostname for staging",
                        "name": "fqdn",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Root password of the node",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "State Id for bootstrapping",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/zones/add": {
            "post": {
                "description": "AddZone creates a new empty zone in the cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Define a cluster Zone.",
                "operationId": "add-zone",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the Zone",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Zone description",
                        "name": "description",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/zones/list": {
            "get": {
                "description": "List all zones in the Cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List cluster zones",
                "operationId": "list-zones",
                "responses": {}
            }
        },
        "/api/v1/zones/remove": {
            "delete": {
                "description": "RemoveZone removes a zone from the cluster, but only if it is empty (no nodes assigned to it).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove an empty cluster Zone",
                "operationId": "remove-zone",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the Zone",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/zones/stats": {
            "get": {
                "description": "ZoneStats returns data about zone.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return Zone stats.",
                "operationId": "zone-stats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the Zone",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/zones/update": {
            "post": {
                "description": "UpdateZone updates a zone data,",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a cluster Zone",
                "operationId": "update-zone",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the Zone",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Zone description",
                        "name": "description",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
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
	swag.Register(swag.Name, &s{})
}
