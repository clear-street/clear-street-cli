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

var v1InstrumentDataGetAllInstrumentEvents = cli.Command{
	Name:    "get-all-instrument-events",
	Usage:   "List instrument events across all securities.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "event-type",
			Usage:     "Filter by event type(s). Comma-delimited list. Example: `event_types=EARNINGS,IPO`.",
			QueryPath: "event_types",
		},
		&requestflag.Flag[string]{
			Name:      "from-date",
			Usage:     "The start date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "from_date",
		},
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Filter by OEMS instrument ID(s). Comma-delimited list of UUIDs. Example: `instrument_ids=550e8400-e29b-41d4-a716-446655440000`.",
			QueryPath: "instrument_ids",
		},
		&requestflag.Flag[string]{
			Name:      "to-date",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "to_date",
		},
	},
	Action:          handleV1InstrumentDataGetAllInstrumentEvents,
	HideHelpCommand: true,
}

var v1InstrumentDataGetInstrumentAnalystConsensus = cli.Command{
	Name:    "get-instrument-analyst-consensus",
	Usage:   "Retrieves analyst ratings and price targets for an instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "The start date for the query range, inclusive (YYYY-MM-DD)",
			QueryPath: "from",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD)",
			QueryPath: "to",
		},
	},
	Action:          handleV1InstrumentDataGetInstrumentAnalystConsensus,
	HideHelpCommand: true,
}

var v1InstrumentDataGetInstrumentBalanceSheetStatements = cli.Command{
	Name:    "get-instrument-balance-sheet-statements",
	Usage:   "Get balance sheet statements for an instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[string]{
			Name:      "from-date",
			Usage:     "The start date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "from_date",
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
		&requestflag.Flag[string]{
			Name:      "to-date",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "to_date",
		},
	},
	Action:          handleV1InstrumentDataGetInstrumentBalanceSheetStatements,
	HideHelpCommand: true,
}

var v1InstrumentDataGetInstrumentCashFlowStatements = cli.Command{
	Name:    "get-instrument-cash-flow-statements",
	Usage:   "Get cash flow statements for an instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[string]{
			Name:      "from-date",
			Usage:     "The start date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "from_date",
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
		&requestflag.Flag[string]{
			Name:      "to-date",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "to_date",
		},
	},
	Action:          handleV1InstrumentDataGetInstrumentCashFlowStatements,
	HideHelpCommand: true,
}

var v1InstrumentDataGetInstrumentEvents = cli.Command{
	Name:    "get-instrument-events",
	Usage:   "Retrieves corporate events (dividends, splits, etc.) for an instrument, grouped\nby event type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[string]{
			Name:      "from-date",
			Usage:     "The start date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "from_date",
		},
		&requestflag.Flag[string]{
			Name:      "to-date",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "to_date",
		},
	},
	Action:          handleV1InstrumentDataGetInstrumentEvents,
	HideHelpCommand: true,
}

var v1InstrumentDataGetInstrumentFundamentals = cli.Command{
	Name:    "get-instrument-fundamentals",
	Usage:   "Retrieves supplemental fundamentals and company profile data for an instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
	},
	Action:          handleV1InstrumentDataGetInstrumentFundamentals,
	HideHelpCommand: true,
}

var v1InstrumentDataGetInstrumentIncomeStatements = cli.Command{
	Name:    "get-instrument-income-statements",
	Usage:   "Retrieves quarterly income statements for a specific instrument, sorted by\nfiscal period (most recent first).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[string]{
			Name:      "from-date",
			Usage:     "The start date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "from_date",
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
		&requestflag.Flag[string]{
			Name:      "to-date",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "to_date",
		},
	},
	Action:          handleV1InstrumentDataGetInstrumentIncomeStatements,
	HideHelpCommand: true,
}

func handleV1InstrumentDataGetAllInstrumentEvents(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := clearstreet.V1InstrumentDataGetAllInstrumentEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.GetAllInstrumentEvents(ctx, params, options...)
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
		Title:          "v1:instrument-data get-all-instrument-events",
		Transform:      transform,
	})
}

func handleV1InstrumentDataGetInstrumentAnalystConsensus(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instrument-id") && len(unusedArgs) > 0 {
		cmd.Set("instrument-id", unusedArgs[0])
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

	params := clearstreet.V1InstrumentDataGetInstrumentAnalystConsensusParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.GetInstrumentAnalystConsensus(
		ctx,
		clearstreet.InstrumentIDOrSymbol(cmd.Value("instrument-id").(string)),
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
		Title:          "v1:instrument-data get-instrument-analyst-consensus",
		Transform:      transform,
	})
}

func handleV1InstrumentDataGetInstrumentBalanceSheetStatements(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instrument-id") && len(unusedArgs) > 0 {
		cmd.Set("instrument-id", unusedArgs[0])
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

	params := clearstreet.V1InstrumentDataGetInstrumentBalanceSheetStatementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.GetInstrumentBalanceSheetStatements(
		ctx,
		clearstreet.InstrumentIDOrSymbol(cmd.Value("instrument-id").(string)),
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
		Title:          "v1:instrument-data get-instrument-balance-sheet-statements",
		Transform:      transform,
	})
}

func handleV1InstrumentDataGetInstrumentCashFlowStatements(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instrument-id") && len(unusedArgs) > 0 {
		cmd.Set("instrument-id", unusedArgs[0])
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

	params := clearstreet.V1InstrumentDataGetInstrumentCashFlowStatementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.GetInstrumentCashFlowStatements(
		ctx,
		clearstreet.InstrumentIDOrSymbol(cmd.Value("instrument-id").(string)),
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
		Title:          "v1:instrument-data get-instrument-cash-flow-statements",
		Transform:      transform,
	})
}

func handleV1InstrumentDataGetInstrumentEvents(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instrument-id") && len(unusedArgs) > 0 {
		cmd.Set("instrument-id", unusedArgs[0])
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

	params := clearstreet.V1InstrumentDataGetInstrumentEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.GetInstrumentEvents(
		ctx,
		clearstreet.InstrumentIDOrSymbol(cmd.Value("instrument-id").(string)),
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
		Title:          "v1:instrument-data get-instrument-events",
		Transform:      transform,
	})
}

func handleV1InstrumentDataGetInstrumentFundamentals(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instrument-id") && len(unusedArgs) > 0 {
		cmd.Set("instrument-id", unusedArgs[0])
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.GetInstrumentFundamentals(ctx, clearstreet.InstrumentIDOrSymbol(cmd.Value("instrument-id").(string)), options...)
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
		Title:          "v1:instrument-data get-instrument-fundamentals",
		Transform:      transform,
	})
}

func handleV1InstrumentDataGetInstrumentIncomeStatements(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instrument-id") && len(unusedArgs) > 0 {
		cmd.Set("instrument-id", unusedArgs[0])
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

	params := clearstreet.V1InstrumentDataGetInstrumentIncomeStatementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.GetInstrumentIncomeStatements(
		ctx,
		clearstreet.InstrumentIDOrSymbol(cmd.Value("instrument-id").(string)),
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
		Title:          "v1:instrument-data get-instrument-income-statements",
		Transform:      transform,
	})
}
