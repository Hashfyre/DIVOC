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
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Digital infra for vaccination certificates",
    "title": "Divoc Portal API",
    "version": "1.0.0"
  },
  "host": "divoc.xiv.in",
  "basePath": "/divoc/admin/api/v1",
  "paths": {
    "/enrollments": {
      "get": {
        "summary": "get enrollments",
        "operationId": "getEnrollments",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload facility csv for bulk ingestion of pre enrollment",
        "parameters": [
          {
            "type": "file",
            "description": "Facility data in the form of csv",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facilities": {
      "get": {
        "summary": "get facilities",
        "operationId": "getFacilities",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Facility"
              }
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload facility csv for bulk ingestion",
        "parameters": [
          {
            "type": "file",
            "description": "Facility data in the form of csv",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/groups": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facillity-admin"
            ]
          }
        ],
        "summary": "Get facility groups",
        "operationId": "getFacilityGroups",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "#/definitions/UserGroup"
              }
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/users": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facillity-admin"
            ]
          }
        ],
        "summary": "Get users of a facility",
        "operationId": "getFacilityUsers",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "#/definitions/FacilityUser"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "hasRole": [
              "facillity-admin"
            ]
          }
        ],
        "summary": "Create Facility User",
        "operationId": "createFacilityUsers",
        "parameters": [
          {
            "description": "Create facility user data",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/FacilityUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/medicines": {
      "get": {
        "security": [
          {
            "hasRole": [
              "user",
              "admin"
            ]
          }
        ],
        "summary": "Get medicines",
        "operationId": "getMedicines",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "#/definitions/CreateMedicineRequest"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "hasRole": [
              "admin"
            ]
          }
        ],
        "summary": "Create Medicine",
        "operationId": "createMedicine",
        "parameters": [
          {
            "description": "Facility data in the form of csv",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/CreateMedicineRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/programs": {
      "get": {
        "security": [
          {
            "hasRole": [
              "admin",
              "user"
            ]
          }
        ],
        "summary": "get program list",
        "operationId": "getPrograms",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Program"
              }
            }
          }
        }
      },
      "post": {
        "summary": "Create program",
        "operationId": "createProgram",
        "parameters": [
          {
            "description": "Vaccination Program",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ProgramRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/vaccinators": {
      "get": {
        "summary": "Get vaccinators",
        "operationId": "getVaccinators",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Vaccinator"
              }
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload facility csv for bulk ingestion",
        "parameters": [
          {
            "type": "file",
            "description": "Facility data in the form of csv",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    }
  },
  "definitions": {
    "Address": {
      "description": "Indian address format",
      "type": "object",
      "title": "Address",
      "required": [
        "addressLine1",
        "addressLine2",
        "district",
        "state",
        "pincode"
      ],
      "properties": {
        "addressLine1": {
          "type": "string"
        },
        "addressLine2": {
          "type": "string"
        },
        "district": {
          "type": "string"
        },
        "pincode": {
          "type": "integer"
        },
        "state": {
          "description": "State of address",
          "type": "string",
          "title": "The state schema",
          "example": [
            "Karnataka"
          ]
        }
      },
      "example": [
        {
          "addressLine1": "no. 23, some lane, some road",
          "addressLine2": "some nagar",
          "district": "bangalore south",
          "pincode": 560000,
          "state": "Karnataka"
        }
      ]
    },
    "CreateMedicineRequest": {
      "type": "object",
      "properties": {
        "effectiveUntil": {
          "description": "Effective until n months after the full vaccination schedule is completed",
          "type": "number"
        },
        "name": {
          "type": "string"
        },
        "price": {
          "description": "Indicative price if fixed or max price available.",
          "type": "number"
        },
        "provider": {
          "type": "string"
        },
        "schedule": {
          "type": "object",
          "properties": {
            "repeatInterval": {
              "description": "Number of times the vaccination should be taken.",
              "type": "number"
            },
            "repeatTimes": {
              "description": "How many times vaccination should be taken",
              "type": "number"
            }
          }
        },
        "status": {
          "type": "string",
          "enum": [
            "Active",
            "Inactive",
            "Blocked"
          ]
        },
        "vaccinationMode": {
          "type": "string",
          "enum": [
            "muscular injection",
            "oral",
            "nasal"
          ]
        }
      }
    },
    "Facility": {
      "properties": {
        "address": {
          "title": "Address",
          "$ref": "#/definitions/Address"
        },
        "admins": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "averageRating": {
          "description": "Average Rating of Facility 0 to 5, 0 for no rating.",
          "type": "number",
          "title": "Average Rating",
          "default": 0
        },
        "category": {
          "type": "string",
          "title": "Category",
          "enum": [
            "GOVT",
            "PRIVATE"
          ]
        },
        "contact": {
          "type": "string",
          "title": "Contact number"
        },
        "facilityCode": {
          "type": "string",
          "title": "Facility Code"
        },
        "facilityName": {
          "type": "string",
          "title": "Facility Name"
        },
        "geoLocation": {
          "type": "string",
          "title": "Geo Location"
        },
        "operatingHourEnd": {
          "type": "integer",
          "title": "Operating hours end of day"
        },
        "operatingHourStart": {
          "type": "integer",
          "title": "Operating hours start of day"
        },
        "serialNum": {
          "type": "integer",
          "title": "Serial Number"
        },
        "stamp": {
          "type": "string"
        },
        "status": {
          "type": "string",
          "title": "Status of Facility",
          "enum": [
            "Active",
            "Inactive",
            "Blocked"
          ]
        },
        "type": {
          "type": "string",
          "title": "Type of Facility",
          "enum": [
            "Fixed location",
            "Mobile",
            "Both"
          ]
        },
        "websiteUrl": {
          "type": "string",
          "title": "Website URL"
        }
      }
    },
    "FacilityUser": {
      "properties": {
        "employeeId": {
          "type": "string",
          "title": "Facility User Id"
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/UserGroup"
          }
        },
        "id": {
          "type": "string",
          "title": "Facility User Id"
        },
        "mobileNumber": {
          "type": "string",
          "title": "Facility User Mobile Number"
        },
        "name": {
          "type": "string",
          "title": "Facility User Name"
        }
      }
    },
    "Program": {
      "type": "object",
      "title": "Program",
      "required": [
        "name",
        "description",
        "startDate"
      ],
      "properties": {
        "description": {
          "type": "string",
          "title": "Description"
        },
        "endDate": {
          "type": "string",
          "format": "date",
          "title": "End Date"
        },
        "logoURL": {
          "type": "string",
          "title": "Logo URL"
        },
        "medicineIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string",
          "title": "Name"
        },
        "startDate": {
          "type": "string",
          "format": "date",
          "title": "Start Date"
        },
        "status": {
          "type": "string",
          "title": "Status",
          "enum": [
            "Active",
            "Inactive"
          ]
        }
      }
    },
    "ProgramRequest": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "endDate": {
          "type": "string",
          "format": "date"
        },
        "logoURL": {
          "type": "string"
        },
        "medicineIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        },
        "startDate": {
          "type": "string",
          "format": "date"
        },
        "status": {
          "type": "string",
          "enum": [
            "Active",
            "Inactive"
          ]
        }
      }
    },
    "UserGroup": {
      "properties": {
        "id": {
          "type": "string",
          "title": "group id"
        },
        "name": {
          "type": "string",
          "title": "group name"
        }
      }
    },
    "Vaccinator": {
      "type": "object",
      "title": "The Vaccinator Schema",
      "required": [
        "serialNum",
        "code",
        "nationalIdentifier",
        "name",
        "facilityIds",
        "mobileNumber",
        "averageRating",
        "trainingCertificate",
        "status"
      ],
      "properties": {
        "averageRating": {
          "type": "number"
        },
        "code": {
          "type": "string"
        },
        "facilityIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "mobileNumber": {
          "type": "string",
          "maxLength": 10,
          "minLength": 10
        },
        "name": {
          "type": "string",
          "title": "Full name"
        },
        "nationalIdentifier": {
          "type": "string"
        },
        "serialNum": {
          "type": "integer"
        },
        "signatures": {
          "type": "array",
          "items": {
            "$ref": "Signature.json#/definitions/Signature"
          }
        },
        "status": {
          "type": "string",
          "enum": [
            "Active",
            "Inactive"
          ]
        },
        "trainingCertificate": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://dummy.oauth.net/auth",
      "tokenUrl": "https://dumy.oauth.net/token",
      "scopes": {
        "admin": "scope of super admin",
        "facillity-admin": "scope of facility admin"
      }
    }
  },
  "security": [
    {
      "hasRole": [
        "admin"
      ]
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Digital infra for vaccination certificates",
    "title": "Divoc Portal API",
    "version": "1.0.0"
  },
  "host": "divoc.xiv.in",
  "basePath": "/divoc/admin/api/v1",
  "paths": {
    "/enrollments": {
      "get": {
        "summary": "get enrollments",
        "operationId": "getEnrollments",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload facility csv for bulk ingestion of pre enrollment",
        "parameters": [
          {
            "type": "file",
            "description": "Facility data in the form of csv",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facilities": {
      "get": {
        "summary": "get facilities",
        "operationId": "getFacilities",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Facility"
              }
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload facility csv for bulk ingestion",
        "parameters": [
          {
            "type": "file",
            "description": "Facility data in the form of csv",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/groups": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facillity-admin"
            ]
          }
        ],
        "summary": "Get facility groups",
        "operationId": "getFacilityGroups",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "#/definitions/UserGroup"
              }
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/users": {
      "get": {
        "security": [
          {
            "hasRole": [
              "facillity-admin"
            ]
          }
        ],
        "summary": "Get users of a facility",
        "operationId": "getFacilityUsers",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "#/definitions/FacilityUser"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "hasRole": [
              "facillity-admin"
            ]
          }
        ],
        "summary": "Create Facility User",
        "operationId": "createFacilityUsers",
        "parameters": [
          {
            "description": "Create facility user data",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/FacilityUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/medicines": {
      "get": {
        "security": [
          {
            "hasRole": [
              "admin",
              "user"
            ]
          }
        ],
        "summary": "Get medicines",
        "operationId": "getMedicines",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "#/definitions/CreateMedicineRequest"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "hasRole": [
              "admin"
            ]
          }
        ],
        "summary": "Create Medicine",
        "operationId": "createMedicine",
        "parameters": [
          {
            "description": "Facility data in the form of csv",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/CreateMedicineRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/programs": {
      "get": {
        "security": [
          {
            "hasRole": [
              "admin",
              "user"
            ]
          }
        ],
        "summary": "get program list",
        "operationId": "getPrograms",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Program"
              }
            }
          }
        }
      },
      "post": {
        "summary": "Create program",
        "operationId": "createProgram",
        "parameters": [
          {
            "description": "Vaccination Program",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ProgramRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/vaccinators": {
      "get": {
        "summary": "Get vaccinators",
        "operationId": "getVaccinators",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Vaccinator"
              }
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload facility csv for bulk ingestion",
        "parameters": [
          {
            "type": "file",
            "description": "Facility data in the form of csv",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid input"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    }
  },
  "definitions": {
    "Address": {
      "description": "Indian address format",
      "type": "object",
      "title": "Address",
      "required": [
        "addressLine1",
        "addressLine2",
        "district",
        "state",
        "pincode"
      ],
      "properties": {
        "addressLine1": {
          "type": "string"
        },
        "addressLine2": {
          "type": "string"
        },
        "district": {
          "type": "string"
        },
        "pincode": {
          "type": "integer"
        },
        "state": {
          "description": "State of address",
          "type": "string",
          "title": "The state schema",
          "example": [
            "Karnataka"
          ]
        }
      },
      "example": [
        {
          "addressLine1": "no. 23, some lane, some road",
          "addressLine2": "some nagar",
          "district": "bangalore south",
          "pincode": 560000,
          "state": "Karnataka"
        }
      ]
    },
    "CreateMedicineRequest": {
      "type": "object",
      "properties": {
        "effectiveUntil": {
          "description": "Effective until n months after the full vaccination schedule is completed",
          "type": "number"
        },
        "name": {
          "type": "string"
        },
        "price": {
          "description": "Indicative price if fixed or max price available.",
          "type": "number"
        },
        "provider": {
          "type": "string"
        },
        "schedule": {
          "type": "object",
          "properties": {
            "repeatInterval": {
              "description": "Number of times the vaccination should be taken.",
              "type": "number"
            },
            "repeatTimes": {
              "description": "How many times vaccination should be taken",
              "type": "number"
            }
          }
        },
        "status": {
          "type": "string",
          "enum": [
            "Active",
            "Inactive",
            "Blocked"
          ]
        },
        "vaccinationMode": {
          "type": "string",
          "enum": [
            "muscular injection",
            "oral",
            "nasal"
          ]
        }
      }
    },
    "CreateMedicineRequestSchedule": {
      "type": "object",
      "properties": {
        "repeatInterval": {
          "description": "Number of times the vaccination should be taken.",
          "type": "number"
        },
        "repeatTimes": {
          "description": "How many times vaccination should be taken",
          "type": "number"
        }
      }
    },
    "Facility": {
      "properties": {
        "address": {
          "title": "Address",
          "$ref": "#/definitions/Address"
        },
        "admins": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "averageRating": {
          "description": "Average Rating of Facility 0 to 5, 0 for no rating.",
          "type": "number",
          "title": "Average Rating",
          "default": 0
        },
        "category": {
          "type": "string",
          "title": "Category",
          "enum": [
            "GOVT",
            "PRIVATE"
          ]
        },
        "contact": {
          "type": "string",
          "title": "Contact number"
        },
        "facilityCode": {
          "type": "string",
          "title": "Facility Code"
        },
        "facilityName": {
          "type": "string",
          "title": "Facility Name"
        },
        "geoLocation": {
          "type": "string",
          "title": "Geo Location"
        },
        "operatingHourEnd": {
          "type": "integer",
          "title": "Operating hours end of day"
        },
        "operatingHourStart": {
          "type": "integer",
          "title": "Operating hours start of day"
        },
        "serialNum": {
          "type": "integer",
          "title": "Serial Number"
        },
        "stamp": {
          "type": "string"
        },
        "status": {
          "type": "string",
          "title": "Status of Facility",
          "enum": [
            "Active",
            "Inactive",
            "Blocked"
          ]
        },
        "type": {
          "type": "string",
          "title": "Type of Facility",
          "enum": [
            "Fixed location",
            "Mobile",
            "Both"
          ]
        },
        "websiteUrl": {
          "type": "string",
          "title": "Website URL"
        }
      }
    },
    "FacilityUser": {
      "properties": {
        "employeeId": {
          "type": "string",
          "title": "Facility User Id"
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/UserGroup"
          }
        },
        "id": {
          "type": "string",
          "title": "Facility User Id"
        },
        "mobileNumber": {
          "type": "string",
          "title": "Facility User Mobile Number"
        },
        "name": {
          "type": "string",
          "title": "Facility User Name"
        }
      }
    },
    "Program": {
      "type": "object",
      "title": "Program",
      "required": [
        "name",
        "description",
        "startDate"
      ],
      "properties": {
        "description": {
          "type": "string",
          "title": "Description"
        },
        "endDate": {
          "type": "string",
          "format": "date",
          "title": "End Date"
        },
        "logoURL": {
          "type": "string",
          "title": "Logo URL"
        },
        "medicineIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string",
          "title": "Name"
        },
        "startDate": {
          "type": "string",
          "format": "date",
          "title": "Start Date"
        },
        "status": {
          "type": "string",
          "title": "Status",
          "enum": [
            "Active",
            "Inactive"
          ]
        }
      }
    },
    "ProgramRequest": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "endDate": {
          "type": "string",
          "format": "date"
        },
        "logoURL": {
          "type": "string"
        },
        "medicineIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        },
        "startDate": {
          "type": "string",
          "format": "date"
        },
        "status": {
          "type": "string",
          "enum": [
            "Active",
            "Inactive"
          ]
        }
      }
    },
    "UserGroup": {
      "properties": {
        "id": {
          "type": "string",
          "title": "group id"
        },
        "name": {
          "type": "string",
          "title": "group name"
        }
      }
    },
    "Vaccinator": {
      "type": "object",
      "title": "The Vaccinator Schema",
      "required": [
        "serialNum",
        "code",
        "nationalIdentifier",
        "name",
        "facilityIds",
        "mobileNumber",
        "averageRating",
        "trainingCertificate",
        "status"
      ],
      "properties": {
        "averageRating": {
          "type": "number"
        },
        "code": {
          "type": "string"
        },
        "facilityIds": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "mobileNumber": {
          "type": "string",
          "maxLength": 10,
          "minLength": 10
        },
        "name": {
          "type": "string",
          "title": "Full name"
        },
        "nationalIdentifier": {
          "type": "string"
        },
        "serialNum": {
          "type": "integer"
        },
        "signatures": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/signature"
          }
        },
        "status": {
          "type": "string",
          "enum": [
            "Active",
            "Inactive"
          ]
        },
        "trainingCertificate": {
          "type": "string"
        }
      }
    },
    "signature": {
      "type": "object",
      "title": "The Signature Schema for the registry",
      "required": [
        "@type",
        "signatureFor",
        "creator",
        "created",
        "signatureValue"
      ],
      "properties": {
        "@type": {
          "type": "string",
          "default": "sc:RsaSignature2018",
          "enum": [
            "sc:LinkedDataSignature2015",
            "sc:GraphSignature2012",
            "sc:RsaSignature2018"
          ],
          "$id": "#/properties/@type"
        },
        "created": {
          "type": "string",
          "format": "date-time",
          "$comment": "Timestamp",
          "$id": "#/properties/created",
          "examples": [
            "2017-09-23T20:21:34Z"
          ]
        },
        "creator": {
          "type": "string",
          "format": "uri",
          "$comment": "IRI where the public key associated could be retrieved",
          "$id": "#/properties/creator",
          "examples": [
            "https://example.com/i/pat/keys/"
          ]
        },
        "nonce": {
          "type": "string",
          "$comment": "Some unique value for tracking",
          "$id": "#/properties/nonce",
          "examples": [
            "guid"
          ]
        },
        "signatureFor": {
          "type": "string",
          "$comment": "The attribute name or entity id you for which this is the signature",
          "$id": "#/properties/signatureFor",
          "examples": [
            "http://localhost:8080/serialNum",
            "http://localhost:8080/9cba6ddd-330c-4a0d-929a-771bb12cb0d3"
          ]
        },
        "signatureValue": {
          "type": "string",
          "$comment": "Hash or signed value",
          "$id": "#/properties/signatureValue",
          "examples": [
            "eyiOiJKJeXAasddOEjgFWFXk"
          ]
        }
      },
      "$id": "#/properties/Signature"
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://dummy.oauth.net/auth",
      "tokenUrl": "https://dumy.oauth.net/token",
      "scopes": {
        "admin": "scope of super admin",
        "facillity-admin": "scope of facility admin"
      }
    }
  },
  "security": [
    {
      "hasRole": [
        "admin"
      ]
    }
  ]
}`))
}
