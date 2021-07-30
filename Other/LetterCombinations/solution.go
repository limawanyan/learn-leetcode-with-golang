package LetterCombinations

import "strconv"

var dict = map[int][]string{
	1:[]string{"a","b","c"},
	2:[]string{"d","e"},
	3:[]string{"f","g","h"},
}

func LetterCombinations(result *[]string,temp,nums string)  {
	if len(nums) == 0 {
		*result = append(*result,temp)
		return
	}
	num,_ := strconv.Atoi(nums[:1])
	for _,v := range dict[num]{
		LetterCombinations(result,temp+v,nums[1:])
	}
}