// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1SavedScreenersCreateScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:saved-screeners", "create-screener",
			"--field-filter", "[{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}]",
			"--filter", "[{left: {name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}, op: {name: GTE, args: [LEFT_INCLUSIVE]}, right: [{value: 1000000000, variable: {name: today, lookback: ONE_WEEK, modifier: {args: [30, DAY], name: SUB}, period: QUARTER}}]}]",
			"--name", "name",
			"--sort-by", "{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}",
			"--sort-direction", "ASC",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1SavedScreenersCreateScreener)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:saved-screeners", "create-screener",
			"--field-filter.name", "market_cap",
			"--field-filter.lookback", "ONE_WEEK",
			"--field-filter.period", "QUARTER",
			"--field-filter.value-type", "DECIMAL",
			"--filter.left", "{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}",
			"--filter.op", "{name: GTE, args: [LEFT_INCLUSIVE]}",
			"--filter.right", "[{value: 1000000000, variable: {name: today, lookback: ONE_WEEK, modifier: {args: [30, DAY], name: SUB}, period: QUARTER}}]",
			"--name", "name",
			"--sort-by.name", "market_cap",
			"--sort-by.lookback", "ONE_WEEK",
			"--sort-by.period", "QUARTER",
			"--sort-by.value-type", "DECIMAL",
			"--sort-direction", "ASC",
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
			"name: name\n" +
			"sort_by:\n" +
			"  name: market_cap\n" +
			"  lookback: ONE_WEEK\n" +
			"  period: QUARTER\n" +
			"  value_type: DECIMAL\n" +
			"sort_direction: ASC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:saved-screeners", "create-screener",
		)
	})
}

func TestV1SavedScreenersDeleteScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:saved-screeners", "delete-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1SavedScreenersGetScreenerByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:saved-screeners", "get-screener-by-id",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestV1SavedScreenersGetScreeners(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:saved-screeners", "get-screeners",
		)
	})
}

func TestV1SavedScreenersReplaceScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:saved-screeners", "replace-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
			"--field-filter", "[{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}]",
			"--filter", "[{left: {name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}, op: {name: GTE, args: [LEFT_INCLUSIVE]}, right: [{value: 1000000000, variable: {name: today, lookback: ONE_WEEK, modifier: {args: [30, DAY], name: SUB}, period: QUARTER}}]}]",
			"--name", "name",
			"--sort-by", "{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}",
			"--sort-direction", "ASC",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1SavedScreenersReplaceScreener)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:saved-screeners", "replace-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
			"--field-filter.name", "market_cap",
			"--field-filter.lookback", "ONE_WEEK",
			"--field-filter.period", "QUARTER",
			"--field-filter.value-type", "DECIMAL",
			"--filter.left", "{name: market_cap, lookback: ONE_WEEK, period: QUARTER, value_type: DECIMAL}",
			"--filter.op", "{name: GTE, args: [LEFT_INCLUSIVE]}",
			"--filter.right", "[{value: 1000000000, variable: {name: today, lookback: ONE_WEEK, modifier: {args: [30, DAY], name: SUB}, period: QUARTER}}]",
			"--name", "name",
			"--sort-by.name", "market_cap",
			"--sort-by.lookback", "ONE_WEEK",
			"--sort-by.period", "QUARTER",
			"--sort-by.value-type", "DECIMAL",
			"--sort-direction", "ASC",
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
			"name: name\n" +
			"sort_by:\n" +
			"  name: market_cap\n" +
			"  lookback: ONE_WEEK\n" +
			"  period: QUARTER\n" +
			"  value_type: DECIMAL\n" +
			"sort_direction: ASC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:saved-screeners", "replace-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}
