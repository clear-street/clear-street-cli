// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1NewsGetNews(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:news", "get-news",
			"--exclude-publishers", "exclude_publishers",
			"--from", "from",
			"--include-publishers", "include_publishers",
			"--instrument-id", "string",
			"--news-type", "NEWS",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--search-query", "search_query",
			"--sector", "BASIC_MATERIALS",
			"--security-id", "string",
			"--security-id-source", "string",
			"--to", "to",
		)
	})
}
