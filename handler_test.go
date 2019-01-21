package schemaorg

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOgHandler(t *testing.T) {
	type args struct {
		tpl  string
		data fmt.Stringer
	}
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want http.HandlerFunc
	// }{
	// 	{"valid" }
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := OgHandler(tt.args.tpl, tt.args.data); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("OgHandler() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestOgSchemaHandler(t *testing.T) {
	type args struct {
		tpl     string
		resolve Resolver
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OgSchemaHandler(tt.args.tpl, tt.args.resolve); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OgSchemaHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
