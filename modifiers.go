package goreloaded

import "fmt"

var hexMap = map[byte]int{
	byte('A'): 10,
	byte('B'): 11,
	byte('C'): 12,
	byte('D'): 13,
	byte('E'): 14,
	byte('F'): 15,
	byte('a'): 10,
	byte('b'): 11,
	byte('c'): 12,
	byte('d'): 13,
	byte('e'): 14,
	byte('f'): 15,
}

func Cap(b []byte) []byte {
	var cappedB []byte
	for i, item := range b {
		if i == 0 && (item >= 97 && item <= 122) {
			cappedB = append(cappedB, item-32)
		} else if i > 0 && (item >= 65 && item <= 90) {

			cappedB = append(cappedB, item+32)
		} else {
			cappedB = append(cappedB, item)

		}
	}
	return cappedB
}

func Up(b []byte) []byte {
	var upperB []byte
	for _, item := range b {
		if item >= 97 && item <= 122 {
			upperB = append(upperB, item-32)
		} else {
			upperB = append(upperB, item)
		}
	}
	return upperB
}

func Low(b []byte) []byte {
	var lowerB []byte
	for _, item := range b {
		if item >= 65 && item <= 90 {
			lowerB = append(lowerB, item+32)
		} else {
			lowerB = append(lowerB, item)
		}
	}
	return lowerB
}

func Hex(b []byte) []byte {
	var sum int
	var sumB []byte
	hexLen := len(b)
	power := 0
	if b[0] == byte(' ') {
		return nil
	}
	for i := hexLen - 1; i >= 0; i-- {
		if _, ok := hexMap[b[i]]; ok {
			hexDigit := hexMap[b[i]]
			sum += hexDigit * Power(16, power)
		} else {
			sum += int(b[i]-48) * Power(16, power)
		}
		power++
	}
	sumStr := fmt.Sprintf("%d", sum)
	for i := 0; i < len(sumStr); i++ {
		sumB = append(sumB, byte(sumStr[i]))
	}
	sumB = append(sumB, byte(' '))
	return sumB
}

func Bin(b []byte) []byte {
	var sum int
	var sumB []byte
	binLen := len(b)
	power := 0
	if b[0] == byte(' ') {
		return nil
	}
	for i := binLen - 1; i >= 0; i-- {
		sum += int(b[i]-48) * Power(2, power)
		power++
	}
	sumStr := fmt.Sprintf("%d", sum)
	for i := 0; i < len(sumStr); i++ {
		sumB = append(sumB, byte(sumStr[i]))
	}
	sumB = append(sumB, byte(' '))
	return sumB
}
