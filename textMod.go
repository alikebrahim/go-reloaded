package goreloaded

import (
	"bytes"
)

func TextMod(t *[][]byte, l *Lexer) {
	for i, item := range l.Tokens {
		if item == 0 {
			NumOfIdens := ModAnalyzer(l.TokenVals[i])
			if bytes.Contains(l.TokenVals[i], []byte("cap")) {
				for j := i - NumOfIdens; j < i; j++ {
					*t = append((*t)[:j], Cap(l.TokenVals[j]))
				}
			} else if bytes.Contains(l.TokenVals[i], []byte("up")) {
				for j := i - NumOfIdens; j < i; j++ {
					*t = append((*t)[:j], Up(l.TokenVals[j]))
				}
			} else if bytes.Contains(l.TokenVals[i], []byte("low")) {
				for j := i - NumOfIdens; j < i; j++ {
					*t = append((*t)[:j], Low(l.TokenVals[j]))
				}
			} //else if strings.Contains(l.TokenVals[i], "hex") {
			// 	for j := i - NumOfIdens; j < i; j++ {
			// 		bs, _ := hex.DecodeString(string(l.TokenVals[j]))
			// 		for _, item := range bs {
			// 			number := fmt.Sprintf("%d", item)
			// 			l.TokenVals[j] = []byte(number)
			// 		}
			// 	}
			// } else if strings.Contains(l.TokenVals[i], "bin") {
			// 	for j := i - NumOfIdens; j < i; j++ {
			// 		if string(l.TokenVals[j]) != " " {
			// 			decimal, _ := strconv.ParseUint(string(l.TokenVals[j]), 2, 64)
			// 			dec := fmt.Sprintf("%d", decimal)
			// 			l.TokenVals[j] = []byte(dec)
			// 		} else {
			// 			continue
			// 		}
			// 	}
			// }
		}
		*t = append(*t, l.TokenVals[i])
	}
}
