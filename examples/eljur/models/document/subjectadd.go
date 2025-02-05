package document

import "github.com/cybertea0x/goapidoc"

type subjectAddInput struct {
	Name          string `json:"name,required"`
	Desription    string `json:"description"`
	Public        bool   `json:"public,required"`
	RequiresGroup string `json:"requires_group"`
	Instructor    int64  `json:"instructor,required"`
	Capacity      int    `json:"capacity,required"`
	IsActive      bool   `json:"is_active"`
}

type subjectAddOutput struct {
	SubjectCode int64  `json:"subject_code,required"`
	Url         string `json:"url,required"`
}

var subjectAddRoute = Route{
	Name: "/subject",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Добавить предмет",
			Description: "Добавить новый учебный предмет",
			OperationId: "subjectAdd",
			Tags:        []string{"auth", "teacher"},
			Parameters:  []goapidoc.Parameter{authParameter},
			RequestBody: goapidoc.NewRequestBody("Параметры предмета", goapidoc.ContentJsonSchemaRef(subjectAddInput{})),
			Responses: map[string]goapidoc.Response{
				"201": *goapidoc.NewResponse("Код и ссылка на страницу предмета", goapidoc.GetSchemaRef(subjectAddOutput{})),
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			subjectAddInput{
				Name:          "Математика",
				Desription:    "Наука о числах и фигурах",
				Public:        true,
				RequiresGroup: "9 класс",
				Instructor:    13,
				Capacity:      30,
				IsActive:      true,
			},
			subjectAddOutput{
				SubjectCode: 1341,
				Url:         "http://example.com/subjects/1341",
			},
		),
	},
}
