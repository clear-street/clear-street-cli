// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1LocatesGetLocateByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:locates", "get-locate-by-id",
			"--account-id", "0",
			"--id", "id",
		)
	})
}

func TestV1LocatesGetLocateRateByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:locates", "get-locate-rate-by-id",
			"--instrument-id", "x",
		)
	})
}

func TestV1LocatesGetLocateRates(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:locates", "get-locate-rates",
			"--instrument-id", "x",
		)
	})
}

func TestV1LocatesGetLocates(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:locates", "get-locates",
			"--account-id", "0",
			"--effective-date", "'2019-12-27'",
			"--instrument-id", "x",
			"--status", "REQUESTED",
		)
	})
}
