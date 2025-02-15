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
			Security:    BearerSecurity,
			Description: "Attach files to a specific lesson",
			OperationId: "attachFilesToLesson",
			Tags:        []string{"auth", "teacher"},
			RequestBody: &goapidoc.RequestBody{
				Description: "Files to attach to the lesson",
				Content: &goapidoc.Content{
					FormData: &goapidoc.ContentSchema{
						Schema: goapidoc.Ref(attachFilesInput{}),
					},
				},
			},
			Responses: map[string]goapidoc.Response{
				"200": goapidoc.ResponseWithJson("Files successfully attached to the lesson", goapidoc.Ref(attachFilesOutput{})),
				"400": Response400,
				"403": Response403,
				"404": goapidoc.ResponseWithJson("There is no such lesson", errorSchema),
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
