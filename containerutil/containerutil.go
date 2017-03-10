package containerutil

import (
	"gotools/reflectutil"
)

func ConcatWithPostfix(a []string, postfix string) string {
	res := ""
	for _, val := range a {
		res += val + postfix
	}
	return res
}

func Concat(a []string) string {
	return ConcatWithPostfix(a, "")
}

// func toArray(input map[interface{}]interface{}) []interface{} {
// 	if len(input) == 0 {
// 		return []interface{}{}
// 	}

// 	for k, v := range input {

// 	}
// }
