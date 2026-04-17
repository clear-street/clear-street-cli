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

var activeV1AccountsOrdersCancelAllOrders = cli.Command{
	Name:    "cancel-all-orders",
	Usage:   "All filter parameters can be used independently or combined. The only constraint\nis that `security_id` and `security_id_source` must be provided together if\neither is specified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:      "security-id",
			Usage:     "Filter by security ID(s). Accepts single value or indexed array.\n\nExamples:\n- Single: `security_id=037833100`\n- Multiple: `security_id[0]=037833100&security_id[1]=594918104`",
			QueryPath: "security_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "security-id-source",
			Usage:     "Source(s) for the security ID filter. Must match the count and order of security_id.\n\nExamples:\n- Single: `security_id_source=CUSIP`\n- Multiple: `security_id_source[0]=CUSIP&security_id_source[1]=FIGI`",
			QueryPath: "security_id_source",
		},
		&requestflag.Flag[string]{
			Name:      "security-type",
			Usage:     "Filter by security type (e.g., COMMON_STOCK, OPTION)",
			QueryPath: "security_type",
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
	Action:          handleActiveV1AccountsOrdersCancelAllOrders,
	HideHelpCommand: true,
}

var activeV1AccountsOrdersCancelOrder = cli.Command{
	Name:    "cancel-order",
	Usage:   "Cancel a specific order",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "order-id",
			Required: true,
		},
	},
	Action:          handleActiveV1AccountsOrdersCancelOrder,
	HideHelpCommand: true,
}

var activeV1AccountsOrdersGetOrderByID = cli.Command{
	Name:    "get-order-by-id",
	Usage:   "Get order by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "order-id",
			Required: true,
		},
	},
	Action:          handleActiveV1AccountsOrdersGetOrderByID,
	HideHelpCommand: true,
}

var activeV1AccountsOrdersGetOrders = cli.Command{
	Name:    "get-orders",
	Usage:   "List orders for an account with optional filtering",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "The start date and time for the query range, inclusive (ISO 8601 format)",
			QueryPath: "from",
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
			Name:      "security-id",
			Usage:     "Filter by security ID(s). Accepts single value or indexed array.\n\nExamples:\n- Single: `security_id=037833100`\n- Multiple: `security_id[0]=037833100&security_id[1]=594918104`",
			QueryPath: "security_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "security-id-source",
			Usage:     "Source(s) for the security ID filter. Must match the count and order of security_id.\n\nExamples:\n- Single: `security_id_source=CUSIP`\n- Multiple: `security_id_source[0]=CUSIP&security_id_source[1]=FIGI`",
			QueryPath: "security_id_source",
		},
		&requestflag.Flag[string]{
			Name:      "security-type",
			Usage:     "Security type filter (e.g., COMMON_STOCK, PREFERRED_STOCK)",
			QueryPath: "security_type",
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
	},
	Action:          handleActiveV1AccountsOrdersGetOrders,
	HideHelpCommand: true,
}

var activeV1AccountsOrdersReplaceOrder = cli.Command{
	Name:    "replace-order",
	Usage:   "Replace an order with new parameters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "order-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "limit-price",
			Usage:    "New limit price for the order",
			BodyPath: "limit_price",
		},
		&requestflag.Flag[any]{
			Name:     "quantity",
			Usage:    "New quantity for the order",
			BodyPath: "quantity",
		},
		&requestflag.Flag[any]{
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
	Action:          handleActiveV1AccountsOrdersReplaceOrder,
	HideHelpCommand: true,
}

var activeV1AccountsOrdersSubmitOrders = cli.Command{
	Name:    "submit-orders",
	Usage:   "Submit new orders",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleActiveV1AccountsOrdersSubmitOrders,
	HideHelpCommand: true,
}

func handleActiveV1AccountsOrdersCancelAllOrders(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountOrderCancelAllOrdersParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.Accounts.Orders.CancelAllOrders(
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
		Title:          "active:v1:accounts:orders cancel-all-orders",
		Transform:      transform,
	})
}

func handleActiveV1AccountsOrdersCancelOrder(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountOrderCancelOrderParams{
		AccountID: cmd.Value("account-id").(int64),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.Accounts.Orders.CancelOrder(
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
		Title:          "active:v1:accounts:orders cancel-order",
		Transform:      transform,
	})
}

func handleActiveV1AccountsOrdersGetOrderByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountOrderGetOrderByIDParams{
		AccountID: cmd.Value("account-id").(int64),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.Accounts.Orders.GetOrderByID(
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
		Title:          "active:v1:accounts:orders get-order-by-id",
		Transform:      transform,
	})
}

func handleActiveV1AccountsOrdersGetOrders(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountOrderGetOrdersParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.Accounts.Orders.GetOrders(
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
		Title:          "active:v1:accounts:orders get-orders",
		Transform:      transform,
	})
}

func handleActiveV1AccountsOrdersReplaceOrder(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountOrderReplaceOrderParams{
		AccountID: cmd.Value("account-id").(int64),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.Accounts.Orders.ReplaceOrder(
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
		Title:          "active:v1:accounts:orders replace-order",
		Transform:      transform,
	})
}

func handleActiveV1AccountsOrdersSubmitOrders(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountOrderSubmitOrdersParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.Accounts.Orders.SubmitOrders(
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
		Title:          "active:v1:accounts:orders submit-orders",
		Transform:      transform,
	})
}
