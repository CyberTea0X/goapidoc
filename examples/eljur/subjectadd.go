package main

import "github.com/CyberTea0X/goapidoc"

type subjectAddInput struct {
	Name          string `json:"name" validate:"required"`
	Desription    string `json:"description"`
	Public        bool   `json:"public" validate:"required"`
	RequiresGroup string `json:"requires_group"`
	Instructor    int64  `json:"instructor" validate:"required"`
	Capacity      int    `json:"capacity" validate:"required"`
	IsActive      bool   `json:"is_active"`
}

type subjectAddOutput struct {
	SubjectCode int64  `json:"subject_code" validate:"required"`
	Url         string `json:"url" validate:"required"`
}

var subjectAddRoute = Route{
	Name: "/subject",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Add item",
			Security:    BearerSecurity,
			Description: "Add a new subject",
			OperationId: "subjectAdd",
			Tags:        []string{"auth", "teacher"},
			RequestBody: goapidoc.RequestWithJson("Item Parameters", goapidoc.Ref(subjectAddInput{}), true),
			Responses: map[string]goapidoc.Response{
				"201": goapidoc.ResponseWithJson("Code and link to the item page", goapidoc.Ref(subjectAddOutput{})),
				"400": Response400,
				"403": Response403,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			subjectAddInput{
				Name:          "Mathematics",
				Desription:    "The Science of Numbers and Shapes",
				Public:        true,
				RequiresGroup: "9 класс",
				Instructor:    13,
				Capacity:      30,
				IsActive:      true,
			},
			subjectAddOutput{
				SubjectCode: 1341,
				Url:         "http://example.com/subjects/1341",
			},
		),
	},
}
