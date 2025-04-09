package q

import (
	"errors"
	"nippo/input"
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
		user, ok := params["user"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		userIdStr := input.Clean(user)

		return QCalls{userIdStr, year, month}, nil
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
		user, ok := params["user"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		userIdStr := input.Clean(user)

		return QReports{userIdStr, year, month}, nil
	case "print":
		report, ok := params["report"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		repoIdStr := input.Clean(report)

		return QPrint{repoIdStr}, nil
	case "download":
		report, ok := params["report"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		repoIdStr := input.Clean(report)

		return QDownload{repoIdStr}, nil
	default:
		return nil, errors.New("unexpected q string")
	}
}
