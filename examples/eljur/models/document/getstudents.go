package document

import "github.com/cybertea0x/goapidoc"

type getStudentsInput struct {
	ClassId int64 `json:"class_id,required"`
}

type getStudentsOutput struct {
	Students []Student `json:"students,required"`
}

type Student struct {
	Id       int64  `json:"id,required"`
	Name     string `json:"name,required"`
	Lastname string `json:"lastname,required"`
	Phone    string `json:"phone,required"`
	Mail     string `json:"mail,required"`
}

var getStudentsRoute = Route{
	Name: "/students",
	Path: goapidoc.Path{
		Get: &goapidoc.Method{
			Summary:     "Get list of students",
			Description: "Get a list of all students in a specified class",
			OperationId: "getStudents",
			Tags:        []string{"auth", "teacher"},
			Parameters:  []goapidoc.Parameter{authParameter},
			Responses: map[string]goapidoc.Response{
				"200": {
					Description: "List of students successfully received",
					Content:     goapidoc.ContentJsonSchemaRef(getStudentsOutput{}),
				},
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			getStudentsInput{
				ClassId: 415,
			},
			getStudentsOutput{
				Students: []Student{
					{
						Id:       1,
						Name:     "Ivan",
						Lastname: "Ivanov",
						Phone:    "+1234566789",
						Mail:     "vlomchetodumat@example.com",
					},
					{
						Id:       2,
						Name:     "Petr",
						Lastname: "Petrov",
						Phone:    "+1234566789",
						Mail:     "vlomchetodumat@example.com",
					},
				},
			},
		),
	},
}
