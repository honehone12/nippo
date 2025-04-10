package q

import (
	"nippo/view/viewprops"

	"github.com/google/uuid"
)

type QDownload struct {
	repoId uuid.UUID
}

func (q QDownload) Query(props *viewprops.Props) ([]byte, error) {
	panic("not implemented")
}
