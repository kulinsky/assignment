package assignment_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kulinsky/assignment/internal/domain/assignment"
)

func TestNewAssignment(t *testing.T) {
	t.Parallel()

	t.Run("should create new assignment", func(t *testing.T) {
		t.Parallel()

		// Given
		id := assignment.ID(uuid.New())
		userID := assignment.UserID(uuid.New())
		courseID := assignment.CourseID(uuid.New())

		// When
		sut := assignment.NewAssignment(id, userID, courseID)

		// Then
		assert.Equal(t, id, sut.ID)
		assert.Equal(t, userID, sut.UserID)
		assert.Equal(t, courseID, sut.CourseID)
		assert.Nil(t, sut.DateStart)
		assert.Nil(t, sut.DateEnd)
		assert.Equal(t, assignment.StatusNew, sut.Status)
	})
}

func TestAssignment_Start(t *testing.T) {
	t.Parallel()

	t.Run("should start assignment", func(t *testing.T) {
		t.Parallel()

		// Given
		id := assignment.ID(uuid.New())
		userID := assignment.UserID(uuid.New())
		courseID := assignment.CourseID(uuid.New())
		sut := assignment.NewAssignment(id, userID, courseID)
		now := time.Date(2021, 10, 10, 0, 0, 0, 0, time.UTC)

		// When
		err := sut.Start(now)

		// Then
		assert.Nil(t, err)
		assert.Equal(t, assignment.StatusInProgress, sut.Status)
		assert.NotNil(t, sut.DateStart)
		assert.Equal(t, now, sut.DateStart.Time())
	})

	t.Run("should return error when assignment is not new", func(t *testing.T) {
		t.Parallel()

		// Given
		id := assignment.ID(uuid.New())
		userID := assignment.UserID(uuid.New())
		courseID := assignment.CourseID(uuid.New())
		sut := assignment.NewAssignment(id, userID, courseID)
		now := time.Date(2021, 10, 10, 0, 0, 0, 0, time.UTC)
		err := sut.Start(now)
		require.Nil(t, err)

		// When
		err = sut.Start(now.Add(time.Hour))

		// Then
		assert.Equal(t, assignment.ErrAssignmentNotReadyToStart, err)
		assert.Equal(t, assignment.StatusInProgress, sut.Status)
		assert.NotNil(t, sut.DateStart)
		assert.Equal(t, now, sut.DateStart.Time())
	})
}

func TestFinish(t *testing.T) {
	t.Parallel()

	t.Run("should finish assignment", func(t *testing.T) {
		t.Parallel()

		// Given
		id := assignment.ID(uuid.New())
		userID := assignment.UserID(uuid.New())
		courseID := assignment.CourseID(uuid.New())
		sut := assignment.NewAssignment(id, userID, courseID)
		now := time.Date(2021, 10, 10, 0, 0, 0, 0, time.UTC)
		err := sut.Start(now)
		require.NoError(t, err)

		// When
		err = sut.Finish(now.Add(time.Hour))

		// Then
		assert.Nil(t, err)
		assert.Equal(t, assignment.StatusFinished, sut.Status)
		assert.NotNil(t, sut.DateEnd)
		assert.Equal(t, now.Add(time.Hour), sut.DateEnd.Time())
	})

	t.Run("should return error when assignment is not in progress", func(t *testing.T) {
		t.Parallel()

		// Given
		id := assignment.ID(uuid.New())
		userID := assignment.UserID(uuid.New())
		courseID := assignment.CourseID(uuid.New())
		sut := assignment.NewAssignment(id, userID, courseID)
		now := time.Date(2021, 10, 10, 0, 0, 0, 0, time.UTC)

		// When
		err := sut.Finish(now)

		// Then
		assert.Equal(t, assignment.ErrAssignmentNotReadyToFinish, err)
		assert.Equal(t, assignment.StatusNew, sut.Status)
		assert.Nil(t, sut.DateEnd)
	})
}
