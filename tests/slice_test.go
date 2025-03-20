package tests

import (
	"testing"

	"github.com/CyberTea0X/goapidoc"
)

func TestSlice(t *testing.T) {
	// Test empty slice
	emptySlice := []int{}
	emptySchema := goapidoc.SchemaFrom(emptySlice)
	if emptySchema.Type != goapidoc.Array {
		t.Errorf("Expected empty slice schema type to be Array, got %v", emptySchema.Type)
	}

	// Test string slice
	strSlice := []string{"one", "two", "three"}
	strSchema := goapidoc.SchemaFrom(strSlice)
	if strSchema.Type != goapidoc.Array {
		t.Errorf("Expected string slice schema type to be Array, got %v", strSchema.Type)
	}

	if len(strSchema.Example.([]string)) != 3 {
		t.Errorf("Expected string slice schema example to be ['one', 'two', 'three'], got %v", strSchema.Example)
	}

	// Test int slice
	intSlice := []int{1, 2, 3, 4}
	intSchema := goapidoc.SchemaFrom(intSlice)
	if intSchema.Type != goapidoc.Array {
		t.Errorf("Expected int slice schema type to be Array, got %v", intSchema.Type)
	}

	// Test int64 slice
	int64Slice := []int64{1, 2, 3, 4}
	int64Schema := goapidoc.SchemaFrom(int64Slice)
	if int64Schema.Type != goapidoc.Array {
		t.Errorf("Expected int64 slice schema type to be Array, got %v", int64Schema.Type)
	}

	if int64Schema.Items.Format != goapidoc.SchemaInt64.Format {
		t.Errorf("Expected int64 slice schema format to be Int64, got %v", int64Schema.Format)
	}

	if len(int64Schema.Example.([]int64)) != 4 {
		t.Errorf("Expected int64 slice schema example to be [1, 2, 3, 4], got %v", int64Schema.Example)
	}
}
