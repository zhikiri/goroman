package goroman

var romans = map[Roman]Arabic{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var arabic = func() map[Arabic]Roman {
	res := make(map[Arabic]Roman)
	for rom, ara := range romans {
		res[ara] = rom
	}
	return res
}()

func getArabic(ch byte) Arabic {
	return romans[Roman(ch)]
}

func getRoman(num int) Roman {
	return arabic[Arabic(num)]
}
