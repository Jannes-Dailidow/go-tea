package util

import (
	"reflect"
	"testing"

	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
)

func TestTryFocusTeaModel(t *testing.T) {
	type args struct {
		m            tea.Model
		parentKeyMap help.KeyMap
	}
	tests := []struct {
		name  string
		args  args
		want  tea.Cmd
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := TryFocusTeaModel(tt.args.m, tt.args.parentKeyMap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TryFocusTeaModel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TryFocusTeaModel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTryBlurTeaModel(t *testing.T) {
	type args struct {
		m tea.Model
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TryBlurTeaModel(tt.args.m); got != tt.want {
				t.Errorf("TryBlurTeaModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
