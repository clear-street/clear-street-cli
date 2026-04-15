// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/clear-street-cli/internal/mocktest"
)

func TestActiveV1IrisRunsCancelRunDeprecated(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:iris:runs", "cancel-run-deprecated",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "account_id",
			"--reason", "reason",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_id\n" +
			"reason: reason\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:iris:runs", "cancel-run-deprecated",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestActiveV1IrisRunsGetRunDeprecated(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:iris:runs", "get-run-deprecated",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "account_id",
			"--page-size", "0",
			"--page-token", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestActiveV1IrisRunsStartRunDeprecated(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:iris:runs", "start-run-deprecated",
			"--account-id", "account_id",
			"--command-text", "command_text",
			"--capability", "NAVIGATE",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--thread-title", "thread_title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_id\n" +
			"command_text: command_text\n" +
			"capabilities:\n" +
			"  - NAVIGATE\n" +
			"thread_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"thread_title: thread_title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:iris:runs", "start-run-deprecated",
		)
	})
}
