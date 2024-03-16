package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`^(?P<type>@|#)(?P<field>\w+):?(?P<cond>[<>]=?|!=)?(?P<val>\w+)?$`)

const (
	typeIndex  = 1
	fieldIndex = 2
	condIndex  = 3
	valIndex   = 4
)

type Cond struct {
	Type      string
	Field     string
	Condition string
	Value     any
}

func Parser(condStr string) (*Cond, error) {

	matches := re.FindStringSubmatch(condStr)
	fmt.Println(matches)
	return &Cond{
		Type:  matches[typeIndex],
		Field: matches[fieldIndex],
		Condition: func() string {
			if matches[condIndex] == "" {
				return "="
			}
			return matches[condIndex]

		}(),
		Value: func() any {
			if matches[valIndex] == "" {
				return nil
			}
			return matches[valIndex]
		}(),
	}, nil
}

func main() {

	fmt.Println(Parser("@office"))

}
