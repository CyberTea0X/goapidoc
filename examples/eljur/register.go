package main

import "github.com/CyberTea0X/goapidoc"

type registerInput struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Mail     string `json:"mail" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type registerOutput struct {
	Id string `json:"id" validate:"required"`
}

var registerRoute = Route{
	Name: "/register",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Registration",
			Description: "New User Registration",
			OperationId: "registerUser",
			Tags:        []string{},
			Parameters:  []goapidoc.Parameter{},
			RequestBody: goapidoc.RequestWithJson("User data", goapidoc.Ref(registerInput{})),
			Responses: map[string]goapidoc.Response{
				"201": goapidoc.ResponseWithJson("User registered successfully", goapidoc.Ref(registerOutput{})),
				"400": Response400,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			registerInput{
				Name:     "Vladislav",
				Surname:  "Petrov",
				Mail:     "vladikdetskysadic@example.com",
				Code:     "43124589",
				Password: "vladikkrutoymalchik2008",
			},
			registerOutput{
				Id: "43124589",
			},
		),
	},
}
