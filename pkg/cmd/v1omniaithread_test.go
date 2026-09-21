// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1OmniAIThreadsCreateMessage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-message",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--text", "Compare that to AMD.",
			"--account-id", "19816",
			"--capability", "PREFILL_ORDER",
			"--context", "{items: [{data: {change_pct: bar, range: bar, ticker: bar}, kind: chart, label: NVDA intraday performance, captured_at: '2019-12-27T18:11:19.117Z'}]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1OmniAIThreadsCreateMessage)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-message",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--text", "Compare that to AMD.",
			"--account-id", "19816",
			"--capability", "PREFILL_ORDER",
			"--context.items", "[{data: {change_pct: bar, range: bar, ticker: bar}, kind: chart, label: NVDA intraday performance, captured_at: '2019-12-27T18:11:19.117Z'}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"text: Compare that to AMD.\n" +
			"account_id: 19816\n" +
			"capabilities:\n" +
			"  - PREFILL_ORDER\n" +
			"context:\n" +
			"  items:\n" +
			"    - data:\n" +
			"        change_pct: bar\n" +
			"        range: bar\n" +
			"        ticker: bar\n" +
			"      kind: chart\n" +
			"      label: NVDA intraday performance\n" +
			"      captured_at: '2019-12-27T18:11:19.117Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-message",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1OmniAIThreadsCreateThread(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "create-thread",
			"--type", "instant",
			"--account-id", "19816",
			"--capability", "PREFILL_ORDER",
			"--context", "{items: [{data: {change_pct: bar, range: bar, ticker: bar}, kind: chart, label: NVDA intraday performance, captured_at: '2019-12-27T18:11:19.117Z'}]}",
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
			"--type", "instant",
			"--account-id", "19816",
			"--capability", "PREFILL_ORDER",
			"--context.items", "[{data: {change_pct: bar, range: bar, ticker: bar}, kind: chart, label: NVDA intraday performance, captured_at: '2019-12-27T18:11:19.117Z'}]",
			"--target.ticker", "ticker",
			"--target.type", "ticker",
			"--text", "What changed in NVDA today?",
			"--thesis", "thesis",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: instant\n" +
			"account_id: 19816\n" +
			"capabilities:\n" +
			"  - PREFILL_ORDER\n" +
			"context:\n" +
			"  items:\n" +
			"    - data:\n" +
			"        change_pct: bar\n" +
			"        range: bar\n" +
			"        ticker: bar\n" +
			"      kind: chart\n" +
			"      label: NVDA intraday performance\n" +
			"      captured_at: '2019-12-27T18:11:19.117Z'\n" +
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
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-messages",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "1",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}

func TestV1OmniAIThreadsGetThreadByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-thread-by-id",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "1",
		)
	})
}

func TestV1OmniAIThreadsGetThreadResponse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-thread-response",
			"--thread-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "1",
		)
	})
}

func TestV1OmniAIThreadsGetThreads(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:threads", "get-threads",
			"--account-id", "1",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}
