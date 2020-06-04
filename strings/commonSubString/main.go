package main

import "fmt"

//LongestCommonSubsequence returns the longes subsequence found in
//s1/s2
func LongestCommonSubsequence(s1 string, s2 string) string {
	dp := make([][]string, len(s1)+1)
	for i := range dp {
		dp[i] = make([]string, len(s2)+1)
	}
	s1, s2 = " "+s1, " "+s2
	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i-1][j-1] + string(s2[j])
			} else {
				if len(dp[i-1][j]) > len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
		fmt.Println(dp[i])
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func main() {
	fmt.Println(LongestCommonSubsequence("ZXVVYZW", "XLYLZPW"))
}
