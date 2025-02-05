package document

import "github.com/cybertea0x/goapidoc"

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
			Summary:     "Получить оценки",
			Description: "Получить список оценок студента. given_at вместе с id урока определяет отдельную оценку. given_at это дата в формате unix",
			OperationId: "getGrades",
			Tags:        []string{"auth", "student", "teacher"},
			Parameters: []goapidoc.Parameter{
				authParameter,
				{Name: "student_id", In: "query", Required: true, Description: "id студента"},
			},
			Responses: map[string]goapidoc.Response{
				"200": {
					Description: "Список оценок успешно получен",
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
						Comment:  "Молодец! Активно отвечал и работал на уроке.",
						GivenAt:  int64(UnixDateExample),
					},
					{
						LessonId: 2,
						Grade:    85,
						Comment:  "Хорошая работа, но можно лучше.",
						GivenAt:  int64(UnixDateExample),
					},
				},
			},
		),
	},
}
