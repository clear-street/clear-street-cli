// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1AccountsOrdersCancelAllOpenOrders(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:orders", "cancel-all-open-orders",
			"--account-id", "0",
			"--instrument-type", "COMMON_STOCK",
			"--security-id", "string",
			"--security-id-source", "string",
			"--side", "BUY",
			"--type", "MARKET",
		)
	})
}

func TestActiveV1AccountsOrdersCancelOpenOrder(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:orders", "cancel-open-order",
			"--account-id", "0",
			"--order-id", "order_id",
		)
	})
}

func TestActiveV1AccountsOrdersGetOrderByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:orders", "get-order-by-id",
			"--account-id", "0",
			"--order-id", "order_id",
		)
	})
}

func TestActiveV1AccountsOrdersGetOrders(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:orders", "get-orders",
			"--account-id", "0",
			"--from", "'2019-12-27T18:11:19.117Z'",
			"--instrument-type", "COMMON_STOCK",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--security-id", "string",
			"--security-id-source", "string",
			"--status", "PENDING_NEW",
			"--symbol", "symbol",
			"--to", "'2019-12-27T18:11:19.117Z'",
			"--underlying-instrument-ids", "underlying_instrument_ids",
		)
	})
}

func TestActiveV1AccountsOrdersReplaceOrder(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:orders", "replace-order",
			"--account-id", "0",
			"--order-id", "order_id",
			"--limit-price", "150.50",
			"--quantity", "125",
			"--stop-price", "148.00",
			"--time-in-force", "DAY",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"limit_price: '150.50'\n" +
			"quantity: '125'\n" +
			"stop_price: '148.00'\n" +
			"time_in_force: DAY\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:accounts:orders", "replace-order",
			"--account-id", "0",
			"--order-id", "order_id",
		)
	})
}

func TestActiveV1AccountsOrdersSubmitOrders(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:orders", "submit-orders",
			"--account-id", "0",
			"--order", "{legs: [{instrument_type: OPTION, ratio: ratio, security: 0193bb84-447a-706f-996f-097254663f02, side: BUY, id: '1', position_effect: OPEN}, {instrument_type: OPTION, ratio: ratio, security: 0193bb84-4db4-78ec-b4fd-cba8be61cf8a, side: SELL, id: '2', position_effect: OPEN}, {instrument_type: OPTION, ratio: ratio, security: 0193bb84-5264-7f20-8fd3-35df82cd6ef0, side: BUY, id: '3', position_effect: OPEN}], order_type: LIMIT, time_in_force: DAY, id: my-mleg-ref-20251001-001, limit_price: '0.50', quantity: '1'}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- legs:\n" +
			"    - instrument_type: OPTION\n" +
			"      ratio: ratio\n" +
			"      security: 0193bb84-447a-706f-996f-097254663f02\n" +
			"      side: BUY\n" +
			"      id: '1'\n" +
			"      position_effect: OPEN\n" +
			"    - instrument_type: OPTION\n" +
			"      ratio: ratio\n" +
			"      security: 0193bb84-4db4-78ec-b4fd-cba8be61cf8a\n" +
			"      side: SELL\n" +
			"      id: '2'\n" +
			"      position_effect: OPEN\n" +
			"    - instrument_type: OPTION\n" +
			"      ratio: ratio\n" +
			"      security: 0193bb84-5264-7f20-8fd3-35df82cd6ef0\n" +
			"      side: BUY\n" +
			"      id: '3'\n" +
			"      position_effect: OPEN\n" +
			"  order_type: LIMIT\n" +
			"  time_in_force: DAY\n" +
			"  id: my-mleg-ref-20251001-001\n" +
			"  limit_price: '0.50'\n" +
			"  quantity: '1'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:accounts:orders", "submit-orders",
			"--account-id", "0",
		)
	})
}
