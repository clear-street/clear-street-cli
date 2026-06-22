// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1OrdersCancelAllOpenOrders(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "cancel-all-open-orders",
			"--account-id", "0",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--instrument-type", "COMMON_STOCK",
			"--side", "BUY",
			"--type", "MARKET",
		)
	})
}

func TestV1OrdersCancelOpenOrder(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "cancel-open-order",
			"--account-id", "0",
			"--order-id", "order_id",
		)
	})
}

func TestV1OrdersGetExecutions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "get-executions",
			"--account-id", "0",
			"--from", "'2019-12-27T18:11:19.117Z'",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--to", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

func TestV1OrdersGetOrderByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "get-order-by-id",
			"--account-id", "0",
			"--order-id", "order_id",
		)
	})
}

func TestV1OrdersGetOrders(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "get-orders",
			"--account-id", "0",
			"--from", "'2019-12-27T18:11:19.117Z'",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--instrument-type", "COMMON_STOCK",
			"--order-id", "string",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--status", "PENDING_NEW",
			"--symbol", "symbol",
			"--to", "'2019-12-27T18:11:19.117Z'",
			"--underlying-instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1OrdersReplaceOrder(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "replace-order",
			"--account-id", "0",
			"--order-id", "order_id",
			"--limit-price", "49.00",
			"--quantity", "1",
			"--stop-price", "52.00",
			"--time-in-force", "DAY",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"limit_price: '49.00'\n" +
			"quantity: '1'\n" +
			"stop_price: '52.00'\n" +
			"time_in_force: DAY\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:orders", "replace-order",
			"--account-id", "0",
			"--order-id", "order_id",
		)
	})
}

func TestV1OrdersSubmitOrders(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "submit-orders",
			"--account-id", "0",
			"--order", "{order_type: LIMIT, quantity: '1', side: BUY, time_in_force: DAY, id: my-ref-id-20251001-002, expires_at: '2025-10-15T16:00:00.000000000Z', extended_hours: true, instrument_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, limit_offset: '0.50', limit_price: '48.00', position_effect: OPEN, stop_price: '52.00', symbol: TSLA, trailing_offset: '2.00', trailing_offset_type: PRICE}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- order_type: LIMIT\n" +
			"  quantity: '1'\n" +
			"  side: BUY\n" +
			"  time_in_force: DAY\n" +
			"  id: my-ref-id-20251001-002\n" +
			"  expires_at: '2025-10-15T16:00:00.000000000Z'\n" +
			"  extended_hours: true\n" +
			"  instrument_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  limit_offset: '0.50'\n" +
			"  limit_price: '48.00'\n" +
			"  position_effect: OPEN\n" +
			"  stop_price: '52.00'\n" +
			"  symbol: TSLA\n" +
			"  trailing_offset: '2.00'\n" +
			"  trailing_offset_type: PRICE\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:orders", "submit-orders",
			"--account-id", "0",
		)
	})
}
