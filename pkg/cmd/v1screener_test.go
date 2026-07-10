// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1ScreenerCreateScreener(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "create-screener",
			"--column", "[{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}]",
			"--filter", "[{left: {name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, op: {name: GREATER_OR_EQUAL, args: [LEFT_INCLUSIVE]}, right: [{value: 1000000000, variable: {name: today, lookback: ONE_DAY, modifier: {args: [30, DAY], name: SUBTRACT}, period: QUARTER}}]}]",
			"--name", "name",
			"--sort", "[{field: {name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, direction: DESC}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1ScreenerCreateScreener)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "create-screener",
			"--column.name", "market_cap",
			"--column.lookback", "ONE_DAY",
			"--column.period", "QUARTER",
			"--column.value-type", "DECIMAL",
			"--filter.left", "{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}",
			"--filter.op", "{name: GREATER_OR_EQUAL, args: [LEFT_INCLUSIVE]}",
			"--filter.right", "[{value: 1000000000, variable: {name: today, lookback: ONE_DAY, modifier: {args: [30, DAY], name: SUBTRACT}, period: QUARTER}}]",
			"--name", "name",
			"--sort.field", "{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}",
			"--sort.direction", "DESC",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"columns:\n" +
			"  - name: market_cap\n" +
			"    lookback: ONE_DAY\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"filters:\n" +
			"  - left:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_DAY\n" +
			"      period: QUARTER\n" +
			"      value_type: DECIMAL\n" +
			"    op:\n" +
			"      name: GREATER_OR_EQUAL\n" +
			"      args:\n" +
			"        - LEFT_INCLUSIVE\n" +
			"    right:\n" +
			"      - value: 1000000000\n" +
			"        variable:\n" +
			"          name: today\n" +
			"          lookback: ONE_DAY\n" +
			"          modifier:\n" +
			"            args:\n" +
			"              - 30\n" +
			"              - DAY\n" +
			"            name: SUBTRACT\n" +
			"          period: QUARTER\n" +
			"name: name\n" +
			"sorts:\n" +
			"  - field:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_DAY\n" +
			"      period: QUARTER\n" +
			"      value_type: DECIMAL\n" +
			"    direction: DESC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:screener", "create-screener",
		)
	})
}

func TestV1ScreenerDeleteScreener(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "delete-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1ScreenerGetScreenerByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "get-screener-by-id",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1ScreenerGetScreeners(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "get-screeners",
		)
	})
}

func TestV1ScreenerReplaceScreener(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "replace-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
			"--column", "[{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}]",
			"--filter", "[{left: {name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, op: {name: GREATER_OR_EQUAL, args: [LEFT_INCLUSIVE]}, right: [{value: 1000000000, variable: {name: today, lookback: ONE_DAY, modifier: {args: [30, DAY], name: SUBTRACT}, period: QUARTER}}]}]",
			"--name", "name",
			"--sort", "[{field: {name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, direction: DESC}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1ScreenerReplaceScreener)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "replace-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
			"--column.name", "market_cap",
			"--column.lookback", "ONE_DAY",
			"--column.period", "QUARTER",
			"--column.value-type", "DECIMAL",
			"--filter.left", "{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}",
			"--filter.op", "{name: GREATER_OR_EQUAL, args: [LEFT_INCLUSIVE]}",
			"--filter.right", "[{value: 1000000000, variable: {name: today, lookback: ONE_DAY, modifier: {args: [30, DAY], name: SUBTRACT}, period: QUARTER}}]",
			"--name", "name",
			"--sort.field", "{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}",
			"--sort.direction", "DESC",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"columns:\n" +
			"  - name: market_cap\n" +
			"    lookback: ONE_DAY\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"filters:\n" +
			"  - left:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_DAY\n" +
			"      period: QUARTER\n" +
			"      value_type: DECIMAL\n" +
			"    op:\n" +
			"      name: GREATER_OR_EQUAL\n" +
			"      args:\n" +
			"        - LEFT_INCLUSIVE\n" +
			"    right:\n" +
			"      - value: 1000000000\n" +
			"        variable:\n" +
			"          name: today\n" +
			"          lookback: ONE_DAY\n" +
			"          modifier:\n" +
			"            args:\n" +
			"              - 30\n" +
			"              - DAY\n" +
			"            name: SUBTRACT\n" +
			"          period: QUARTER\n" +
			"name: name\n" +
			"sorts:\n" +
			"  - field:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_DAY\n" +
			"      period: QUARTER\n" +
			"      value_type: DECIMAL\n" +
			"    direction: DESC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:screener", "replace-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1ScreenerSearchScreener(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:screener", "search-screener",
			"--column", "[{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, {name: price, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, {name: volume, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}]",
			"--filter", "[{left: {name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, op: {name: GREATER_OR_EQUAL, args: [LEFT_INCLUSIVE]}, right: [{value: 1000000000, variable: {name: today, lookback: ONE_DAY, modifier: {args: [30, DAY], name: SUBTRACT}, period: QUARTER}}]}]",
			"--page-size", "25",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--sort-case-sensitive=true",
			"--sort", "[{field: {name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}, direction: DESC}]",
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
			"--column.name", "market_cap",
			"--column.lookback", "ONE_DAY",
			"--column.period", "QUARTER",
			"--column.value-type", "DECIMAL",
			"--column.name", "price",
			"--column.lookback", "ONE_DAY",
			"--column.period", "QUARTER",
			"--column.value-type", "DECIMAL",
			"--column.name", "volume",
			"--column.lookback", "ONE_DAY",
			"--column.period", "QUARTER",
			"--column.value-type", "DECIMAL",
			"--filter.left", "{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}",
			"--filter.op", "{name: GREATER_OR_EQUAL, args: [LEFT_INCLUSIVE]}",
			"--filter.right", "[{value: 1000000000, variable: {name: today, lookback: ONE_DAY, modifier: {args: [30, DAY], name: SUBTRACT}, period: QUARTER}}]",
			"--page-size", "25",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--sort-case-sensitive=true",
			"--sort.field", "{name: market_cap, lookback: ONE_DAY, period: QUARTER, value_type: DECIMAL}",
			"--sort.direction", "DESC",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"columns:\n" +
			"  - name: market_cap\n" +
			"    lookback: ONE_DAY\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"  - name: price\n" +
			"    lookback: ONE_DAY\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"  - name: volume\n" +
			"    lookback: ONE_DAY\n" +
			"    period: QUARTER\n" +
			"    value_type: DECIMAL\n" +
			"filters:\n" +
			"  - left:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_DAY\n" +
			"      period: QUARTER\n" +
			"      value_type: DECIMAL\n" +
			"    op:\n" +
			"      name: GREATER_OR_EQUAL\n" +
			"      args:\n" +
			"        - LEFT_INCLUSIVE\n" +
			"    right:\n" +
			"      - value: 1000000000\n" +
			"        variable:\n" +
			"          name: today\n" +
			"          lookback: ONE_DAY\n" +
			"          modifier:\n" +
			"            args:\n" +
			"              - 30\n" +
			"              - DAY\n" +
			"            name: SUBTRACT\n" +
			"          period: QUARTER\n" +
			"page_size: 25\n" +
			"page_token: U3RhaW5sZXNzIHJvY2tz\n" +
			"sort_case_sensitive: true\n" +
			"sorts:\n" +
			"  - field:\n" +
			"      name: market_cap\n" +
			"      lookback: ONE_DAY\n" +
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
