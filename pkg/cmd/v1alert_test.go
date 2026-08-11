// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1AlertsCreateAlert(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:alerts", "create-alert",
			"--condition", "{conditions: [{op: lte, signal: market.day_change_pct, subject: {instrument_id: NVDA}, value: -5}], match: all}",
			"--schedule", "every_1m",
			"--trigger", "once",
			"--account-id", "19816",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"condition:\n" +
			"  conditions:\n" +
			"    - op: lte\n" +
			"      signal: market.day_change_pct\n" +
			"      subject:\n" +
			"        instrument_id: NVDA\n" +
			"      value: -5\n" +
			"  match: all\n" +
			"schedule: every_1m\n" +
			"trigger: once\n" +
			"account_id: 19816\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:alerts", "create-alert",
		)
	})
}

func TestV1AlertsDeleteAlert(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:alerts", "delete-alert",
			"--alert-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1AlertsGetAlertByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:alerts", "get-alert-by-id",
			"--alert-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1AlertsGetAlerts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:alerts", "get-alerts",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--status", "status",
		)
	})
}
