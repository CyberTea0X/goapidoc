package main

import "github.com/CyberTea0X/goapidoc"

type addTeacherInput struct {
	FirstName   string  `json:"first_name,required"` // Имя учителя
	LastName    string  `json:"last_name,required"`  // Фамилия учителя
	SubjectIds  []int64 `json:"subject_ids"`         // Идентификаторы предметов, которые преподает учитель
	Email       string  `json:"email,required"`      // Электронная почта учителя
	PhoneNumber string  `json:"phone_number"`        // Номер телефона учителя (опционально)
}

type addTeacherOutput struct {
	TeacherId int64  `json:"teacher_id,required"` // Идентификатор добавленного учителя
	Code      string `json:"code,required"`       // Код для доступа (например, для входа в систему)
	Password  string `json:"password,required"`   // Пароль для доступа
}

var addTeacherRoute = Route{
	Name: "/teacher",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Add new teacher",
			Description: "Adds new teacher into system",
			OperationId: "addTeacher",
			Tags:        []string{"auth", "admin"},
			Parameters:  []goapidoc.Parameter{authParameter},
			RequestBody: &goapidoc.RequestBody{
				Description: "Teacher info",
				Content:     goapidoc.ContentJsonSchemaRef(addTeacherInput{}),
			},
			Responses: map[string]goapidoc.Response{
				"201": {
					Description: "Teacher added",
					Content:     goapidoc.ContentJsonSchemaRef(addTeacherOutput{}),
				},
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
