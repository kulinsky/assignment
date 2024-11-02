package assignment

import "errors"

type Status int

const (
	StatusNew Status = iota + 1
	StatusInProgress
	StatusFinished
)

var ErrInvalidStatus = errors.New("invalid status")

func NewStatusFromInt(s int) (Status, error) {
	switch s {
	case 1:
		return StatusNew, nil
	case 2:
		return StatusInProgress, nil
	case 3:
		return StatusFinished, nil
	default:
		return 0, ErrInvalidStatus
	}
}
