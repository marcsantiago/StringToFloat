package stringtofloat

import (
	"strconv"
	"strings"
)

func normalizeEurope(old string) string {
	count := strings.Count(old, ".")
	s := strings.Replace(old, ",", ".", -1)
	return strings.Replace(s, ".", "", count)

}
func normalizeAmericanBritain(old string) string {
	return strings.Replace(old, ",", "", -1)
}

func reverse(a []byte) string {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return string(a)
}

// Convert determines locale there are really only two types, US/Britain and rest of Europe
func Convert(fs string) (f float64, err error) {
	// if a decimal point comes before the comma then its US/Britain
	rev := reverse([]byte(fs))
	point := strings.Index(rev, ".")
	comma := strings.Index(rev, ",")

	// no commas so need to normalize
	if comma == -1 {
		f, err = strconv.ParseFloat(fs, 64)
		if err != nil {
			return
		}
	}

	if point < comma {
		f, err = strconv.ParseFloat(normalizeAmericanBritain(fs), 64)
		if err != nil {
			return
		}
	} else {
		f, err = strconv.ParseFloat(normalizeEurope(fs), 64)
		if err != nil {
			return
		}
	}
	return
}
