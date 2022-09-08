package cm

import (
	"fmt"
)

var _ fmt.Stringer = (*ID)(nil)

type ID string

func (id ID) String() string {
	return string(id)
}

const EmptyID = ID("")
