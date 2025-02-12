package main

import "github.com/CyberTea0X/goapidoc"

type addLessonInput struct {
	Index          int    `json:"index,required"`
	Name           string `json:"name,required"`
	Description    string `json:"description"`
	IsOnline       bool   `json:"is_online"`     // Проводится ли урок онлайн (true/false)
	Date           string `json:"date,required"` // Дата проведения урока (например, "2023-10-15")
	SubjectId      int64  `json:"subject_id,required"`
	ClassId        int64  `json:"class_id"`        // Идентификатор класса, для которого проводится урок
	TeacherId      int64  `json:"teacher_id"`      // Идентификатор учителя, который ведет
	ScheduleNumber int    `json:"schedule_number"` // Указывает каким по номеру в расписании идёт урок
}

type addLessonOutput struct {
	LessonId int64 `json:"lesson_id,required"`
}

var addLessonRoute = Route{
	Name: "/lesson",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Create new lesson",
			Description: "Create new lesson",
			OperationId: "createLesson",
			Tags:        []string{"auth", "teacher"},
			Parameters:  []goapidoc.Parameter{authParameter},
			RequestBody: goapidoc.RequestWithJson("Lesson info", goapidoc.Ref(addLessonInput{})),
			Responses: map[string]goapidoc.Response{
				"201": goapidoc.ResponseWithJson("Lesson successfully added", goapidoc.Ref(addLessonOutput{})),
				"400": Response400,
				"403": Response403,
				"409": Response409,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			addLessonInput{
				Index:          1,
				Name:           "Math first lesson",
				Description:    "Addition and Subtraction",
				IsOnline:       true,
				Date:           "2023-10-15",
				SubjectId:      1,
				ClassId:        1,
				TeacherId:      1,
				ScheduleNumber: 1,
			},
			addLessonOutput{
				LessonId: 1,
			},
		),
	},
}
