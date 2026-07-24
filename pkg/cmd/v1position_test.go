// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1PositionsCancelPositionInstruction(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:positions", "cancel-position-instruction",
			"--account-id", "0",
			"--instruction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1PositionsClosePosition(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:positions", "close-position",
			"--account-id", "0",
			"--instrument-id", "x",
			"--cancel-orders=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("cancel_orders: false")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:positions", "close-position",
			"--account-id", "0",
			"--instrument-id", "x",
		)
	})
}

func TestV1PositionsClosePositions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:positions", "close-positions",
			"--account-id", "0",
			"--cancel-orders=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("cancel_orders: false")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:positions", "close-positions",
			"--account-id", "0",
		)
	})
}

func TestV1PositionsGetPositionInstructions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:positions", "get-position-instructions",
			"--account-id", "0",
			"--instrument-id", "x",
		)
	})
}

func TestV1PositionsGetPositions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:positions", "get-positions",
			"--account-id", "0",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--sort-by", "SYMBOL",
			"--sort-direction", "ASC",
		)
	})
}

func TestV1PositionsSubmitPositionInstructions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:positions", "submit-position-instructions",
			"--account-id", "0",
			"--instruction", "{instruction_type: EXERCISE, instrument_id: 0195f6d0-a1b2-7c3d-8e4f-5a6b7c8d9e02, quantity: '1', instruction_id: ui-20260424-001}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PositionsSubmitPositionInstructions)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:positions", "submit-position-instructions",
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
			"v1:positions", "submit-position-instructions",
			"--account-id", "0",
		)
	})
}
