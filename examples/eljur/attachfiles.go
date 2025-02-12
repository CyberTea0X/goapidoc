package main

import "github.com/CyberTea0X/goapidoc"

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
				Description: "Files to attach to the lesson",
				Content: &goapidoc.Content{
					FormData: &goapidoc.ContentSchema{
						Schema: goapidoc.GetSchemaRef(attachFilesInput{}),
					},
				},
			},
			Responses: map[string]goapidoc.Response{
				"200": {
					Description: "Files successfully attached to the lesson",
					Content:     goapidoc.ContentJsonSchemaRef(attachFilesOutput{}),
				},
				"400": Response400,
				"403": Response403,
				"404": *goapidoc.NewResponse("There is no such lesson", errorSchema),
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
