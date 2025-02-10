package document

import "github.com/CyberTea0X/goapidoc"

type addStudentInput struct {
	ClassId int64 `json:"class_id,required"`
}

type addStudentOutput struct {
	Code     string `json:"code,required"`
	Password string `json:"password,required"`
}

var addStudentRoute = Route{
	Name: "/student",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Add new student",
			Description: "Add new student to the specified class",
			OperationId: "addStudent",
			Tags:        []string{"auth", "teacher"},
			Parameters:  []goapidoc.Parameter{authParameter},
			RequestBody: &goapidoc.RequestBody{
				Description: "student info",
				Content:     goapidoc.ContentJsonSchemaRef(addStudentInput{}),
			},
			Responses: map[string]goapidoc.Response{
				"201": {
					Description: "Student successfully added into class",
					Content:     goapidoc.ContentJsonSchemaRef(addStudentOutput{}),
				},
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			addStudentInput{
				ClassId: 415,
			},
			addStudentOutput{
				Code:     "14843928",
				Password: "vladikbrutal2009",
			},
		),
	},
}
