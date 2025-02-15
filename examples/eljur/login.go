package main

import "github.com/CyberTea0X/goapidoc"

type loginInput struct {
	Code     string `json:"code" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginOutput struct {
	AccessToken      string `json:"access_token" validate:"required"`
	AccessExpiresAt  int64  `json:"access_expires_at" validate:"required"`
	RefreshToken     string `json:"refresh_token" validate:"required"`
	RefreshExpiresAt int64  `json:"refresh_expires_at" validate:"required"`
}

var loginRoute = Route{
	Name: "/login",
	Path: goapidoc.Path{
		Get: &goapidoc.Method{
			Summary:     "Login",
			Description: "Login by field and password",
			OperationId: "loginUser",
			Tags:        []string{},
			Parameters:  goapidoc.ParametersFromStruct(loginInput{Code: "43124589", Password: "vladikkrutoymalchik2008"}, "query"),
			Responses: map[string]goapidoc.Response{
				"201": goapidoc.ResponseWithJson("The user has successfully logged in.", goapidoc.Ref(loginOutput{})),
				"401": goapidoc.ResponseWithJson("Wrong password or login", errorSchema),
				"400": Response400,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			loginOutput{
				AccessToken:      "thisis.access.token",
				AccessExpiresAt:  1738677032,
				RefreshToken:     "thisis.refresh.token",
				RefreshExpiresAt: 1738677032,
			},
		),
	},
}
