// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1CalendarsMarketHoursGetMarketHoursCalendar(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:calendars:market-hours", "get-market-hours-calendar",
			"--date", "date",
			"--market", "us_equities",
		)
	})
}
