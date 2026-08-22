// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1PrivateMarketsCreateIoi(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "create-ioi",
			"--account-id", "0",
			"--notional-amount", "100000.00",
			"--offering-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nda-acceptance", "{accepted: true, agreement_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, authority_confirmed: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PrivateMarketsCreateIoi)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "create-ioi",
			"--account-id", "0",
			"--notional-amount", "100000.00",
			"--offering-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nda-acceptance.accepted=true",
			"--nda-acceptance.agreement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nda-acceptance.authority-confirmed=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"notional_amount: '100000.00'\n" +
			"offering_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"nda_acceptance:\n" +
			"  accepted: true\n" +
			"  agreement_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  authority_confirmed: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:private-markets", "create-ioi",
			"--account-id", "0",
		)
	})
}

func TestV1PrivateMarketsDeleteIoi(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "delete-ioi",
			"--ioi-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}

func TestV1PrivateMarketsGetCompanyByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "get-company-by-id",
			"--company-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}

func TestV1PrivateMarketsGetIois(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "get-iois",
			"--account-id", "0",
		)
	})
}

func TestV1PrivateMarketsGetSpvByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "get-spv-by-id",
			"--spv-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}

func TestV1PrivateMarketsUpdateIoi(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "update-ioi",
			"--ioi-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
			"--notional-amount", "125000.00",
			"--nda-acceptance", "{accepted: true, agreement_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, authority_confirmed: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1PrivateMarketsUpdateIoi)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:private-markets", "update-ioi",
			"--ioi-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
			"--notional-amount", "125000.00",
			"--nda-acceptance.accepted=true",
			"--nda-acceptance.agreement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nda-acceptance.authority-confirmed=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"notional_amount: '125000.00'\n" +
			"nda_acceptance:\n" +
			"  accepted: true\n" +
			"  agreement_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  authority_confirmed: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:private-markets", "update-ioi",
			"--ioi-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "0",
		)
	})
}
