package crypt

import (
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		c        string
		wantText string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotText := Decode(tt.c); gotText != tt.wantText {
				t.Errorf("Decode() = %v, want %v", gotText, tt.wantText)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		wantC string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Encode(tt.text); gotC != tt.wantC {
				t.Errorf("Encode() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
