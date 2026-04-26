// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestActiveV1OmniAIThreadsCreateThread(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:omni-ai:threads", "create-thread",
			"--account-id", "19816",
			"--type", "instant",
			"--capability", "PREFILL_ORDER",
			"--target", "{ticker: ticker, type: ticker}",
			"--text", "What changed in NVDA today?",
			"--thesis", "thesis",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(activeV1OmniAIThreadsCreateThread)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:omni-ai:threads", "create-thread",
			"--account-id", "19816",
			"--type", "instant",
			"--capability", "PREFILL_ORDER",
			"--target.ticker", "ticker",
			"--target.type", "ticker",
			"--text", "What changed in NVDA today?",
			"--thesis", "thesis",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: 19816\n" +
			"type: instant\n" +
			"capabilities:\n" +
			"  - PREFILL_ORDER\n" +
			"target:\n" +
			"  ticker: ticker\n" +
			"  type: ticker\n" +
			"text: What changed in NVDA today?\n" +
			"thesis: thesis\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:omni-ai:threads", "create-thread",
		)
	})
}

func TestActiveV1OmniAIThreadsGetThread(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:omni-ai:threads", "get-thread",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}

func TestActiveV1OmniAIThreadsListThreads(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:omni-ai:threads", "list-threads",
			"--account-id", "0",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}

func TestActiveV1OmniAIThreadsResponse(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:omni-ai:threads", "response",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}
