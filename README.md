# goapidoc

This is a simple and fast library for writing OpenAPI specification in Go.

## Main features

 - Typesafety
 - Struct to component conversion
 - Write examples as golang structs
 - Supports openapi v 3.1.0

## Examples

Some golang struct

```golang
//...
type AddProfileInput struct {
	Name      string            `json:"name" binding:"required"`
	Birthdate time.Time         `json:"birthdate" binding:"required"`
	Contacts  string            `json:"contacts"`
	About     string            `json:"about"`
	Gender    string            `json:"gender" binding:"required"`
	City      string            `json:"city" binding:"required"`
	Position  *models.PointJson `json:"position" binding:"required"`
	Tags      []int64           `json:"tags" binding:"required"`
}

var someProfile = AddProfileInput{
	Name:      "Anatoliy",
	Birthdate: time.Now().AddDate(18, 0, 0),
	Contacts:  "+1234566789",
	About:     "Golang programmer",
	Gender:    "male",
	City:      "Some city",
	Position:  &models.PointJson{X: 0, Y: 0},
	Tags:      []int64{1, 4, 10},
}
//...
```

generated openapi docs for this:
```json
"AddProfileInput": {
    "type": "object",
    "properties": {
        "about": {
            "type": "string",
            "example": "Golang programmer"
        },
        "birthdate": {
            "type": "string",
            "example": "2043-01-25T21:23:50.373924787+03:00"
        },
        "city": {
            "type": "string",
            "example": "Some city"
        },
        "contacts": {
            "type": "string",
            "example": "+1234566789"
        },
        "gender": {
            "type": "string",
            "example": "male"
        },
        "name": {
            "type": "string",
            "example": "Anatoliy"
        },
        "position": {
            "type": "object",
            "example": {
                "type": "object",
                "properties": {
                    "x": {
                        "type": "number",
                        "example": 0
                    },
                    "y": {
                        "type": "number",
                        "example": 0
                    }
                }
            }
        },
        "tags": {
            "type": "array",
            "example": [
                1,
                4,
                10
            ],
            "items": {
                "type": "integer"
            }
        }
    }
},
```

check out petstore example
