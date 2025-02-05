package document

import "github.com/cybertea0x/goapidoc"

type registerInput struct {
	Name     string `json:"name,required"`
	Surname  string `json:"surname,required"`
	Phone    string `json:"phone,required"`
	Mail     string `json:"mail,required"`
	Code     string `json:"code,required"`
	Password string `json:"password,required"`
}

type registerOutput struct {
	Id string `json:"id,required"`
}

var registerRoute = Route{
	Name: "/register",
	Path: goapidoc.Path{
		Post: &goapidoc.Method{
			Summary:     "Регистрация",
			Description: "Регистрация нового пользователя",
			OperationId: "registerUser",
			Tags:        []string{},
			Parameters:  []goapidoc.Parameter{},
			RequestBody: goapidoc.NewRequestBody("Пользовательские данные", goapidoc.ContentJsonSchemaRef(registerInput{})),
			Responses: map[string]goapidoc.Response{
				"201": *goapidoc.NewResponse("Пользователь успешно зарегистрирован", goapidoc.GetSchemaRef(registerOutput{})),
				"400": Response400,
				"500": Response500,
			},
		},
	},
	Components: goapidoc.Components{
		Schemas: goapidoc.SchemasOf(
			registerInput{
				Name:     "Владислав",
				Surname:  "Петров",
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
