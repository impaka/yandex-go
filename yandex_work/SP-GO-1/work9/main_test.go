package main_test

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		wantLen   int
		wantError error
	}{
		{
			name:      "valid ASCII",
			input:     []byte("hello"),
			wantLen:   5,
			wantError: nil,
		},
		{
			name:      "valid UTF-8 runes",
			input:     []byte("привет"), // 6 символов
			wantLen:   6,
			wantError: nil,
		},
		{
			name:      "emoji",
			input:     []byte("👍👍"),
			wantLen:   2,
			wantError: nil,
		},
		{
			name:      "invalid UTF-8",
			input:     []byte{0xff, 0xfe, 0xfd},
			wantLen:   0,
			wantError: ErrInvalidUTF8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen, gotErr := GetUTFLength(tt.input)

			if gotLen != tt.wantLen {
				t.Errorf("expected length %d, got %d", tt.wantLen, gotLen)
			}

			if !errors.Is(gotErr, tt.wantError) {
				t.Errorf("expected error %v, got %v", tt.wantError, gotErr)
			}
		})
	}
}
