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
    "description": "A School Management System",
    "title": "Apna School",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/student": {
      "post": {
        "operationId": "addStudent",
        "parameters": [
          {
            "description": "student's details",
            "name": "student",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "student added",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "409": {
            "description": "student already exists"
          }
        }
      }
    },
    "/student/signup": {
      "post": {
        "operationId": "SignUpStudent",
        "parameters": [
          {
            "description": "student's details",
            "name": "student",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "student added",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "409": {
            "description": "student already exists"
          }
        }
      }
    },
    "/student/{ID}": {
      "get": {
        "description": "return student based on UUID",
        "operationId": "getStudent",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the student",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "student response",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "404": {
            "description": "student not found"
          }
        }
      },
      "put": {
        "operationId": "editStudent",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the student",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "student updated",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "operationId": "deleteStudent",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the student",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "student deleted"
          },
          "404": {
            "description": "student not found"
          }
        }
      }
    },
    "/teacher": {
      "post": {
        "operationId": "addTeacher",
        "parameters": [
          {
            "description": "teacher's details",
            "name": "teacher",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "teacher added",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "409": {
            "description": "teacher already exists"
          }
        }
      }
    },
    "/teacher/signup": {
      "post": {
        "operationId": "SignUpTeacher",
        "parameters": [
          {
            "description": "teacher's details",
            "name": "teacher",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "teacher added",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "409": {
            "description": "teacher already exists"
          }
        }
      }
    },
    "/teacher/{ID}": {
      "get": {
        "description": "return teacher based on UUID",
        "operationId": "getTeacher",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the teacher",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "teacher response",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "404": {
            "description": "teacher not found"
          }
        }
      },
      "put": {
        "operationId": "editTeacher",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the teacher",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "teacher updated",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "operationId": "deleteTeacher",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the teacher",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "teacher deleted"
          },
          "404": {
            "description": "teacher not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Student": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "string"
        },
        "Department": {
          "type": "string"
        },
        "Grade": {
          "type": "integer"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "Teacher": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "string"
        },
        "Department": {
          "type": "string"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        },
        "Salary": {
          "type": "integer"
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
    "description": "A School Management System",
    "title": "Apna School",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/student": {
      "post": {
        "operationId": "addStudent",
        "parameters": [
          {
            "description": "student's details",
            "name": "student",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "student added",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "409": {
            "description": "student already exists"
          }
        }
      }
    },
    "/student/signup": {
      "post": {
        "operationId": "SignUpStudent",
        "parameters": [
          {
            "description": "student's details",
            "name": "student",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "student added",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "409": {
            "description": "student already exists"
          }
        }
      }
    },
    "/student/{ID}": {
      "get": {
        "description": "return student based on UUID",
        "operationId": "getStudent",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the student",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "student response",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "404": {
            "description": "student not found"
          }
        }
      },
      "put": {
        "operationId": "editStudent",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the student",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "student updated",
            "schema": {
              "$ref": "#/definitions/Student"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "operationId": "deleteStudent",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the student",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "student deleted"
          },
          "404": {
            "description": "student not found"
          }
        }
      }
    },
    "/teacher": {
      "post": {
        "operationId": "addTeacher",
        "parameters": [
          {
            "description": "teacher's details",
            "name": "teacher",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "teacher added",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "409": {
            "description": "teacher already exists"
          }
        }
      }
    },
    "/teacher/signup": {
      "post": {
        "operationId": "SignUpTeacher",
        "parameters": [
          {
            "description": "teacher's details",
            "name": "teacher",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "teacher added",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "409": {
            "description": "teacher already exists"
          }
        }
      }
    },
    "/teacher/{ID}": {
      "get": {
        "description": "return teacher based on UUID",
        "operationId": "getTeacher",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the teacher",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "teacher response",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "404": {
            "description": "teacher not found"
          }
        }
      },
      "put": {
        "operationId": "editTeacher",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the teacher",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "teacher updated",
            "schema": {
              "$ref": "#/definitions/Teacher"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "operationId": "deleteTeacher",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the teacher",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "teacher deleted"
          },
          "404": {
            "description": "teacher not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Student": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "string"
        },
        "Department": {
          "type": "string"
        },
        "Grade": {
          "type": "integer"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "Teacher": {
      "type": "object",
      "properties": {
        "Age": {
          "type": "string"
        },
        "Department": {
          "type": "string"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        },
        "Salary": {
          "type": "integer"
        }
      }
    }
  }
}`))
}
