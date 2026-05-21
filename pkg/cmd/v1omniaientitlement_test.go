// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1OmniAIEntitlementsCreateEntitlements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:entitlements", "create-entitlements",
			"--account-id", "100019",
			"--account-id", "100021",
			"--agreement-id", "01JZ0000000000000000000000",
			"--entitlement-code", "omni.account_data",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_ids:\n" +
			"  - 100019\n" +
			"  - 100021\n" +
			"agreement_id: 01JZ0000000000000000000000\n" +
			"entitlement_codes:\n" +
			"  - omni.account_data\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:omni-ai:entitlements", "create-entitlements",
		)
	})
}

func TestV1OmniAIEntitlementsDeleteEntitlement(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:entitlements", "delete-entitlement",
			"--entitlement-id", "entitlement_id",
		)
	})
}

func TestV1OmniAIEntitlementsGetEntitlementAgreements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:entitlements", "get-entitlement-agreements",
		)
	})
}

func TestV1OmniAIEntitlementsGetEntitlements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:omni-ai:entitlements", "get-entitlements",
			"--trading-account-id", "0",
		)
	})
}
