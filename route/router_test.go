package route

import (
	_ "prueba-meli/docs"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootHandler(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rootHandler(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("rootHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
