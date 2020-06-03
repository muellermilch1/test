package main

import "fmt"

func romanToInt(s string) int {
	dict := map[rune]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
	var sum int
	for i,v:=range s{
			subt := false
			for j:=i;j<len(s)-1&&!subt;j++{
					if dict[rune(s[j+1])]>dict[v]{
							subt =true
					}
			} 
			if subt {
					sum -= dict[v]            
			}else{
					sum += dict[v]            
			}
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}