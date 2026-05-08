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
			"--agreement-id", "01JZ0000000000000000000000",
			"--requested-entitlement-code", "omni.account_data",
			"--trading-account-id", "100019",
			"--trading-account-id", "100021",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agreement_id: 01JZ0000000000000000000000\n" +
			"requested_entitlement_codes:\n" +
			"  - omni.account_data\n" +
			"trading_account_ids:\n" +
			"  - 100019\n" +
			"  - 100021\n")
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
