package goroman

type Roman string

func (r Roman) isException(index int) bool {
	if index >= len(r) || index+1 >= len(r) {
		return false
	}
	return getArabic(r[index]) < getArabic(r[index+1])
}

func (r Roman) toArabic() Arabic {
	var res Arabic
	for i := 0; i < len(r); i++ {
		if r.isException(i) {
			res += getArabic(r[i+1]) - getArabic(r[i])
			i++
			continue
		}
		res += getArabic(r[i])
	}
	return res
}
