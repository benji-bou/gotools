package tools

import (
	"regexp"
)

type RegexpTools struct {
	regexp.Regexp
}

func (rp *RegexpTools) ArrayMatch(input []string) []string {
	output := make([]string, 0)
	for _, value := range input {
		if rp.MatchString(value) == true {
			output = append(output, value)
		}
	}
	return output
}

func Compile(expr string) (*RegexpTools, error) {
	res, err := regexp.Compile(expr)
	return &RegexpTools{*res}, err
}

func CompilePOSIX(expr string) (*RegexpTools, error) {
	res, err := regexp.CompilePOSIX(expr)
	return &RegexpTools{*res}, err
}

func MustCompile(str string) *RegexpTools {
	res := regexp.MustCompile(str)
	return &RegexpTools{*res}
}
func MustCompilePOSIX(str string) *RegexpTools {
	res := regexp.MustCompilePOSIX(str)
	return &RegexpTools{*res}
}
