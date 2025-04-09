package q

import (
	"errors"
	"nippo/view/viewprops"
)

type Q interface {
	Query(props *viewprops.Props) ([]byte, error)
}

func FromParams(params map[string]string) (Q, error) {
	q, ok := params["q"]
	if !ok {
		return nil, errors.New("q not found")
	}

	switch q {
	case "users":
		panic("not implemented")
	case "reports":
		panic("not implemented")
	case "print":
		panic("not implemented")
	case "calls":
		panic("not implemented")
	case "download":
		panic("not implemented")
	default:
		return nil, errors.New("unexpected q string")
	}
}
