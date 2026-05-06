// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1AccountsExercisesCancelExercise(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:exercises", "cancel-exercise",
			"--account-id", "0",
			"--exercise-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1AccountsExercisesGetExercises(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:exercises", "get-exercises",
			"--account-id", "0",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1AccountsExercisesSubmitExercises(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:exercises", "submit-exercises",
			"--account-id", "0",
			"--exercise", "{action: EXERCISE, instrument_id: 0195f6d0-a1b2-7c3d-8e4f-5a6b7c8d9e02, quantity: '1', client_exercise_id: ui-20260424-001}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1AccountsExercisesSubmitExercises)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:exercises", "submit-exercises",
			"--account-id", "0",
			"--exercise.action", "EXERCISE",
			"--exercise.instrument-id", "0195f6d0-a1b2-7c3d-8e4f-5a6b7c8d9e02",
			"--exercise.quantity", "1",
			"--exercise.client-exercise-id", "ui-20260424-001",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- action: EXERCISE\n" +
			"  instrument_id: 0195f6d0-a1b2-7c3d-8e4f-5a6b7c8d9e02\n" +
			"  quantity: '1'\n" +
			"  client_exercise_id: ui-20260424-001\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:accounts:exercises", "submit-exercises",
			"--account-id", "0",
		)
	})
}
