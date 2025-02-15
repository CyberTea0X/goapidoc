package main

import "github.com/CyberTea0X/goapidoc"

type getStudentsInput struct {
	ClassId int64 `json:"class_id" validate:"required"`
}

type getStudentsOutput struct {
	Students []Student `json:"students" validate:"required"`
}

type Student struct {
	Id       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Mail     string `json:"mail" validate:"required"`
}

var getStudentsRoute = Route{
	Name: "/students",
	Path: goapidoc.Path{
		Get: &goapidoc.Method{
			Summary:     "Get list of students",
			Security:    BearerSecurity,
			Description: "Get a list of all students in a specified class",
			OperationId: "getStudents",
			Tags:        []string{"auth", "teacher"},
			Responses: map[string]goapidoc.Response{
				"200": goapidoc.ResponseWithJson("List of students successfully received", goapidoc.Ref(getStudentsOutput{})),
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
