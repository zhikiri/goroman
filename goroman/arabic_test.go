package goroman

import (
	"testing"
)

func TestRomanToArabic(t *testing.T) {
	assertions := map[Roman]Arabic{
		"I":      1,
		"II":     2,
		"III":    3,
		"IV":     4,
		"V":      5,
		"VI":     6,
		"VII":    7,
		"VIII":   8,
		"IX":     9,
		"X":      10,
		"XXXIX":  39,
		"CCXLVI": 246,
		"CCVII":  207,
		"MLXVI":  1066,
		"MMXIV":  2014,
		"CDIX":   409,
	}

	var ara Arabic
	for rom, exp := range assertions {
		ara = rom.toArabic()
		if ara != exp {
			t.Errorf("Assertion is wrong, %d is not equal to %d", ara, exp)
		}
	}
}
