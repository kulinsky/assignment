package command_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kulinsky/assignment/internal/app/command"
	"github.com/kulinsky/assignment/internal/domain/assignment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAssignmentCmdHandler_Handle(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		// Given
		repo := &assignment.MockRepository{}
		repo.On("Add", mock.Anything).Return(nil).Once()
		rawuuid, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
		uuidgen := assignment.MockUUIDGenerator{Value: rawuuid}.Generate
		handler := command.NewCreateAssignmentCmdHandler(repo, uuidgen)
		cmd := &command.CreateAssignmentCmd{
			CourseID: uuid.New().String(),
			UserID:   uuid.New().String(),
		}

		// When
		id, err := handler.Handle(cmd)

		// Then
		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, rawuuid.String(), id.String())
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		// Given
		repo := &assignment.MockRepository{}
		rawuuid, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
		uuidgen := assignment.MockUUIDGenerator{Value: rawuuid}.Generate
		handler := command.NewCreateAssignmentCmdHandler(repo, uuidgen)
		cmd := &command.CreateAssignmentCmd{
			CourseID: "course_id",
			UserID:   "user_id",
		}

		// When
		id, err := handler.Handle(cmd)

		// Then
		assert.Error(t, err)
		assert.Nil(t, id)
		repo.AssertExpectations(t)
	})
}
