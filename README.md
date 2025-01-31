# goapidoc

This is a simple and fast library for writing OpenAPI specification in Go.

## Main features

 - Typesafety
 - Struct to component conversion
 - Write examples as golang structs
 - Supports openapi v 3.1.0

## Examples

### Write doc in golang

```golang
  type Pet struct {
  	Id   int64  `json:"id" binding:"required"`
  	Name string `json:"name" binding:"required"`
  	Tag  string `json:"tag"`
  }

  type Pets []Pet

	doc := oapi.Document{
		OpenApiVersion: "3.1.0",
		Info: oapi.Info{
			Version: "1.0.0",
			Title:   "Swagger Petstore",
			License: &oapi.License{
				Name: "MIT",
				Url:  "https://opensource.org/licenses/MIT",
			},
		},
		Servers: []oapi.Server{
			{Url: "http://petstore.swagger.io/v1"},
		},
		Paths: map[string]oapi.Path{
			"/pets": {
				Get: &oapi.Method{
					Summary:     "List all pets",
					OperationId: "listPets",
					Tags:        []string{"pets"},
					Parameters: []oapi.Parameter{
						{
							Name:        "limit",
							In:          "query",
							Description: "How many items to return at one time (max 100)",
							Required:    false,
							Schema:      oapi.MustBuildSchemaFrom(int32(1)),
						},
					},
					Responses: map[string]oapi.Response{
						"200": {
							Description: "A paged array of pets",
							Headers: map[string]oapi.Header{
								"x-next": {
									Description: "A link to the next page of responses",
									Schema: oapi.Schema{
										Type: oapi.String,
									},
								},
							},
							Content: oapi.ContentJsonSchemaRef(Pets{}),
						},
						"default": DefaultResponse,
					},
				},
			},
			Components: &oapi.Components{
			Schemas: oapi.SchemasOf(
				Pet{
					Id:   1,
					Name: "Dog",
					Tag:  "dogs",
				},
				Pets{{Id: 1, Name: "Dog", Tag: "dogs"}, {Id: 2, Name: "Cat", Tag: "cats"}},
				Error{
					Code:    500,
					Message: "server crushed",
				},
			),
		},
	}
	doc.SaveAsJson("petstore.json")
```

Spec:

```json
{
  "openapi": "3.1.0",
  "info": {
    "title": "Swagger Petstore",
    "version": "1.0.0",
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/MIT"
    }
  },
  "paths": {
    "/pets": {
      "get": {
        "summary": "List all pets",
        "operationId": "listPets",
        "tags": [
          "pets"
        ],
        "parameters": [
          {
            "name": "limit",
            "in": "query",
            "description": "How many items to return at one time (max 100)",
            "required": false,
            "schema": {
              "type": "integer",
              "format": "int32",
              "example": 1
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A paged array of pets",
            "headers": {
              "x-next": {
                "description": "A link to the next page of responses",
                "schema": {
                  "type": "string"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Pets"
                }
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Error"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "Error": {
        "type": "object",
        "properties": {
          "code": {
            "type": "integer",
            "format": "int32",
            "example": 500
          },
          "message": {
            "type": "string",
            "example": "server crushed"
          }
        }
      },
      "Pet": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer",
            "format": "int64",
            "example": 1
          },
          "name": {
            "type": "string",
            "example": "Dog"
          },
          "tag": {
            "type": "string",
            "example": "dogs"
          }
        }
      },
      "Pets": {
        "type": "array",
        "items": {
          "type": "object"
        },
        "example": [
          {
            "id": 1,
            "name": "Dog",
            "tag": "dogs"
          },
          {
            "id": 2,
            "name": "Cat",
            "tag": "cats"
          }
        ]
      }
    }
  },
  "servers": [
    {
      "url": "http://petstore.swagger.io/v1"
    }
  ]
}
```


### Convert golang struct into spec

```golang
//...
type AddProfileInput struct {
	Name      string            `json:"name" binding:"required"`
	Birthdate time.Time         `json:"birthdate" binding:"required"`
	Contacts  string            `json:"contacts"`
	About     string            `json:"about"`
	Gender    string            `json:"gender" binding:"required"`
	City      string            `json:"city" binding:"required"`
	Position  *models.PointJson `json:"position" binding:"required"`
	Tags      []int64           `json:"tags" binding:"required"`
}

var someProfile = AddProfileInput{
	Name:      "Anatoliy",
	Birthdate: time.Now().AddDate(18, 0, 0),
	Contacts:  "+1234566789",
	About:     "Golang programmer",
	Gender:    "male",
	City:      "Some city",
	Position:  &models.PointJson{X: 0, Y: 0},
	Tags:      []int64{1, 4, 10},
}
//...
```

generated openapi docs for this:
```json
"AddProfileInput": {
    "type": "object",
    "properties": {
        "about": {
            "type": "string",
            "example": "Golang programmer"
        },
        "birthdate": {
            "type": "string",
            "example": "2043-01-25T21:23:50.373924787+03:00"
        },
        "city": {
            "type": "string",
            "example": "Some city"
        },
        "contacts": {
            "type": "string",
            "example": "+1234566789"
        },
        "gender": {
            "type": "string",
            "example": "male"
        },
        "name": {
            "type": "string",
            "example": "Anatoliy"
        },
        "position": {
            "type": "object",
            "properties": {
                "x": {
                    "type": "number",
                    "example": 44.957813
                },
                "y": {
                    "type": "number",
                    "example": 34.109547
                }
            }
        },
        "tags": {
            "type": "array",
            "items": {
                "type": "integer"
            },
            "example": [
                1,
                4,
                10
            ]
        }
    }
}
```

check out petstore example
