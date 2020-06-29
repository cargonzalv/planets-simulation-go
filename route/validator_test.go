package route

import (
	"reflect"
	"testing"
)

func TestNewValidator(t *testing.T) {
	tests := []struct {
		name string
		want *Validator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidator_Validate(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		v       *Validator
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Validate(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Validator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
