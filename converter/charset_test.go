package converter

import "testing"

func TestSimpleCharset_CharFor(t *testing.T) {
	tests := []struct {
		name      string
		grayValue uint8
		invert    bool
		want      rune
	}{
		{"black_normal", 0, false, '$'},
		{"white_normal", 255, false, ' '},
		{"black_inverted", 0, true, ' '},
		{"white_inverted", 255, true, '$'},
		{"mid_gray_normal", 128, false, 'n'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := NewSimpleCharset(tt.invert, 0, 255)
			got := cs.CharFor(tt.grayValue)
			if got != tt.want {
				t.Errorf("CharFor(%d) = %c, want %c", tt.grayValue, got, tt.want)
			}
		})
	}
}

func TestSimpleCharset_BlackpointWhitepoint(t *testing.T) {
	tests := []struct {
		name       string
		grayValue  uint8
		blackpoint uint8
		whitepoint uint8
		wantBlack  bool // true if should map to darkest char
		wantWhite  bool // true if should map to lightest char
	}{
		{"below_blackpoint", 30, 50, 200, true, false},
		{"at_blackpoint", 50, 50, 200, true, false},
		{"above_whitepoint", 220, 50, 200, false, true},
		{"at_whitepoint", 200, 50, 200, false, true},
		{"midpoint", 125, 50, 200, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := NewSimpleCharset(false, tt.blackpoint, tt.whitepoint)
			got := cs.CharFor(tt.grayValue)
			if tt.wantBlack && got != '$' {
				t.Errorf("CharFor(%d) = %c, want '$' (black)", tt.grayValue, got)
			}
			if tt.wantWhite && got != ' ' {
				t.Errorf("CharFor(%d) = %c, want ' ' (white)", tt.grayValue, got)
			}
		})
	}
}
