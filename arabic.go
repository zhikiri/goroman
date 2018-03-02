package goroman

import (
	"bytes"
)

type Arabic int

func (a Arabic) getReducer() Arabic {
	if a == 4 || a == 9 {
		return 1
	}
	var reducer Arabic
	for ara := range arabic {
		if a >= ara && reducer < ara {
			reducer = ara
		}
	}
	return reducer
}

func (a Arabic) getSubstructionValues(red, lim int) (Roman, Roman) {
	return getRoman(1 * int(red)), getRoman(lim * int(red))
}

func (a Arabic) applySubstration(red *Arabic, buf *bytes.Buffer) bool {
	var first, second Roman
	if a == 4 || a / *red == 4 {
		first, second = a.getSubstructionValues(int(*red), 5)
		*red *= 4
	} else if a == 9 || a / *red == 9 {
		first, second = a.getSubstructionValues(int(*red), 10)
		*red *= 9
	}
	if first != "" && second != "" {
		buf.WriteString(string(first))
		buf.WriteString(string(second))
		return true
	}
	return false
}

func (a Arabic) toRoman() Roman {
	var buffer bytes.Buffer
	reducer := a.getReducer()
	for a > 0 {
		if !a.applySubstration(&reducer, &buffer) {
			buffer.WriteString(string(getRoman(int(reducer))))
		}
		a -= reducer
		reducer = a.getReducer()
	}
	return Roman(buffer.String())
}
