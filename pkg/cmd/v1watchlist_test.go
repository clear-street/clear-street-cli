// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1WatchlistAddWatchlistItem(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlist", "add-watchlist-item",
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
			"v1:watchlist", "add-watchlist-item",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1WatchlistCreateWatchlist(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlist", "create-watchlist",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: name")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:watchlist", "create-watchlist",
		)
	})
}

func TestV1WatchlistDeleteWatchlist(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlist", "delete-watchlist",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1WatchlistDeleteWatchlistItem(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlist", "delete-watchlist-item",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
			"--item-id", "660e8400-e29b-41d4-a716-446655440001",
		)
	})
}

func TestV1WatchlistGetWatchlistByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlist", "get-watchlist-by-id",
			"--watchlist-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1WatchlistGetWatchlists(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:watchlist", "get-watchlists",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}
