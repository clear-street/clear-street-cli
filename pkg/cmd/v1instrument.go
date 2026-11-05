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

var v1InstrumentsGetInstrumentByID = cli.Command{
	Name:    "get-instrument-by-id",
	Usage:   "Retrieves detailed information for a specific instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "Instrument identifier: either an instrument UUID or a symbol (symbol for equities, OSI for options). Non-UUID inputs are resolved server-side.",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[*bool]{
			Name:      "include-options-expiry-dates",
			Usage:     "When true, include unique options expiry dates for this instrument",
			QueryPath: "include_options_expiry_dates",
		},
	},
	Action:          handleV1InstrumentsGetInstrumentByID,
	HideHelpCommand: true,
}

var v1InstrumentsGetInstruments = cli.Command{
	Name:    "get-instruments",
	Usage:   "Retrieves a list of tradeable instruments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "easy-to-borrow",
			Usage:     "Filter by easy to borrow status",
			QueryPath: "easy_to_borrow",
		},
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-separated instrument identifiers: unique identifiers or symbols (symbol for equities, OSI for options)",
			QueryPath: "instrument_ids",
		},
		&requestflag.Flag[string]{
			Name:      "instrument-type",
			Usage:     "Filter by instrument type (e.g. COMMON_STOCK, OPTION)",
			QueryPath: "instrument_type",
		},
		&requestflag.Flag[bool]{
			Name:      "is-liquidation-only",
			Usage:     "Filter by liquidation only status",
			QueryPath: "is_liquidation_only",
		},
		&requestflag.Flag[bool]{
			Name:      "is-marginable",
			Usage:     "Filter by marginable status",
			QueryPath: "is_marginable",
		},
		&requestflag.Flag[bool]{
			Name:      "is-ptp",
			Usage:     "Filter by publicly traded partnership (PTP) status",
			QueryPath: "is_ptp",
		},
		&requestflag.Flag[bool]{
			Name:      "is-short-prohibited",
			Usage:     "Filter by short prohibited status",
			QueryPath: "is_short_prohibited",
		},
		&requestflag.Flag[bool]{
			Name:      "is-threshold-security",
			Usage:     "Filter by threshold security status",
			QueryPath: "is_threshold_security",
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
	},
	Action:          handleV1InstrumentsGetInstruments,
	HideHelpCommand: true,
}

var v1InstrumentsGetOptionContracts = cli.Command{
	Name:    "get-option-contracts",
	Usage:   "List options contracts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "contract-id",
			Usage:     "Comma-separated contract instrument IDs (UUID) or OSI option symbols to look up directly, bypassing underlier expansion. Mutually exclusive with root/underlying_instrument_id; up to 100 values.",
			QueryPath: "contract_ids",
		},
		&requestflag.Flag[string]{
			Name:      "contract-type",
			Usage:     "Filter by contract type: CALL or PUT",
			QueryPath: "contract_type",
		},
		&requestflag.Flag[any]{
			Name:      "expiry",
			Usage:     "Filter to contracts expiring on this date (YYYY-MM-DD)",
			QueryPath: "expiry",
		},
		&requestflag.Flag[bool]{
			Name:      "is-settle-on-open",
			Usage:     "Filter by settlement cycle: true for early-settling (AM, settle-on-open) contracts, false for normal (PM) contracts. Omit to return both.",
			QueryPath: "is_settle_on_open",
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
			Name:      "root",
			Usage:     "OSI option root symbol (e.g., AAPL, SPXW). Index roots also include their weekly series (e.g., SPX includes SPXW).",
			QueryPath: "root",
		},
		&requestflag.Flag[string]{
			Name:      "underlier",
			Usage:     "Deprecated: use `root`. If both are sent, the values must match.",
			QueryPath: "underlier",
		},
		&requestflag.Flag[string]{
			Name:      "underlying-instrument-id",
			Usage:     "Instrument identifier or symbol of the underlying equity/index",
			QueryPath: "underlying_instrument_id",
		},
	},
	Action:          handleV1InstrumentsGetOptionContracts,
	HideHelpCommand: true,
}

var v1InstrumentsGetPredictionCategories = cli.Command{
	Name:            "get-prediction-categories",
	Usage:           "> **Alpha** — this endpoint is experimental and may change or be removed at any\n> time.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleV1InstrumentsGetPredictionCategories,
	HideHelpCommand: true,
}

var v1InstrumentsGetPredictionEvents = requestflag.WithInnerFlags(cli.Command{
	Name:    "get-prediction-events",
	Usage:   "> **Alpha** — this endpoint is experimental and may change or be removed at any\n> time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "categories",
			Usage:     "Comma-separated categories. Keeps events with a listable market in any of them, ignoring case.",
			QueryPath: "categories",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "close-at",
			QueryPath: "close_at",
		},
		&requestflag.Flag[bool]{
			Name:      "include-markets",
			Usage:     "Embed a summary of every market of each event, in any status. Default false.",
			QueryPath: "include_markets",
		},
		&requestflag.Flag[bool]{
			Name:      "include-settled",
			Usage:     "Also list settled and expired events, not only events with a listable market. Every market then counts as listable for the other filters, `title`, `categories`, and the open and close times. Default false. A request that takes longer than 30 seconds fails with `504`; narrow `search` or the filters and try again.",
			QueryPath: "include_settled",
		},
		&requestflag.Flag[string]{
			Name:      "market-ids",
			Usage:     "Comma-separated market ids (UUID) or tickers, 1-100. Keeps events of those markets, in any status; the other filters still apply. Ids that match no market are reported: `207`, or `404` when none match.",
			QueryPath: "market_ids",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "open-at",
			QueryPath: "open_at",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   20,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Keeps events whose code, series code, or a market's question or title contains every word, matched by prefix and English stem, ignoring case. Searches listable events, or every event with `include_settled`. A word of four or more letters that matches nothing is replaced by up to three close spellings. Without `sort_by`, orders by relevance, closest spellings first.",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "series-codes",
			Usage:     "Comma-separated series codes. Keeps events of those series, ignoring case.",
			QueryPath: "series_codes",
		},
		&requestflag.Flag[string]{
			Name:      "sort-by",
			Usage:     "Sort key. Defaults to relevance when `search` is set, else the event code. Ties break on event code ascending.",
			QueryPath: "sort_by",
		},
		&requestflag.Flag[string]{
			Name:      "sort-direction",
			Usage:     "Sort direction. Defaults to ascending. Ignored when sorting by relevance.",
			QueryPath: "sort_direction",
		},
	},
	Action:          handleV1InstrumentsGetPredictionEvents,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"close-at": {
		&requestflag.InnerFlag[string]{
			Name:       "close-at.gt",
			Usage:      "Return only rows where `close_at` is strictly after the given value. A bare `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches from the start of the following day. See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "gt",
		},
		&requestflag.InnerFlag[string]{
			Name:       "close-at.gte",
			Usage:      "Return only rows where `close_at` is on or after the given value. A bare `YYYY-MM-DD` date expands to the start of that day (UTC). See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[string]{
			Name:       "close-at.lt",
			Usage:      "Return only rows where `close_at` is strictly before the given value. A bare `YYYY-MM-DD` date expands to the start of that day (UTC). See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "lt",
		},
		&requestflag.InnerFlag[string]{
			Name:       "close-at.lte",
			Usage:      "Return only rows where `close_at` is on or before the given value. A bare `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches through the end of that day. See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "lte",
		},
	},
	"open-at": {
		&requestflag.InnerFlag[string]{
			Name:       "open-at.gt",
			Usage:      "Return only rows where `open_at` is strictly after the given value. A bare `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches from the start of the following day. See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "gt",
		},
		&requestflag.InnerFlag[string]{
			Name:       "open-at.gte",
			Usage:      "Return only rows where `open_at` is on or after the given value. A bare `YYYY-MM-DD` date expands to the start of that day (UTC). See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[string]{
			Name:       "open-at.lt",
			Usage:      "Return only rows where `open_at` is strictly before the given value. A bare `YYYY-MM-DD` date expands to the start of that day (UTC). See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "lt",
		},
		&requestflag.InnerFlag[string]{
			Name:       "open-at.lte",
			Usage:      "Return only rows where `open_at` is on or before the given value. A bare `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches through the end of that day. See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "lte",
		},
	},
})

var v1InstrumentsGetPredictionMarkets = cli.Command{
	Name:    "get-prediction-markets",
	Usage:   "> **Alpha** — this endpoint is experimental and may change or be removed at any\n> time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "event-code",
			Usage:     "Event code; returns the markets of the event. Mutually exclusive with `market_ids`.",
			QueryPath: "event_code",
		},
		&requestflag.Flag[string]{
			Name:      "market-ids",
			Usage:     "Comma-separated market ids (UUID) or tickers, 1-100, in any status. Mutually exclusive with `event_code`.",
			QueryPath: "market_ids",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   100,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[string]{
			Name:      "statuses",
			Usage:     "Comma-separated statuses to keep: `OPEN`, `UNOPENED`, `PAUSED`, `CLOSED`, `SETTLED`. Defaults to `OPEN` with `event_code` and to every status with `market_ids`. Unrecognized values are ignored.",
			QueryPath: "statuses",
		},
	},
	Action:          handleV1InstrumentsGetPredictionMarkets,
	HideHelpCommand: true,
}

var v1InstrumentsGetPredictionSeries = cli.Command{
	Name:    "get-prediction-series",
	Usage:   "> **Alpha** — this endpoint is experimental and may change or be removed at any\n> time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "categories",
			Usage:     "Comma-separated categories. Keeps series with a market in any of them, ignoring case.",
			QueryPath: "categories",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   100,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
	},
	Action:          handleV1InstrumentsGetPredictionSeries,
	HideHelpCommand: true,
}

var v1InstrumentsSearchInstruments = cli.Command{
	Name:    "search-instruments",
	Usage:   "Search instruments by symbol, alternate identifier, or company name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Search term applied case-insensitively to ticker symbols, alternate identifiers (CUSIP, ISIN, OPRA root), and company names for non-option instruments. Option searches match symbols and alternate identifiers.",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "asset-class",
			Usage:     "Comma-separated asset classes (EQUITY|OPTION|WARRANT|BOND|FX|OTHER). Defaults to EQUITY.",
			QueryPath: "asset_class",
		},
		&requestflag.Flag[string]{
			Name:      "country",
			Usage:     "Optional listing-country filter (e.g., US).",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "currency",
			Usage:     "Optional ISO currency filter (e.g., USD).",
			QueryPath: "currency",
		},
		&requestflag.Flag[bool]{
			Name:      "include-inactive",
			Usage:     "Include inactive instruments. Default false.",
			QueryPath: "include_inactive",
		},
		&requestflag.Flag[bool]{
			Name:      "include-ptp",
			Usage:     "Include publicly traded partnership (PTP) instruments. Default true (penalized in ranking).",
			QueryPath: "include_ptp",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   100,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
	},
	Action:          handleV1InstrumentsSearchInstruments,
	HideHelpCommand: true,
}

func handleV1InstrumentsGetInstrumentByID(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentGetInstrumentByIDParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.GetInstrumentByID(
		ctx,
		cmd.Value("instrument-id").(string),
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
		Title:          "v1:instruments get-instrument-by-id",
		Transform:      transform,
	})
}

func handleV1InstrumentsGetInstruments(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentGetInstrumentsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.GetInstruments(ctx, params, options...)
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
		Title:          "v1:instruments get-instruments",
		Transform:      transform,
	})
}

func handleV1InstrumentsGetOptionContracts(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentGetOptionContractsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.GetOptionContracts(ctx, params, options...)
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
		Title:          "v1:instruments get-option-contracts",
		Transform:      transform,
	})
}

func handleV1InstrumentsGetPredictionCategories(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.GetPredictionCategories(ctx, options...)
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
		Title:          "v1:instruments get-prediction-categories",
		Transform:      transform,
	})
}

func handleV1InstrumentsGetPredictionEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentGetPredictionEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.GetPredictionEvents(ctx, params, options...)
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
		Title:          "v1:instruments get-prediction-events",
		Transform:      transform,
	})
}

func handleV1InstrumentsGetPredictionMarkets(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentGetPredictionMarketsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.GetPredictionMarkets(ctx, params, options...)
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
		Title:          "v1:instruments get-prediction-markets",
		Transform:      transform,
	})
}

func handleV1InstrumentsGetPredictionSeries(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentGetPredictionSeriesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.GetPredictionSeries(ctx, params, options...)
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
		Title:          "v1:instruments get-prediction-series",
		Transform:      transform,
	})
}

func handleV1InstrumentsSearchInstruments(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentSearchInstrumentsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.SearchInstruments(ctx, params, options...)
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
		Title:          "v1:instruments search-instruments",
		Transform:      transform,
	})
}
