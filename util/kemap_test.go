package util

import (
	"reflect"
	"testing"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func TestAnnounceKeyMapCmd(t *testing.T) {
	type args struct {
		keyMaps []help.KeyMap
	}
	tests := []struct {
		name string
		args args
		want tea.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnnounceKeyMapCmd(tt.args.keyMaps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnnounceKeyMapCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeKeyMaps(t *testing.T) {
	type args struct {
		keyMaps []help.KeyMap
	}
	tests := []struct {
		name string
		args args
		want help.KeyMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeKeyMaps(tt.args.keyMaps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeKeyMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergedKeyMaps_ShortHelp(t *testing.T) {
	tests := []struct {
		name string
		m    MergedKeyMaps
		want []key.Binding
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ShortHelp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergedKeyMaps.ShortHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergedKeyMaps_FullHelp(t *testing.T) {
	tests := []struct {
		name string
		m    MergedKeyMaps
		want [][]key.Binding
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.FullHelp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergedKeyMaps.FullHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyBindingList_ShortHelp(t *testing.T) {
	tests := []struct {
		name string
		km   KeyBindingList
		want []key.Binding
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.km.ShortHelp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyBindingList.ShortHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyBindingList_FullHelp(t *testing.T) {
	tests := []struct {
		name string
		km   KeyBindingList
		want [][]key.Binding
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.km.FullHelp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyBindingList.FullHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}
