package table2

import (
	"fmt"
	"strconv"
)

// Body ...
type Body []interface{}

// NewBody ...
func NewBody(values ...interface{}) Body {
	return append(make(Body, 0, len(values)), values...)
}

// Copy ...
func (b Body) Copy() Body {
	return append(make(Body, 0, len(b)), b...)
}

// Equal ...
func (b Body) Equal(bdy Body) bool {
	if len(b) != len(bdy) {
		return false
	}

	var i int
	for ; i < len(b) && b[i] == bdy[i]; i++ {
	}

	return i == len(b)
}

// Strings ...
func (b Body) Strings() []string {
	ss := make([]string, 0, len(b))
	for i := 0; i < len(b); i++ {
		switch Fmt(b[i]) {
		case Flt:
			ss = append(ss, strconv.FormatFloat(b[i].(float64), 'f', -1, 64))
		case Int:
			ss = append(ss, strconv.Itoa(b[i].(int)))
		case Str:
			ss = append(ss, b[i].(string))
		default:
			// TODO: Should this panic?
			ss = append(ss, fmt.Sprintf("%v", b[i]))
		}
	}

	return ss
}
