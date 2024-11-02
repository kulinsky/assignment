package assignment

import (
	"github.com/google/uuid"
)

type (
	ID       uuid.UUID
	CourseID uuid.UUID
	UserID   uuid.UUID
)

type UUIDGenerator func() (uuid.UUID, error)

func NewV7UUID() (uuid.UUID, error) {
	return uuid.NewV7()
}

func AssignmentIDFromString(s string) (ID, error) {
	rawid, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}

	return ID(rawid), nil
}

func CourseIDFromString(s string) (CourseID, error) {
	rawid, err := uuid.Parse(s)
	if err != nil {
		return CourseID{}, err
	}

	return CourseID(rawid), nil
}

func UserIDFromString(s string) (UserID, error) {
	rawid, err := uuid.Parse(s)
	if err != nil {
		return UserID{}, err
	}

	return UserID(rawid), nil
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func (c CourseID) String() string {
	return uuid.UUID(c).String()
}

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

type MockUUIDGenerator struct {
	Value uuid.UUID
}

func (m MockUUIDGenerator) Generate() (uuid.UUID, error) {
	return m.Value, nil
}
