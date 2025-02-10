package goapidoc

var (
	SchemaInt     = Schema{Type: Integer}
	SchemaInt64   = Schema{Type: Integer, Format: "int64"}
	SchemaInt32   = Schema{Type: Integer, Format: "int32"}
	SchemaFloat   = Schema{Type: Number, Format: "float"}
	SchemaDouble  = Schema{Type: Number, Format: "double"}
	SchemaString  = Schema{Type: String}
	SchemaBoolean = Schema{Type: Boolean}
	SchemaNumber  = Schema{Type: Number}
)
