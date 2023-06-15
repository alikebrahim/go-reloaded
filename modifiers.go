package goreloaded

// ASCII:A==65-Z==90 || a==97-z==122
func Cap(b []byte) []byte {
	var cappedB []byte
	for i, item := range b {
		if i == 0 && item > 90 {
			cappedB = append(cappedB, item-32)
		} else {

			cappedB = append(cappedB, item)
		}
	}
	return cappedB
}

func Up(b []byte) []byte {
	var upperB []byte
	for _, item := range b {
		if item == byte(' ') {
			upperB = append(upperB, item)
		} else {
			upperB = append(upperB, item-32)
		}
	}
	return upperB
}

func Low(b []byte) []byte {
	var lowerB []byte
	for _, item := range b {
		if item == byte(' ') {
			lowerB = append(lowerB, item)
		} else {
			lowerB = append(lowerB, item+32)
		}
	}
	return lowerB
}

// func TrimSpace(b []byte) []byte {
// 	var trimmedB []byte
// 	for _, item := range b {
// 		if item == byte(' ') {
// 			continue
// 		} else {
// 			trimmedB = append(trimmedB, item)
// 		}
// 	}
// 	return trimmedB
// }
