// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/clear-street/clear-street-cli/internal/apiquery"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
	"github.com/clear-street/clear-street-go"
	"github.com/clear-street/clear-street-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var v1OrdersCancelAllOpenOrders = cli.Command{
	Name:    "cancel-all-open-orders",
	Usage:   "Cancel all orders for an account",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-separated instrument IDs (UUID) or symbols (equity tickers or OSI option symbols).",
			QueryPath: "instrument_ids",
		},
		&requestflag.Flag[string]{
			Name:      "instrument-type",
			Usage:     "Filter by instrument type (e.g., COMMON_STOCK, OPTION)",
			QueryPath: "instrument_type",
		},
		&requestflag.Flag[string]{
			Name:      "side",
			Usage:     "Filter by order side (BUY or SELL)",
			QueryPath: "side",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter by order type (e.g., MARKET, LIMIT)",
			QueryPath: "type",
		},
	},
	Action:          handleV1OrdersCancelAllOpenOrders,
	HideHelpCommand: true,
}

var v1OrdersCancelOpenOrder = cli.Command{
	Name:    "cancel-open-order",
	Usage:   "Cancel a specific order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "order-id",
			Required:  true,
			PathParam: "order_id",
		},
	},
	Action:          handleV1OrdersCancelOpenOrder,
	HideHelpCommand: true,
}

var v1OrdersGetExecutions = cli.Command{
	Name:    "get-executions",
	Usage:   "Retrieves filled and partially-filled execution reports for the specified\ntrading account, ordered by transaction time (nanosecond precision) descending.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "The start date and time for the query range, inclusive (ISO 8601 format)",
			QueryPath: "from",
		},
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-separated instrument identifiers (UUIDs) or symbols (e.g. `AAPL`) to filter by. When provided, only executions for any of the listed instruments are returned.",
			QueryPath: "instrument_ids",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   1000,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "The end date and time for the query range, inclusive (ISO 8601 format)",
			QueryPath: "to",
		},
	},
	Action:          handleV1OrdersGetExecutions,
	HideHelpCommand: true,
}

var v1OrdersGetOrderByID = cli.Command{
	Name:    "get-order-by-id",
	Usage:   "Fetch a single order. The `{order_id}` path parameter accepts either the order's\n`id` or its `client_order_id`. A `client_order_id` can only be used while the\norder is open; after that, use the `id` returned in every order response, or\nfind the order with the list-orders endpoint's `order_ids` filter, which accepts\nboth identifiers at any time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "order-id",
			Required:  true,
			PathParam: "order_id",
		},
	},
	Action:          handleV1OrdersGetOrderByID,
	HideHelpCommand: true,
}

var v1OrdersGetOrders = cli.Command{
	Name:    "get-orders",
	Usage:   "List orders for an account with optional filtering",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "The start date and time for the query range, inclusive (ISO 8601 format)",
			QueryPath: "from",
		},
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-separated instrument IDs (UUID) or symbols (equity tickers or OSI option symbols).",
			QueryPath: "instrument_ids",
		},
		&requestflag.Flag[string]{
			Name:      "instrument-type",
			Usage:     "Instrument type filter (e.g., COMMON_STOCK, OPTION)",
			QueryPath: "instrument_type",
		},
		&requestflag.Flag[[]string]{
			Name:      "order-id",
			Usage:     "Comma-separated list of order identifiers. Each value may be an order's `id` or its `client_order_id`; only orders matching one of the given identifiers are returned.",
			QueryPath: "order_ids",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   1000,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[[]string]{
			Name:      "status",
			Usage:     "Comma-separated order statuses to filter by",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "symbol",
			Usage:     "Filter by symbol",
			QueryPath: "symbol",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "The end date and time for the query range, inclusive (ISO 8601 format)",
			QueryPath: "to",
		},
		&requestflag.Flag[[]string]{
			Name:      "underlying-instrument-id",
			Usage:     "Comma-separated instrument IDs (UUID) or symbols (equity tickers or OSI option symbols). Matches options orders whose resolved underlier is any of the given instruments.",
			QueryPath: "underlying_instrument_ids",
		},
	},
	Action:          handleV1OrdersGetOrders,
	HideHelpCommand: true,
}

var v1OrdersReplaceOrder = cli.Command{
	Name:    "replace-order",
	Usage:   "Replace an order with new parameters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "order-id",
			Required:  true,
			PathParam: "order_id",
		},
		&requestflag.Flag[*string]{
			Name:     "limit-price",
			Usage:    "New limit price for the order",
			BodyPath: "limit_price",
		},
		&requestflag.Flag[*string]{
			Name:     "quantity",
			Usage:    "New quantity for the order",
			BodyPath: "quantity",
		},
		&requestflag.Flag[*string]{
			Name:     "stop-price",
			Usage:    "New stop price for the order",
			BodyPath: "stop_price",
		},
	},
	Action:          handleV1OrdersReplaceOrder,
	HideHelpCommand: true,
}

var v1OrdersSubmitOrders = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit-orders",
	Usage:   "Submit new orders",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "order",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleV1OrdersSubmitOrders,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"order": {
		&requestflag.InnerFlag[string]{
			Name:       "order.order-type",
			Usage:      "Strict order-type enum for order submission/replacement requests.",
			InnerField: "order_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "order.quantity",
			Usage:      "Quantity to trade. For COMMON_STOCK: shares (may be fractional if supported).\nFor OPTION (single-leg): contracts (must be an integer)",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[string]{
			Name:       "order.side",
			Usage:      "Side of an order",
			InnerField: "side",
		},
		&requestflag.InnerFlag[string]{
			Name:       "order.time-in-force",
			Usage:      "Strict time-in-force enum for order submission requests.",
			InnerField: "time_in_force",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.id",
			Usage:      "Optional client-provided unique ID (idempotency). Required to be unique per account.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "order.expires-at",
			Usage:      "The timestamp when the order should expire (UTC). Required when time_in_force is GOOD_TILL_DATE.",
			InnerField: "expires_at",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "order.extended-hours",
			Usage:      "Allow trading outside regular trading hours. Some brokers disallow options outside RTH.",
			InnerField: "extended_hours",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.instrument-id",
			Usage:      "Instrument identifier: either an instrument UUID or a symbol (symbol for equities, OSI for options). Non-UUID inputs are resolved server-side.",
			InnerField: "instrument_id",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.limit-offset",
			Usage:      "Limit offset for trailing stop-limit orders (signed)",
			InnerField: "limit_offset",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.limit-price",
			Usage:      "Limit price (required for LIMIT and STOP_LIMIT orders)",
			InnerField: "limit_price",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.stop-price",
			Usage:      "Stop price (required for STOP and STOP_LIMIT orders)",
			InnerField: "stop_price",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.symbol",
			Usage:      "Trading symbol. For equities, use the ticker symbol (e.g., \"TSLA\").\nFor options, use the OSI symbol (e.g., \"TSLA  250117C00190000\").\nEither `symbol` or `instrument_id` must be provided.",
			InnerField: "symbol",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.trailing-offset",
			Usage:      "Trailing offset amount (required for trailing orders)",
			InnerField: "trailing_offset",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "order.trailing-offset-type",
			Usage:      "Trailing offset type for trailing stop orders.",
			InnerField: "trailing_offset_type",
		},
	},
})

func handleV1OrdersCancelAllOpenOrders(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1OrderCancelAllOpenOrdersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Orders.CancelAllOpenOrders(
		ctx,
		cmd.Value("account-id").(int64),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:orders cancel-all-open-orders",
		Transform:      transform,
	})
}

func handleV1OrdersCancelOpenOrder(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1OrderCancelOpenOrderParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Orders.CancelOpenOrder(
		ctx,
		cmd.Value("order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:orders cancel-open-order",
		Transform:      transform,
	})
}

func handleV1OrdersGetExecutions(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1OrderGetExecutionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Orders.GetExecutions(
		ctx,
		cmd.Value("account-id").(int64),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:orders get-executions",
		Transform:      transform,
	})
}

func handleV1OrdersGetOrderByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1OrderGetOrderByIDParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Orders.GetOrderByID(
		ctx,
		cmd.Value("order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:orders get-order-by-id",
		Transform:      transform,
	})
}

func handleV1OrdersGetOrders(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1OrderGetOrdersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Orders.GetOrders(
		ctx,
		cmd.Value("account-id").(int64),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:orders get-orders",
		Transform:      transform,
	})
}

func handleV1OrdersReplaceOrder(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1OrderReplaceOrderParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Orders.ReplaceOrder(
		ctx,
		cmd.Value("order-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:orders replace-order",
		Transform:      transform,
	})
}

func handleV1OrdersSubmitOrders(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1OrderSubmitOrdersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Orders.SubmitOrders(
		ctx,
		cmd.Value("account-id").(int64),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:orders submit-orders",
		Transform:      transform,
	})
}
