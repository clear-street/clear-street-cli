// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestActiveV1SavedScreenersCreateScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:saved-screeners", "create-screener",
			"--field-filter", "[string]",
			"--filter", "[{field_name: field_name, operation: operation, value: value}]",
			"--name", "name",
			"--sort-by", "sort_by",
			"--sort-direction", "ASC",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(activeV1SavedScreenersCreateScreener)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:saved-screeners", "create-screener",
			"--field-filter", "[string]",
			"--filter.field-name", "field_name",
			"--filter.operation", "operation",
			"--filter.value", "value",
			"--name", "name",
			"--sort-by", "sort_by",
			"--sort-direction", "ASC",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"field_filter:\n" +
			"  - string\n" +
			"filters:\n" +
			"  - field_name: field_name\n" +
			"    operation: operation\n" +
			"    value: value\n" +
			"name: name\n" +
			"sort_by: sort_by\n" +
			"sort_direction: ASC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:saved-screeners", "create-screener",
		)
	})
}

func TestActiveV1SavedScreenersDeleteScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:saved-screeners", "delete-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestActiveV1SavedScreenersGetScreenerByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:saved-screeners", "get-screener-by-id",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestActiveV1SavedScreenersListScreeners(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:saved-screeners", "list-screeners",
		)
	})
}

func TestActiveV1SavedScreenersUpdateScreener(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:saved-screeners", "update-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
			"--field-filter", "[string]",
			"--filter", "[{field_name: field_name, operation: operation, value: value}]",
			"--name", "name",
			"--sort-by", "sort_by",
			"--sort-direction", "ASC",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(activeV1SavedScreenersUpdateScreener)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:saved-screeners", "update-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
			"--field-filter", "[string]",
			"--filter.field-name", "field_name",
			"--filter.operation", "operation",
			"--filter.value", "value",
			"--name", "name",
			"--sort-by", "sort_by",
			"--sort-direction", "ASC",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"field_filter:\n" +
			"  - string\n" +
			"filters:\n" +
			"  - field_name: field_name\n" +
			"    operation: operation\n" +
			"    value: value\n" +
			"name: name\n" +
			"sort_by: sort_by\n" +
			"sort_direction: ASC\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:saved-screeners", "update-screener",
			"--screener-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}
