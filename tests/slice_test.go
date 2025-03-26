package tests

import (
	"testing"

	"github.com/CyberTea0X/goapidoc"
)

func TestEmptySlice(t *testing.T) {
	emptySlice := []int{}
	emptySchema := goapidoc.SchemaFrom(emptySlice)
	if emptySchema.Type != goapidoc.Array {
		t.Errorf("Expected empty slice schema type to be Array, got %v", emptySchema.Type)
	}
}

func TestStringSlice(t *testing.T) {
	strSlice := []string{"one", "two", "three"}
	strSchema := goapidoc.SchemaFrom(strSlice)
	if strSchema.Type != goapidoc.Array {
		t.Errorf("Expected string slice schema type to be Array, got %v", strSchema.Type)
	}

	if len(strSchema.Example.([]string)) != 3 {
		t.Errorf("Expected string slice schema example to be ['one', 'two', 'three'], got %v", strSchema.Example)
	}
}

func TestIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4}
	intSchema := goapidoc.SchemaFrom(intSlice)
	if intSchema.Type != goapidoc.Array {
		t.Errorf("Expected int slice schema type to be Array, got %v", intSchema.Type)
	}
}

func TestInt64Slice(t *testing.T) {
	int64Slice := []int64{1, 2, 3, 4}
	int64Schema := goapidoc.SchemaFrom(int64Slice)
	if int64Schema.Type != goapidoc.Array {
		t.Errorf("Expected int64 slice schema type to be Array, got %v", int64Schema.Type)
	}
	if int64Schema.Items.Format != goapidoc.SchemaInt64.Format {
		t.Errorf("Expected int64 slice schema format to be Int64, got %v. Items: %v", int64Schema.Format, int64Schema.Items)
	}

	if len(int64Schema.Example.([]int64)) != 4 {
		t.Errorf("Expected int64 slice schema example to be [1, 2, 3, 4], got %v", int64Schema.Example)
	}
}

func TestStructSlice(t *testing.T) {
	type TestStruct struct {
		Name  string
		Value int
	}

	structSlice := []TestStruct{
		{Name: "first", Value: 1},
		{Name: "second", Value: 2},
	}
	structSchema := goapidoc.SchemaFrom(structSlice)
	if structSchema.Type != goapidoc.Array {
		t.Errorf("Expected struct slice schema type to be Array, got %v", structSchema.Type)
	}
	if structSchema.Items.Type != goapidoc.Object {
		t.Errorf("Expected struct slice items type to be Object, got %v", structSchema.Items.Type)
	}
	if len(structSchema.Example.([]TestStruct)) != 2 {
		t.Error("Expected struct slice schema example to have 2 items")
	}
}

func TestStructWithSliceField(t *testing.T) {
	testStruct := struct {
		Name  string
		Value []int64
	}{
		Name:  "test",
		Value: []int64{1, 2, 3, 4},
	}
	property, ok := goapidoc.SchemaFrom(testStruct).Properties["value"]
	if !ok {
		t.Errorf("Expected value property, got %v", property)
		return
	}
	sliceSchema, ok := property.(goapidoc.Schema)
	if !ok {
		t.Error("expected value property to be a schema")
		return
	}
	sliceFormat := sliceSchema.Items.Format
	if sliceFormat != goapidoc.SchemaInt64.Format {
		t.Errorf("Expected int64 slice schema format to be Int64, got %v", sliceFormat)
	}
}
