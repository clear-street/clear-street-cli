// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1InstrumentsOptionsContractsGetOptionContracts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:instruments:options:contracts", "get-option-contracts",
			"--contract-type", "CALL",
			"--expiry", "'2019-12-27'",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--underlier", "underlier",
			"--underlier-instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--underlier-security-id", "underlier_security_id",
			"--underlier-security-id-source", "CMS",
		)
	})
}
