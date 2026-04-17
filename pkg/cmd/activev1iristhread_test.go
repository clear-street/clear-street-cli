// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1IrisThreadsGetThreadDeprecated(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:iris:threads", "get-thread-deprecated",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "account_id",
		)
	})
}

func TestActiveV1IrisThreadsListThreadsDeprecated(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:iris:threads", "list-threads-deprecated",
			"--account-id", "account_id",
			"--page-size", "0",
			"--page-token", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
