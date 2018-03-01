package goroman

import (
	"bytes"
)

type Arabic int

func (a Arabic) getReducer() Arabic {
	var reducer Arabic
	for ara := range arabic {
		if a >= ara && reducer < ara {
			reducer = ara
		}
	}
	return reducer
}

func (a Arabic) toRoman() Roman {
	var buffer bytes.Buffer
	reducer := a.getReducer()
	for a > 0 {
		if a == 4 || a == 9 {
			reducer = 1
		}
		if a == 4 || a/reducer == 4 {
			buffer.WriteString(string(getRoman(1 * int(reducer))))
			buffer.WriteString(string(getRoman(5 * int(reducer))))
			reducer = 4 * reducer
		} else if a == 9 || a/reducer == 9 {
			buffer.WriteString(string(getRoman(1 * int(reducer))))
			buffer.WriteString(string(getRoman(10 * int(reducer))))
			reducer = 9 * reducer
		} else {
			buffer.WriteString(string(getRoman(int(reducer))))
		}
		a -= reducer
		reducer = a.getReducer()
	}
	return Roman(buffer.String())
}
