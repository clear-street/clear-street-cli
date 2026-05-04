// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1OmniAIMessagesFeedback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:messages", "feedback",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
			"--score", "0",
			"--comment", "comment",
			"--metadata", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: 0\n" +
			"score: 0\n" +
			"comment: comment\n" +
			"metadata: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:omni-ai:messages", "feedback",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1OmniAIMessagesGetMessage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:messages", "get-message",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}
