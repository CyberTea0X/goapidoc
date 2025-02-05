package document

import "github.com/cybertea0x/goapidoc"

type addClassInput struct {
	Name        string `json:"name,required"`
	Description string `json:"description"` // Описание класса (например, "Математический класс")
	TeacherID   int64  `json:"teacher_id"`  // Идентификатор учителя, который будет вести класс
}

type addClassOutput struct {
	ClassId int64 `json:"class_id"`
}

var addClassRoute = Route{
	Name: "/class",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Add new class",
			Description: "Adds new class",
			OperationId: "addClass",
			Tags:        []string{"auth"},
			Parameters:  []goapidoc.Parameter{authParameter},
			RequestBody: &goapidoc.RequestBody{
				Description: "Class info",
				Content:     goapidoc.ContentJsonSchemaRef(addClassInput{}),
			},
			Responses: map[string]goapidoc.Response{
				"201": {
					Description: "Класс успешно добавлен",
					Content:     goapidoc.ContentJsonSchemaRef(addClassOutput{}),
				},
				"400": Response400,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			addClassInput{Name: "8-A", Description: "8 класс школы такой-то", TeacherID: 542},
			addClassOutput{ClassId: 142},
		),
	},
}
