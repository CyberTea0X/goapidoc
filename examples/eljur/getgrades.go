package main

import "github.com/CyberTea0X/goapidoc"

type getGradesOutput struct {
	Grades []Grade `json:"grades"`
}

type Grade struct {
	LessonId int64  `json:"lesson_id"`
	Grade    int64  `json:"grade"`
	GivenAt  int64  `json:"given_at"`
	Comment  string `json:"comment"`
}

var getGradesRoute = Route{
	Name: "/grades",
	Path: goapidoc.Path{
		Get: &goapidoc.Method{
			Summary:     "Get ratings",
			Description: "Get a list of student grades. given_at together with lesson id defines a single grade. given_at is a unix date",
			OperationId: "getGrades",
			Tags:        []string{"auth", "student", "teacher"},
			Parameters: []goapidoc.Parameter{
				authParameter,
				{Name: "student_id", In: "query", Required: true, Description: "id студента"},
			},
			Responses: map[string]goapidoc.Response{
				"200": {
					Description: "The list of ratings has been successfully received.",
					Content:     goapidoc.ContentJsonSchemaRef(getGradesOutput{}),
				},
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			getGradesOutput{
				Grades: []Grade{
					{
						LessonId: 1,
						Grade:    95,
						Comment:  "Well done! You answered actively and worked in class.",
						GivenAt:  int64(UnixDateExample),
					},
					{
						LessonId: 2,
						Grade:    85,
						Comment:  "Good job, but could be better.",
						GivenAt:  int64(UnixDateExample),
					},
				},
			},
		),
	},
}
