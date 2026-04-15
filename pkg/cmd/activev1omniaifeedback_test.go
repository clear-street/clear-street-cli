// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1OmniAIFeedbackCreateFeedback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:omni-ai:feedback", "create-feedback",
			"--account-id", "account_id",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--score", "0",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--comment", "comment",
			"--metadata", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_id\n" +
			"message_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"score: 0\n" +
			"thread_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"comment: comment\n" +
			"metadata: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:omni-ai:feedback", "create-feedback",
		)
	})
}
