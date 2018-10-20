package util

import (
	"fmt"
	"strings"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IntInSlice(needle int, haystack []int) (int, bool) {
	for index, straw := range haystack {
		if straw == needle {
			return index, true
		}
	}

	return 0, false
}

func DeleteFromIntSlice(s []int, index int) []int {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func IntSliceJoin(s []int, delim string) string {
	// https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string/37533144

	return strings.Trim(strings.Replace(fmt.Sprint(s), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
