package nonrepeatingsubstr

func lengthOfNonRepeatingSubStr(s string) int {
	//lastOccurred := make(map[rune]int)
	// 用空间换时间
	// 中文的大小 0xffff 大概够了
	//rune 是 int32 其实是

	lastOccurred := make([]int, 0xffff)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	//lastOccurred[0x65] = 1
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
