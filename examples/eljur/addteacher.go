package main

import "github.com/CyberTea0X/goapidoc"

type addTeacherInput struct {
	FirstName   string  `json:"first_name" validate:"required"` // Имя учителя
	LastName    string  `json:"last_name" validate:"required"`  // Фамилия учителя
	SubjectIds  []int64 `json:"subject_ids"`                    // Идентификаторы предметов, которые преподает учитель
	Email       string  `json:"email" validate:"required"`      // Электронная почта учителя
	PhoneNumber string  `json:"phone_number"`                   // Номер телефона учителя (опционально)
}

type addTeacherOutput struct {
	TeacherId int64  `json:"teacher_id" validate:"required"` // Идентификатор добавленного учителя
	Code      string `json:"code" validate:"required"`       // Код для доступа (например, для входа в систему)
	Password  string `json:"password" validate:"required"`   // Пароль для доступа
}

var addTeacherRoute = Route{
	Name: "/teacher",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Add new teacher",
			Security:    BearerSecurity,
			Description: "Adds new teacher into system",
			OperationId: "addTeacher",
			Tags:        []string{"auth", "admin"},
			RequestBody: goapidoc.RequestWithJson("Teacher info", goapidoc.Ref(addTeacherInput{}), true),
			Responses: map[string]goapidoc.Response{
				"201": goapidoc.ResponseWithJson("Teacher added", goapidoc.Ref(addTeacherOutput{})),
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			addTeacherInput{
				FirstName:   "Ivan",
				LastName:    "Ivanov",
				SubjectIds:  []int64{101, 102}, // Математика и Физика
				Email:       "ivan.ivanov@school.com",
				PhoneNumber: "+79123456789",
			},
			addTeacherOutput{
				TeacherId: 123,
				Code:      "teacher123",
				Password:  "securepassword123",
			},
		),
	},
}
