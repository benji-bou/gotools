package tools

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
