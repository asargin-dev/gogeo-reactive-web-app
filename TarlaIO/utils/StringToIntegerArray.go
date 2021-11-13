package utils

import (
	"strconv"
	"strings"
)

// We are collecting string that is selected in table
// So it should convert to string to int array
func StringToIntArray(s string) []int {
	strs := strings.Split(s, ",")
	res := make([]int, len(strs))
	for i := range res {
		res[i], _ = strconv.Atoi(strs[i])
	}
	return res
}
