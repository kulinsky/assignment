package assignment

import (
	"errors"
	"time"
)

type Assignment struct {
	ID        ID
	UserID    UserID
	CourseID  CourseID
	DateStart *DateStart
	DateEnd   *DateEnd
	Status    Status
}

func NewAssignment(
	id ID,
	userID UserID,
	courseID CourseID,
) *Assignment {
	return &Assignment{
		ID:       id,
		UserID:   userID,
		CourseID: courseID,
		Status:   StatusNew,
	}
}

var (
	ErrAssignmentNotReadyToStart  = errors.New("assignment is not ready to start")
	ErrAssignmentNotReadyToFinish = errors.New("assignment is not ready to finish")
	ErrAssignmentNotFound         = errors.New("assignment not found")
)

func (a *Assignment) Start(now time.Time) error {
	if a.Status != StatusNew {
		return ErrAssignmentNotReadyToStart
	}

	dt := DateStart(now)

	a.Status = StatusInProgress
	a.DateStart = &dt

	return nil
}

func (a *Assignment) Finish(now time.Time) error {
	if a.Status != StatusInProgress {
		return ErrAssignmentNotReadyToFinish
	}

	dt := DateEnd(now)

	a.Status = StatusFinished
	a.DateEnd = &dt

	return nil
}
