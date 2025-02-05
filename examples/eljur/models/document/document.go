package document

import (
	"github.com/cybertea0x/goapidoc"
	oapi "github.com/cybertea0x/goapidoc"
)

type Route struct {
	Name       string
	Path       oapi.Path
	Components oapi.Components
}

func addRoute(document oapi.Document, route Route) {
	document.Components.Schemas.Merge(route.Components.Schemas)
	document.AddPath(route.Name, route.Path)
}

var authParameter = oapi.Parameter{
	Name:        "Authorization",
	In:          "header",
	Description: "Authentication parameter",
	Required:    true,
	Schema: goapidoc.Schema{
		Type:    goapidoc.String,
		Example: "Bearer <token>",
	},
}

var errorSchema = oapi.Schema{
	Type: "object",
	Properties: map[string]any{
		"error": oapi.Schema{
			Type:    "string",
			Example: "Описание ошибки",
		},
	},
}
var infoOk = oapi.Schema{
	Type: "object",
	Properties: map[string]any{
		"info": oapi.Schema{
			Type:    oapi.String,
			Example: "Всё хорошо",
		},
	},
}

var Response200Empty = oapi.Response{Description: "Операция успешно выполнена. Возвращать результат смысла нет."}
var Response201Empty = oapi.Response{Description: "Ресурс успешно создан. Возвращать смысла нет"}
var Response204 = oapi.Response{Description: "No content"}
var Response400 = *oapi.NewResponse("Неверный ввод", errorSchema)
var Response403 = oapi.Response{Description: "Доступ запрещён"}
var Response409 = oapi.Response{Description: "Конфликт с другим ресурсом"}
var Response500 = oapi.Response{Description: "Ошибка сервера"}
var InfoResponse = *oapi.NewResponse("Операция успешно. Json содержит info с доп.информацией", infoOk)
var UnixDateExample = int64(10250045)

func BuildDocument() *oapi.Document {
	doc := oapi.Document{
		OpenApiVersion: "3.1.0",
		Info: oapi.Info{
			Title:       "Сервис электронного журнала",
			Description: "Сервис обеспечивающий работу электронного журнала",
			Version:     "1.0",
			Contact:     map[string]string{},
			License:     &oapi.License{},
		},
		Tags: []oapi.Tag{
			{Name: "auth", Description: "authentication required"},
			{Name: "teacher", Description: "teacher role required"},
			{Name: "admin", Description: "admin role required"},
		},
		Paths:      map[string]oapi.Path{},
		Components: &oapi.Components{Schemas: oapi.Schemas{}},
	}
	addRoute(doc, registerRoute)
	addRoute(doc, loginRoute)
	addRoute(doc, addStudentRoute)
	addRoute(doc, getStudentsRoute)
	addRoute(doc, addClassRoute)
	addRoute(doc, subjectAddRoute)
	addRoute(doc, addLessonRoute)
	addRoute(doc, attachFilesRoute)
	addRoute(doc, attachAttendanceRoute)
	addRoute(doc, giveGradeRoute)
	addRoute(doc, getGradesRoute)
	addRoute(doc, addTeacherRoute)
	return &doc
}
