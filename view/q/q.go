package q

import (
	"errors"
	"nippo/timeconv"
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
		return QUsers{}, nil
	case "calls":
		y, ok := params["y"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		m, ok := params["m"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		year, month, err := timeconv.AToYM(y, m)
		if err != nil {
			return nil, err
		}

		return QCalls{year, month}, nil
	case "reports":
		y, ok := params["y"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		m, ok := params["m"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		year, month, err := timeconv.AToYM(y, m)
		if err != nil {
			return nil, err
		}

		return QReports{year, month}, nil
	case "print":
		panic("not implemented")
	case "download":
		panic("not implemented")
	default:
		return nil, errors.New("unexpected q string")
	}
}
