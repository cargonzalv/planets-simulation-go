package db

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}

func TestRetryHandler(t *testing.T) {
	type args struct {
		n int
		f func() (bool, error)
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
			if err := RetryHandler(tt.args.n, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("RetryHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
