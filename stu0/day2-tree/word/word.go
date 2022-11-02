package main

import "fmt"

//dp[i] 表示字符串i是否满足拆分条件
//若dp[i] = true。则拆分为两个字符串 s[0:j] s[j:i]
//s[0:j]要可拆分，即s[j]=true
//dp[i] = dp[j] && s[j:i] in wordDict
func wordBreak(s string, wordDict []string) {
	l := len(s)
	dict := make(map[string]bool)
	dp := make([]bool, l+1)
	dp[0] = true
	for _, word := range wordDict {
		dict[word] = true
	}
	//	dict[""] = true
	for i := 0; i <= l; i++ {
		for j := 0; j <= i; j++ {
			if ok := dict[s[j:i]]; ok && dp[j] {
				fmt.Printf("%v %v\n", s[0:j], s[j:i])
				dp[i] = true
			}
		}
	}
}
func main() {
	s := "helloworldchina"
	wordDict := []string{"hello", "world", "china"}
	wordBreak(s, wordDict)
}
