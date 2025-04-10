package q

import (
	"errors"
	"nippo/input"
	"nippo/timeconv"
	"nippo/view/viewprops"

	"github.com/google/uuid"
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
		userId, err := uuid.Parse(string(userIdStr))
		if err != nil {
			return nil, err
		}

		return QCalls{userId, year, month}, nil
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
		userId, err := uuid.Parse(string(userIdStr))
		if err != nil {
			return nil, err
		}

		return QReports{userId, year, month}, nil
	case "print":
		report, ok := params["report"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		repoIdStr := input.Clean(report)
		repoId, err := uuid.Parse(string(repoIdStr))
		if err != nil {
			return nil, err
		}

		return QPrint{repoId}, nil
	case "download":
		report, ok := params["report"]
		if !ok {
			return nil, errors.New("unexpected params")
		}
		repoIdStr := input.Clean(report)
		repoId, err := uuid.Parse(string(repoIdStr))
		if err != nil {
			return nil, err
		}

		return QDownload{repoId}, nil
	default:
		return nil, errors.New("unexpected q string")
	}
}
