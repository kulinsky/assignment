package assignment_test

import (
	"testing"

	"github.com/kulinsky/assignment/internal/domain/assignment"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	t.Parallel()

	t.Run("NewStatusFromInt", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			name        string
			input       int
			expected    assignment.Status
			expectedErr error
		}{
			{
				name:        "New",
				input:       1,
				expected:    assignment.StatusNew,
				expectedErr: nil,
			},
			{
				name:        "InProgress",
				input:       2,
				expected:    assignment.StatusInProgress,
				expectedErr: nil,
			},
			{
				name:        "Finished",
				input:       3,
				expected:    assignment.StatusFinished,
				expectedErr: nil,
			},
			{
				name:        "Invalid",
				input:       4,
				expected:    0,
				expectedErr: assignment.ErrInvalidStatus,
			},
			{
				name:        "Invalid",
				input:       0,
				expected:    0,
				expectedErr: assignment.ErrInvalidStatus,
			},
		}

		for _, tc := range cases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				actual, err := assignment.NewStatusFromInt(tc.input)
				if tc.expectedErr != nil {
					if err == nil {
						t.Errorf("expected error %v, got nil", tc.expectedErr)
					} else if !assert.ErrorIs(t, err, tc.expectedErr) {
						t.Errorf("expected error %v, got %v", tc.expectedErr, err)
					}
				} else {
					if err != nil {
						t.Errorf("expected no error, got %v", err)
					}
					if actual != tc.expected {
						t.Errorf("expected %v, got %v", tc.expected, actual)
					}
				}
			})
		}
	})
}
