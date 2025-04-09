package q

import (
	"nippo/schema"
	"nippo/view/viewprops"
)

type QDownload struct {
	repoIdStr schema.CleanText
}

func (q QDownload) Query(props *viewprops.Props) ([]byte, error) {
	panic("not implemented")
}
