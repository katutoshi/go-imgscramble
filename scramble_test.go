package imgscramble

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name       string
		givenSlice []uint8
		givenIndex []uint64
		wantSlice  []uint8
	}{
		{"normal",
			[]uint8{10, 20, 30, 40, 50, 60, 70, 80},
			[]uint64{0, 1, 2, 3, 4, 5, 6, 7},
			[]uint8{10, 20, 30, 40, 50, 60, 70, 80},
		},
		{"reversed",
			[]uint8{10, 20, 30, 40, 50, 60, 70, 80},
			[]uint64{7, 6, 5, 4, 3, 2, 1, 0},
			[]uint8{80, 70, 60, 50, 40, 30, 20, 10},
		},
		{"random",
			[]uint8{10, 20, 30, 40, 50, 60, 70, 80},
			[]uint64{2, 4, 5, 0, 3, 7, 1, 6},
			[]uint8{30, 50, 60, 10, 40, 80, 20, 70},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Shuffle(tt.givenSlice, tt.givenIndex)
			if cmp.Diff(tt.givenSlice, tt.wantSlice) != "" {
				t.Errorf("Shuffle() = %v, want %v", tt.givenSlice, tt.wantSlice)
			}
		})
	}
}
