package q

import "nippo/view/viewprops"

type QUsers struct{}

func (q QUsers) Query(props *viewprops.Props) ([]byte, error) {
	panic("not implemented")
}
