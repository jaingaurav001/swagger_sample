// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "swaggertest",
    "title": "swaggertest",
    "version": "0.0.1"
  },
  "host": "localhost:8000",
  "basePath": "/",
  "paths": {
    "/search": {
      "get": {
        "tags": [
          "user"
        ],
        "summary": "get user information",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "user id",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "user_id"
      ],
      "properties": {
        "name": {
          "description": "name",
          "type": "string"
        },
        "user_id": {
          "description": "user id",
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "swaggertest",
    "title": "swaggertest",
    "version": "0.0.1"
  },
  "host": "localhost:8000",
  "basePath": "/",
  "paths": {
    "/search": {
      "get": {
        "tags": [
          "user"
        ],
        "summary": "get user information",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "user id",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "user_id"
      ],
      "properties": {
        "name": {
          "description": "name",
          "type": "string"
        },
        "user_id": {
          "description": "user id",
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
}
