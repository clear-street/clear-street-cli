// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1OmniAIMessagesGetMessageByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:messages", "get-message-by-id",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "1",
		)
	})
}

func TestV1OmniAIMessagesSubmitFeedback(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:messages", "submit-feedback",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--score", "0",
			"--account-id", "1",
			"--comment", "comment",
			"--metadata", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"score: 0\n" +
			"account_id: 1\n" +
			"comment: comment\n" +
			"metadata: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:omni-ai:messages", "submit-feedback",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
