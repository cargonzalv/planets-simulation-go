package api

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestSimulacionHandler(t *testing.T) {
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
			if err := SimulacionHandler(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SimulacionHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
