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
			Usage:     "OEMS instrument UUID",
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
			Usage:     "Comma-separated OEMS instrument UUIDs",
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
		&requestflag.Flag[string]{
			Name:      "contract-type",
			Usage:     "The type of options contract",
			QueryPath: "contract_type",
		},
		&requestflag.Flag[any]{
			Name:      "expiry",
			Usage:     "Filter to contracts expiring on this date (YYYY-MM-DD)",
			QueryPath: "expiry",
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
			Name:      "underlier",
			Usage:     "Underlier symbol (e.g., AAPL, SPX)",
			QueryPath: "underlier",
		},
		&requestflag.Flag[string]{
			Name:      "underlying-instrument-id",
			Usage:     "OEMS instrument UUID",
			QueryPath: "underlying_instrument_id",
		},
	},
	Action:          handleV1InstrumentsGetOptionContracts,
	HideHelpCommand: true,
}

var v1InstrumentsSearchInstruments = cli.Command{
	Name:    "search-instruments",
	Usage:   "Search instruments by symbol, alternate identifier, or company name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Search term applied case-insensitively to ticker symbols, alternate identifiers (CUSIP, ISIN, OPRA root, CMS), and company names for non-option instruments. Option searches match symbols and alternate identifiers.",
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
		apiquery.ArrayQueryFormatIndices,
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
		apiquery.ArrayQueryFormatIndices,
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
		apiquery.ArrayQueryFormatIndices,
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

func handleV1InstrumentsSearchInstruments(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
