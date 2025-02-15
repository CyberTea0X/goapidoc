package main

import "github.com/CyberTea0X/goapidoc"

type attachAttendanceInput struct {
	LessonId   int64              `json:"lesson_id" validate:"required"`
	Attendance []attendanceRecord `json:"attendance" validate:"required"`
}

type attendanceRecord struct {
	StudentId int64 `json:"student_id" validate:"required"`
	Presence  bool  `json:"presence" validate:"required"`
}

var attachAttendanceRoute = Route{
	Name: "/attendance",
	Path: goapidoc.Path{
		Put: &goapidoc.Method{
			Summary:     "Attach attendance",
			Security:    BearerSecurity,
			Description: "Attaches the attendance list to the lesson. If attendance was already connected for the student, it is updated",
			OperationId: "attachAttendance",
			Tags:        []string{"auth", "teacher"},
			RequestBody: goapidoc.RequestWithJson(
				"List of records where the lesson and the fact of attendance/absence are indicated",
				goapidoc.Ref(attachAttendanceInput{}),
			),
			Responses: map[string]goapidoc.Response{
				"200": Response200Empty,
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			attachAttendanceInput{
				LessonId: 1,
				Attendance: []attendanceRecord{
					{StudentId: 1, Presence: true},
					{StudentId: 2, Presence: true},
					{StudentId: 2, Presence: false},
				},
			},
		),
	},
}
