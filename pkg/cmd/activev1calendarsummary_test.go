// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/clear-street-cli/internal/mocktest"
)

func TestActiveV1CalendarsSummaryGetCalendarSummary(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:calendars:summary", "get-calendar-summary",
			"--from", "'2019-12-27'",
			"--to", "'2019-12-27'",
		)
	})
}
