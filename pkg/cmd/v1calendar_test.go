// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1CalendarGetClock(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:calendar", "get-clock",
		)
	})
}

func TestV1CalendarGetEconomicEventsCalendar(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:calendar", "get-economic-events-calendar",
			"--country", "country",
			"--impact", "NONE",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--timestamp", "{gt: gt, gte: gte, lt: lt, lte: lte}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CalendarGetEconomicEventsCalendar)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:calendar", "get-economic-events-calendar",
			"--country", "country",
			"--impact", "NONE",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--timestamp.gt", "gt",
			"--timestamp.gte", "gte",
			"--timestamp.lt", "lt",
			"--timestamp.lte", "lte",
		)
	})
}

func TestV1CalendarGetMarketHoursCalendar(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:calendar", "get-market-hours-calendar",
			"--date", "date",
			"--market", "us_equities",
		)
	})
}
