package main

import (
	oapi "github.com/CyberTea0X/goapidoc"
)

type Pet struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Tag  string `json:"tag"`
}

type Pets []Pet

type Error struct {
	Code    int32  `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
}

var DefaultResponse = oapi.ResponseWithJson("unexpected error", oapi.Ref(Error{}))

func main() {
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
							Schema:      oapi.SchemaFrom(int32(1)),
						},
					},
					Responses: map[string]oapi.Response{
						"200": oapi.ResponseWithJson("A paged array of pets", oapi.Ref(Pets{})).WithHeaders(
							map[string]oapi.Header{"x-next": {
								Description: "A link to the next page of responses",
								Schema: oapi.Schema{
									Type: oapi.String,
								},
							},
							},
						),
						"default": DefaultResponse,
					},
				},
				Post: &oapi.Method{
					Summary:     "Create a pet",
					OperationId: "createPets",
					Tags:        []string{"pets"},
					Responses: map[string]oapi.Response{
						"201": {
							Description: "Null response",
						},
						"default": DefaultResponse,
					},
				},
			},
			"/pets/{petId}": {
				Get: &oapi.Method{
					Summary:     "Info for a specific pet",
					OperationId: "showPetById",
					Tags:        []string{"pets"},
					Parameters: []oapi.Parameter{
						{
							Name:        "petId",
							In:          "path",
							Required:    true,
							Description: "The id of the pet to retrieve",
							Schema:      oapi.Schema{Type: oapi.String},
						},
					},
					Responses: map[string]oapi.Response{
						"200":     oapi.ResponseWithJson("Expected response to a valid request", oapi.Ref(Pet{})),
						"default": DefaultResponse,
					},
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
}
