// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1OmniAIThreadsMessagesListMessages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:omni-ai:threads:messages", "list-messages",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "account_id",
			"--after-seq", "0",
			"--page-size", "0",
			"--page-token", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
