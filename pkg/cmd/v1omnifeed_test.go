// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1OmniFeedGetFeed(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-feed", "get-feed",
			"--account-id", "0",
			"--cursor", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "0",
		)
	})
}

func TestV1OmniFeedPostFeedEvent(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-feed", "post-feed-event",
			"--event", "{item_id: 0198f3a2-4b3d-7c1e-9f2a-3b4c5d6e7f80, type: seen}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1OmniFeedPostFeedEvent)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-feed", "post-feed-event",
			"--event.item-id", "0198f3a2-4b3d-7c1e-9f2a-3b4c5d6e7f80",
			"--event.type", "seen",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"event:\n" +
			"  item_id: 0198f3a2-4b3d-7c1e-9f2a-3b4c5d6e7f80\n" +
			"  type: seen\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:omni-feed", "post-feed-event",
		)
	})
}
