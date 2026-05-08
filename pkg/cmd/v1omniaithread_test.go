// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1OmniAIThreadsCreateMessage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-message",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "19816",
			"--text", "Compare that to AMD.",
			"--capability", "PREFILL_ORDER",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: 19816\n" +
			"text: Compare that to AMD.\n" +
			"capabilities:\n" +
			"  - PREFILL_ORDER\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-message",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1OmniAIThreadsCreateThread(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-thread",
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
		requestflag.CheckInnerFlags(v1OmniAIThreadsCreateThread)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-thread",
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
			"v1:omni-ai:threads", "create-thread",
		)
	})
}

func TestV1OmniAIThreadsGetMessages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-messages",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}

func TestV1OmniAIThreadsGetThreadByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-thread-by-id",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}

func TestV1OmniAIThreadsGetThreadResponse(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-thread-response",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}

func TestV1OmniAIThreadsGetThreads(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-threads",
			"--account-id", "0",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}
