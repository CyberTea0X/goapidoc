package document

import "github.com/cybertea0x/goapidoc"

type giveGradeInput struct {
	LessonId  int64  `json:"lesson_id,required"`
	StudentId int64  `json:"student_id,required"`
	Grade     int64  `json:"grade,required"`
	Comment   string `json:"comment"`
}

type giveGradeOutput struct {
	GivenAt int64 `json:"given_at"`
}

var giveGradeRoute = Route{
	Name: "/grade",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Поставить оценку",
			Description: "Поставить оценку",
			OperationId: "giveGrade",
			Tags:        []string{"auth", "teacher"},
			Parameters:  []goapidoc.Parameter{authParameter},
			RequestBody: &goapidoc.RequestBody{
				Description: "Описание оценки",
			},
			Responses: map[string]goapidoc.Response{
				"201": *goapidoc.NewResponse(
					"Оценка поставлена. given_at это unix дата когда она была выставлена, возвращается",
					goapidoc.MustBuildSchemaFrom(giveGradeOutput{GivenAt: UnixDateExample}),
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
				Comment:   "Молодец! Активно отвечал и работал на уроке.",
			},
		),
	},
}
