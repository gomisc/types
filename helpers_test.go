package types_test

import (
	"testing"
	"unsafe"

	"git.corout.in/golibs/types"
)

func TestIsPrimitiveValues(t *testing.T) {
	type args struct {
		values []interface{}
	}
	type SomeType struct {
		Name string
	}

	type AliasString = string
	type SomeString string

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "int",
			args: args{values: []interface{}{
				1,
				int8(1),
				int16(1),
				int32(1),
				int64(1),
				uint(1),
				uint8(1),
				uint16(1),
				uint32(1),
				uint64(1),
			}},
			want: true,
		},
		{
			name: "float",
			args: args{values: []interface{}{float32(1), float64(1)}},
			want: true,
		},
		{
			name: "string",
			args: args{values: []interface{}{"Hello", ",", "World", "!"}},
			want: true,
		},
		{
			name: "bool",
			args: args{values: []interface{}{false, true}},
			want: true,
		},
		{
			name: "uintptr",
			args: args{values: []interface{}{uintptr(unsafe.Pointer(&SomeType{}))}},
			want: true,
		},
		{
			name: "new types",
			args: args{values: []interface{}{AliasString("alias string"), SomeString("some string type")}},
			want: true,
		},
		{
			name: "not primitive",
			args: args{values: []interface{}{SomeType{Name: ""}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := types.IsPrimitiveValues(tt.args.values...); got != tt.want {
				t.Errorf("IsPrimitiveValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPrimitive(b *testing.B) {
	elements := []interface{}{1, float32(32), "", []byte("some text"), true}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = types.IsPrimitiveValues(elements)
	}
}
