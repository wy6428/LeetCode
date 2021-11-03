package main

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]int, len(s))
	var ans int
	head := 0
	for head < len(s) {
		var subStrLen int
		tail, nextHead := head, head+1
		var dupChar byte
		for ; tail < len(s); tail++ {
			dupChar = s[tail]
			if v, ok := charSet[dupChar]; ok && v >= head && head != tail {
				nextHead = v + 1 //下次直接从重复的第一个字符的下一个开始遍历, 不可以直接从第二个开始：dvdf
				break
			}
		}
		charSet[dupChar] = tail
		subStrLen = tail - head
		ans = max(subStrLen, ans)
		head = nextHead
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	s := "anviaj"
	println(lengthOfLongestSubstring(s))

}
