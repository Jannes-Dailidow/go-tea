package util

import (
	"reflect"
	"testing"
)

func TestClamp(t *testing.T) {
	type args struct {
		_min    T
		_wanted T
		_max    T
	}
	tests := []struct {
		name string
		args args
		want T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args._min, tt.args._wanted, tt.args._max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
