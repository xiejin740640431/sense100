package swag

import "fmt"

func CheckSchemaType(typeName string) {
	switch typeName {
	case "string", "number", "integer", "boolean", "array", "object":

	default:
		panic(fmt.Errorf("%s is not basic types", typeName))
		return
	}

}

func TransToValidSchemeType(typeName string) string {
	switch typeName {
	case "uint", "int", "uint8", "int8", "uint16", "int16", "byte":
		return "integer"
	case "uint32", "int32", "rune":
		return "integer"
	case "uint64", "int64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	case "string":
		return "string"
	case "JsonDate","JsonTime","JsonDateTime","JsonTimeHS":
		return "string"
	default:
		panic(fmt.Errorf("%s is not valid go basic types", typeName))
	}
}
