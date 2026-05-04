// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1WatchlistsCreateWatchlist(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlists", "create-watchlist",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: name")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:watchlists", "create-watchlist",
		)
	})
}

func TestV1WatchlistsDeleteWatchlist(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlists", "delete-watchlist",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1WatchlistsGetWatchlistByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlists", "get-watchlist-by-id",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1WatchlistsGetWatchlists(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlists", "get-watchlists",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}
