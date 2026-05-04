// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1WatchlistsItemsAddWatchlistItem(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlists:items", "add-watchlist-item",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("instrument_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:watchlists:items", "add-watchlist-item",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1WatchlistsItemsDeleteWatchlistItem(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlists:items", "delete-watchlist-item",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
			"--item-id", "660e8400-e29b-41d4-a716-446655440001",
		)
	})
}
