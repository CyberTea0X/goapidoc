package goapidoc

import "reflect"

type OapiType string

const (
	Integer OapiType = "integer"
	Boolean OapiType = "boolean"
	Number  OapiType = "number"
	String  OapiType = "string"
	Object  OapiType = "object"
	Array   OapiType = "array"
)

func toOapiType(t reflect.Type) OapiType {
	switch t.Kind() {
	case reflect.Int:
		return Integer
	case reflect.Int32:
		return Integer
	case reflect.Int64:
		return Integer
	case reflect.Uint:
		return Integer
	case reflect.Uint64:
		return Integer
	case reflect.Uint32:
		return Integer
	case reflect.Bool:
		return Boolean
	case reflect.Float64:
		return Number
	case reflect.Float32:
		return Number
	case reflect.String:
		return String
	case reflect.Slice:
		return Array
	case reflect.Struct:
		if t.Name() == "Time" {
			return String
		}
		return Object
	}
	panic("unhandled type: " + t.String())
}
