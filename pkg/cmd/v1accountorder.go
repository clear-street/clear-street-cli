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

var v1AccountsOrdersCancelAllOpenOrders = cli.Command{
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
			Usage:     "Comma-separated OEMS instrument UUIDs",
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
	Action:          handleV1AccountsOrdersCancelAllOpenOrders,
	HideHelpCommand: true,
}

var v1AccountsOrdersCancelOpenOrder = cli.Command{
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
	Action:          handleV1AccountsOrdersCancelOpenOrder,
	HideHelpCommand: true,
}

var v1AccountsOrdersGetOrderByID = cli.Command{
	Name:    "get-order-by-id",
	Usage:   "Get Order By ID",
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
	Action:          handleV1AccountsOrdersGetOrderByID,
	HideHelpCommand: true,
}

var v1AccountsOrdersGetOrders = cli.Command{
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
			Usage:     "Comma-separated OEMS instrument UUIDs",
			QueryPath: "instrument_ids",
		},
		&requestflag.Flag[string]{
			Name:      "instrument-type",
			Usage:     "Instrument type filter (e.g., COMMON_STOCK, OPTION)",
			QueryPath: "instrument_type",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   1000,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next page of results. Contains encoded pagination state (limit + offset).\nWhen provided, page_size is ignored.",
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
		&requestflag.Flag[string]{
			Name:      "underlying-instrument-ids",
			Usage:     "Comma-separated OEMS instrument UUIDs. Matches options orders whose resolved underlier is any of the given IDs.",
			QueryPath: "underlying_instrument_ids",
		},
	},
	Action:          handleV1AccountsOrdersGetOrders,
	HideHelpCommand: true,
}

var v1AccountsOrdersReplaceOrder = cli.Command{
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
		&requestflag.Flag[string]{
			Name:     "time-in-force",
			Usage:    "Time in force",
			BodyPath: "time_in_force",
		},
	},
	Action:          handleV1AccountsOrdersReplaceOrder,
	HideHelpCommand: true,
}

var v1AccountsOrdersSubmitOrders = cli.Command{
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
	Action:          handleV1AccountsOrdersSubmitOrders,
	HideHelpCommand: true,
}

func handleV1AccountsOrdersCancelAllOpenOrders(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1AccountOrderCancelAllOpenOrdersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Orders.CancelAllOpenOrders(
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
		Title:          "v1:accounts:orders cancel-all-open-orders",
		Transform:      transform,
	})
}

func handleV1AccountsOrdersCancelOpenOrder(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1AccountOrderCancelOpenOrderParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Orders.CancelOpenOrder(
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
		Title:          "v1:accounts:orders cancel-open-order",
		Transform:      transform,
	})
}

func handleV1AccountsOrdersGetOrderByID(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1AccountOrderGetOrderByIDParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Orders.GetOrderByID(
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
		Title:          "v1:accounts:orders get-order-by-id",
		Transform:      transform,
	})
}

func handleV1AccountsOrdersGetOrders(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1AccountOrderGetOrdersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Orders.GetOrders(
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
		Title:          "v1:accounts:orders get-orders",
		Transform:      transform,
	})
}

func handleV1AccountsOrdersReplaceOrder(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1AccountOrderReplaceOrderParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Orders.ReplaceOrder(
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
		Title:          "v1:accounts:orders replace-order",
		Transform:      transform,
	})
}

func handleV1AccountsOrdersSubmitOrders(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1AccountOrderSubmitOrdersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Orders.SubmitOrders(
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
		Title:          "v1:accounts:orders submit-orders",
		Transform:      transform,
	})
}
