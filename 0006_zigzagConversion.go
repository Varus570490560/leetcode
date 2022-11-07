package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var strArrays []string
	for i := 0; i < numRows; i++ {
		strArrays = append(strArrays, "")
	}
	var isDown bool = true
	curStrIndex := 0
	for i := 0; i < len(s); i++ {
		strArrays[curStrIndex] += string(s[i])
		if isDown && curStrIndex == numRows-1 {
			isDown = !isDown
			curStrIndex--
			continue
		}
		if !isDown && curStrIndex == 0 {
			isDown = !isDown
			curStrIndex++
			continue
		}
		if isDown {
			curStrIndex++
		} else {
			curStrIndex--
		}
	}
	res := ""
	for i := 0; i < len(strArrays); i++ {
		res += strArrays[i]
	}
	return res

}
