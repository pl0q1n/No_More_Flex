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
    "title": "No more flex",
    "version": "0.0.1"
  },
  "paths": {
    "/transactions/add": {
      "post": {
        "summary": "Add new transaction",
        "operationId": "addTransaction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/transaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created"
          },
          "default": {
            "description": "generic server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/transactions/filter": {
      "get": {
        "summary": "Filter transactions",
        "operationId": "filterTransactions",
        "parameters": [
          {
            "type": "string",
            "name": "category",
            "in": "query"
          },
          {
            "type": "string",
            "name": "receiver",
            "in": "query"
          },
          {
            "type": "string",
            "name": "sender",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "from",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "to",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list of transactions",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/transaction"
              }
            }
          },
          "default": {
            "description": "generic server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "transaction": {
      "type": "object",
      "required": [
        "sender",
        "receiver",
        "value",
        "time"
      ],
      "properties": {
        "category": {
          "type": "string"
        },
        "receiver": {
          "type": "string"
        },
        "sender": {
          "type": "string"
        },
        "time": {
          "type": "integer",
          "format": "int64"
        },
        "value": {
          "type": "integer",
          "format": "int64"
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
    "title": "No more flex",
    "version": "0.0.1"
  },
  "paths": {
    "/transactions/add": {
      "post": {
        "summary": "Add new transaction",
        "operationId": "addTransaction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/transaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created"
          },
          "default": {
            "description": "generic server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/transactions/filter": {
      "get": {
        "summary": "Filter transactions",
        "operationId": "filterTransactions",
        "parameters": [
          {
            "type": "string",
            "name": "category",
            "in": "query"
          },
          {
            "type": "string",
            "name": "receiver",
            "in": "query"
          },
          {
            "type": "string",
            "name": "sender",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "from",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "to",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list of transactions",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/transaction"
              }
            }
          },
          "default": {
            "description": "generic server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "transaction": {
      "type": "object",
      "required": [
        "sender",
        "receiver",
        "value",
        "time"
      ],
      "properties": {
        "category": {
          "type": "string"
        },
        "receiver": {
          "type": "string"
        },
        "sender": {
          "type": "string"
        },
        "time": {
          "type": "integer",
          "format": "int64"
        },
        "value": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}`))
}
