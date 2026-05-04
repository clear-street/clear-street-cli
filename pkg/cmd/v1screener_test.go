// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1ScreenerGetScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "get-screener",
			"--field-filter", "string",
			"--filter", "{foo: string}",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--sort-by", "sort_by",
			"--sort-direction", "ASC",
		)
	})
}

func TestV1ScreenerSearchScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "search-screener",
			"--field-filter", "[{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}, {name: price, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}, {name: volume, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}]",
			"--filter", "[{left: {name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}, op: {name: GTE, args: [LEFT_INCLUSIVE]}, right: [{value: 1000000000, variable: {name: today, lookback: ONE_WEEK, modifier: {args: [30, DAY], name: SUB}, period: QUARTER}}]}]",
			"--page-size", "25",
			"--page-token", "page_token",
			"--sort-by", "{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}",
			"--sort-case-sensitive=true",
			"--sort-direction", "ASC",
			"--sort", "[{field: {name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}, direction: DESC}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1ScreenerSearchScreener)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "search-screener",
			"--field-filter.name", "market_cap",
			"--field-filter.lookback", "ONE_WEEK",
			"--field-filter.period", "QUARTER",
			"--field-filter.value-type", "DECIMAL",
			"--field-filter.name", "price",
			"--field-filter.lookback", "ONE_WEEK",
			"--field-filter.period", "QUARTER",
			"--field-filter.value-type", "DECIMAL",
			"--field-filter.name", "volume",
			"--field-filter.lookback", "ONE_WEEK",
			"--field-filter.period", "QUARTER",
			"--field-filter.value-type", "DECIMAL",
			"--filter.left", "{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}",
			"--filter.op", "{name: GTE, args: [LEFT_INCLUSIVE]}",
			"--filter.right", "[{value: 1000000000, variable: {name: today, lookback: ONE_WEEK, modifier: {args: [30, DAY], name: SUB}, period: QUARTER}}]",
			"--page-size", "25",
			"--page-token", "page_token",
			"--sort-by.name", "market_cap",
			"--sort-by.lookback", "ONE_WEEK",
			"--sort-by.period", "QUARTER",
			"--sort-by.value-type", "DECIMAL",
			"--sort-case-sensitive=true",
			"--sort-direction", "ASC",
			"--sort.field", "{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}",
			"--sort.direction", "DESC",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"field_filter:\n" +
			"  - name: market_cap\n" +
			"    lookback: ONE_WEEK\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"  - name: price\n" +
			"    lookback: ONE_WEEK\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"  - name: volume\n" +
			"    lookback: ONE_WEEK\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"filters:\n" +
			"  - left:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_WEEK\n" +
			"      period: QUARTER\n" +
			"      value_type: DECIMAL\n" +
			"    op:\n" +
			"      name: GTE\n" +
			"      args:\n" +
			"        - LEFT_INCLUSIVE\n" +
			"    right:\n" +
			"      - value: 1000000000\n" +
			"        variable:\n" +
			"          name: today\n" +
			"          lookback: ONE_WEEK\n" +
			"          modifier:\n" +
			"            args:\n" +
			"              - 30\n" +
			"              - DAY\n" +
			"            name: SUB\n" +
			"          period: QUARTER\n" +
			"page_size: 25\n" +
			"page_token: page_token\n" +
			"sort_by:\n" +
			"  name: market_cap\n" +
			"  lookback: ONE_WEEK\n" +
			"  period: QUARTER\n" +
			"  value_type: DECIMAL\n" +
			"sort_case_sensitive: true\n" +
			"sort_direction: ASC\n" +
			"sorts:\n" +
			"  - field:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_WEEK\n" +
			"      period: QUARTER\n" +
			"      value_type: DECIMAL\n" +
			"    direction: DESC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:screener", "search-screener",
		)
	})
}
