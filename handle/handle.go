package handle

import "strconv"

func GetNumber(num int) string {
	resStr := ""
	if num%3 == 0 {
		resStr = resStr + "FIZZ"
	}
	if num%5 == 0 {
		resStr = resStr + "BUZZ"
	}
	if resStr == "" {
		resStr = strconv.Itoa(num)
	}
	return resStr
}
