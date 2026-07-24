// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1InstrumentDataGetAllInstrumentEvents(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data", "get-all-instrument-events",
			"--event-type", "EARNINGS",
			"--from-date", "from_date",
			"--instrument-id", "x",
			"--to-date", "to_date",
		)
	})
}

func TestV1InstrumentDataGetInstrumentAnalystConsensus(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data", "get-instrument-analyst-consensus",
			"--instrument-id", "x",
			"--from", "'2019-12-27'",
			"--to", "'2019-12-27'",
		)
	})
}

func TestV1InstrumentDataGetInstrumentBalanceSheetStatements(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data", "get-instrument-balance-sheet-statements",
			"--instrument-id", "x",
			"--from-date", "from_date",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--to-date", "to_date",
		)
	})
}

func TestV1InstrumentDataGetInstrumentCashFlowStatements(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data", "get-instrument-cash-flow-statements",
			"--instrument-id", "x",
			"--from-date", "from_date",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--to-date", "to_date",
		)
	})
}

func TestV1InstrumentDataGetInstrumentEvents(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data", "get-instrument-events",
			"--instrument-id", "x",
			"--from-date", "from_date",
			"--to-date", "to_date",
		)
	})
}

func TestV1InstrumentDataGetInstrumentFundamentals(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data", "get-instrument-fundamentals",
			"--instrument-id", "x",
		)
	})
}

func TestV1InstrumentDataGetInstrumentIncomeStatements(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data", "get-instrument-income-statements",
			"--instrument-id", "x",
			"--from-date", "from_date",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--to-date", "to_date",
		)
	})
}
