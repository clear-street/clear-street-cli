// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1AccountsPositionsInstructionsCancelPositionInstruction(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:positions:instructions", "cancel-position-instruction",
			"--account-id", "0",
			"--instruction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1AccountsPositionsInstructionsGetPositionInstructions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:positions:instructions", "get-position-instructions",
			"--account-id", "0",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1AccountsPositionsInstructionsSubmitPositionInstructions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:positions:instructions", "submit-position-instructions",
			"--account-id", "0",
			"--instruction", "{instruction_type: EXERCISE, instrument_id: 0195f6d0-a1b2-7c3d-8e4f-5a6b7c8d9e02, quantity: '1', instruction_id: ui-20260424-001}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1AccountsPositionsInstructionsSubmitPositionInstructions)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:positions:instructions", "submit-position-instructions",
			"--account-id", "0",
			"--instruction.instruction-type", "EXERCISE",
			"--instruction.instrument-id", "0195f6d0-a1b2-7c3d-8e4f-5a6b7c8d9e02",
			"--instruction.quantity", "1",
			"--instruction.instruction-id", "ui-20260424-001",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- instruction_type: EXERCISE\n" +
			"  instrument_id: 0195f6d0-a1b2-7c3d-8e4f-5a6b7c8d9e02\n" +
			"  quantity: '1'\n" +
			"  instruction_id: ui-20260424-001\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:accounts:positions:instructions", "submit-position-instructions",
			"--account-id", "0",
		)
	})
}
