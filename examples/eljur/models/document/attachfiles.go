package document

import "github.com/cybertea0x/goapidoc"

type attachFilesInput struct {
	LessonId int64    `form:"lesson_id,required"` // Идентификатор урока, к которому прикрепляются файлы
	Files    []string `form:"files,required"`     // Список файлов для загрузки
}

type attachFilesOutput struct {
	Success bool `json:"success,required"` // Успешно ли прикреплены файлы
}

var attachFilesRoute = Route{
	Name: "/lesson/attach-files",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Attach files to lesson",
			Description: "Attach files to a specific lesson",
			OperationId: "attachFilesToLesson",
			Tags:        []string{"auth", "teacher"},
			Parameters:  []goapidoc.Parameter{authParameter},
			RequestBody: &goapidoc.RequestBody{
				Description: "Файлы для прикрепления к уроку",
				Content: &goapidoc.Content{
					FormData: &goapidoc.ContentSchema{
						Schema: goapidoc.GetSchemaRef(attachFilesInput{}),
					},
				},
			},
			Responses: map[string]goapidoc.Response{
				"200": {
					Description: "Файлы успешно прикреплены к уроку",
					Content:     goapidoc.ContentJsonSchemaRef(attachFilesOutput{}),
				},
				"400": Response400,
				"403": Response403,
				"404": *goapidoc.NewResponse("Такого урока не существует", errorSchema),
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			attachFilesInput{
				LessonId: 1,
				Files:    []string{"file1.pdf", "file2.jpg"},
			},
			attachFilesOutput{
				Success: true,
			},
		),
	},
}
