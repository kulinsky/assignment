package command

import "github.com/kulinsky/assignment/internal/domain/assignment"

type CreateAssignmentCmd struct {
	CourseID string `json:"course_id"`
	UserID   string `json:"user_id"`
}

type CreateAssignmentCmdHandler struct {
	repo    assignment.Repository
	uuidgen assignment.UUIDGenerator
}

func NewCreateAssignmentCmdHandler(
	repo assignment.Repository,
	uuidgen assignment.UUIDGenerator,
) *CreateAssignmentCmdHandler {
	return &CreateAssignmentCmdHandler{
		repo:    repo,
		uuidgen: uuidgen,
	}
}

func (h *CreateAssignmentCmdHandler) Handle(cmd *CreateAssignmentCmd) (*assignment.ID, error) {
	rawID, err := h.uuidgen()
	if err != nil {
		return nil, err
	}

	id := assignment.ID(rawID)

	courseID, err := assignment.CourseIDFromString(cmd.CourseID)
	if err != nil {
		return nil, err
	}

	userID, err := assignment.UserIDFromString(cmd.UserID)
	if err != nil {
		return nil, err
	}

	assignment := assignment.NewAssignment(
		id,
		userID,
		courseID,
	)

	if err := h.repo.Add(assignment); err != nil {
		return nil, err
	}

	return &id, nil
}
