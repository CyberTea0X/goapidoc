package main

import (
	"github.com/CyberTea0X/goapidoc"
	oapi "github.com/CyberTea0X/goapidoc"
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
			Example: "Error description",
		},
	},
}
var infoOk = oapi.Schema{
	Type: "object",
	Properties: map[string]any{
		"info": oapi.Schema{
			Type:    oapi.String,
			Example: "Everything is fine",
		},
	},
}

var Response200Empty = oapi.Response{Description: "The operation was completed successfully. There is no point in returning the result."}
var Response201Empty = oapi.Response{Description: "The resource has been successfully created. There is no point in returning it."}
var Response204 = oapi.Response{Description: "No content"}
var Response400 = oapi.ResponseWithJson("Invalid input", errorSchema)
var Response403 = oapi.Response{Description: "Access Denied"}
var Response409 = oapi.Response{Description: "Conflict with another resource"}
var Response500 = oapi.Response{Description: "Server error"}
var InfoResponse = oapi.ResponseWithJson("Operation successful. Json contains info with additional information", infoOk)
var UnixDateExample = int64(10250045)

func BuildDocument() *oapi.Document {
	doc := oapi.Document{
		OpenApiVersion: "3.1.0",
		Info: oapi.Info{
			Title:       "Electronic journal service",
			Description: "Service providing operation of the electronic journal",
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
