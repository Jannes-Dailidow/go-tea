package util

import (
	"reflect"
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestSize_UpdateFromMsg(t *testing.T) {
	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name string
		s    *Size
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdateFromMsg(tt.args.msg); got != tt.want {
				t.Errorf("Size.UpdateFromMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize_ToMsg(t *testing.T) {
	tests := []struct {
		name string
		s    Size
		want tea.WindowSizeMsg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToMsg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Size.ToMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSizeFromView(t *testing.T) {
	type args struct {
		view string
	}
	tests := []struct {
		name string
		args args
		want Size
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SizeFromView(tt.args.view); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SizeFromView() = %v, want %v", got, tt.want)
			}
		})
	}
}
