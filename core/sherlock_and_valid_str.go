package core

// https://www.hackerrank.com/challenges/one-month-preparation-kit-sherlock-and-valid-string/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-three
func SherlockStr(s string) string {
	ocr := make(map[string]int)
	ocr["high"] = 0
	ocr["low"] = 0
	for i := 0; i < len(s); i++ {
		ocr[string(s[i])] += 1
		if ocr[string(s[i])] > ocr["high"] {
			ocr["high"] = ocr[string(s[i])]
		}
		if ocr[string(s[i])] < ocr["high"] || ocr[string(s[i])] == 1 {
			ocr["low"] = ocr[string(s[i])]
		}
	}

	var counterH, counterL int
	for key := range ocr {
		if key == "low" || key == "high" {
			continue
		}
		if ocr[key] == ocr["high"] && ocr[key]-1 == ocr["low"] {
			counterH++
			continue
		}
		if ocr[key] == ocr["low"] {
			counterL++
			continue
		}
		if ocr[key] == 1 {
			counterL++
			continue
		}
		return "NO"
	}
	if counterL == 1 || counterL == 0 {
		return "YES"
	}
	if counterH > 1 {
		return "NO"
	}

	return "YES"
}
