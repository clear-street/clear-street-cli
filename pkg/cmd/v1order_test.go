// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1OrdersCancelAllOpenOrders(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "cancel-all-open-orders",
			"--account-id", "0",
			"--instrument-id", "x",
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
			"--instrument-id", "x",
			"--order-id", "string",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--to", "'2019-12-27T18:11:19.117Z'",
			"--underlying-instrument-id", "x",
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
			"--instrument-id", "x",
			"--instrument-type", "COMMON_STOCK",
			"--order-id", "string",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--status", "PENDING_NEW",
			"--symbol", "symbol",
			"--to", "'2019-12-27T18:11:19.117Z'",
			"--underlying-instrument-id", "x",
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
			"--limit-offset", "0.50",
			"--limit-price", "49.00",
			"--quantity", "1",
			"--stop-price", "52.00",
			"--trailing-offset", "2.00",
			"--trailing-offset-type", "PRICE",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"limit_offset: '0.50'\n" +
			"limit_price: '49.00'\n" +
			"quantity: '1'\n" +
			"stop_price: '52.00'\n" +
			"trailing_offset: '2.00'\n" +
			"trailing_offset_type: PRICE\n")
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
			"--order", "{order_type: LIMIT, quantity: '1', side: BUY, time_in_force: DAY, id: my-ref-id-20251001-002, expires_at: '2025-10-15T16:00:00.000000000Z', extended_hours: true, instrument_id: x, limit_offset: '0.50', limit_price: '48.00', stop_price: '52.00', symbol: TSLA, trailing_offset: '2.00', trailing_offset_type: PRICE}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1OrdersSubmitOrders)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:orders", "submit-orders",
			"--account-id", "0",
			"--order.order-type", "LIMIT",
			"--order.quantity", "1",
			"--order.side", "BUY",
			"--order.time-in-force", "DAY",
			"--order.id", "my-ref-id-20251001-002",
			"--order.expires-at", "2025-10-15T16:00:00.000000000Z",
			"--order.extended-hours=true",
			"--order.instrument-id", "x",
			"--order.limit-offset", "0.50",
			"--order.limit-price", "48.00",
			"--order.stop-price", "52.00",
			"--order.symbol", "TSLA",
			"--order.trailing-offset", "2.00",
			"--order.trailing-offset-type", "PRICE",
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
			"  instrument_id: x\n" +
			"  limit_offset: '0.50'\n" +
			"  limit_price: '48.00'\n" +
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
