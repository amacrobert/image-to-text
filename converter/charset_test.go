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
			cs := NewSimpleCharset(tt.invert)
			got := cs.CharFor(tt.grayValue)
			if got != tt.want {
				t.Errorf("CharFor(%d) = %c, want %c", tt.grayValue, got, tt.want)
			}
		})
	}
}
