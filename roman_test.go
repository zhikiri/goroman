package goroman

import (
	"testing"
)

func TestArabicToRoman(t *testing.T) {
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
	}

	var rom Roman
	for exp, ara := range assertions {
		rom = ara.toRoman()
		if exp != rom {
			t.Errorf("Assertion is wrong, %s is not equal to %s", rom, exp)
		}
	}
	/*
		ara := Arabic(249)
		rom := ara.toRoman()
		fmt.Printf("\n\n%s\n", rom)
	*/
}
