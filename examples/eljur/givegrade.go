package main

import "github.com/CyberTea0X/goapidoc"

type giveGradeInput struct {
	LessonId  int64  `json:"lesson_id" validate:"required"`
	StudentId int64  `json:"student_id" validate:"required"`
	Grade     int64  `json:"grade" validate:"required"`
	Comment   string `json:"comment"`
}

type giveGradeOutput struct {
	GivenAt int64 `json:"given_at"`
}

var giveGradeRoute = Route{
	Name: "/grade",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Rate this",
			Security:    BearerSecurity,
			Description: "Rate this",
			OperationId: "giveGrade",
			Tags:        []string{"auth", "teacher"},
			RequestBody: &goapidoc.RequestBody{
				Description: "Rate this",
			},
			Responses: map[string]goapidoc.Response{
				"201": goapidoc.ResponseWithJson(
					"The rating has been set. given_at is the unix date when it was set, returned",
					goapidoc.SchemaFrom(giveGradeOutput{GivenAt: UnixDateExample}),
				),
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			giveGradeInput{
				LessonId:  1,
				StudentId: 1,
				Grade:     95,
				Comment:   "Well done! You answered actively and worked in class.",
			},
		),
	},
}
