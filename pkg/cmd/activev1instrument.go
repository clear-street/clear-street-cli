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

var activeV1InstrumentsGetInstrumentByID = cli.Command{
	Name:    "get-instrument-by-id",
	Usage:   "Retrieves detailed information for a specific instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "security-id-source",
			Usage:    "Security identifier source",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "security-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "include-options-expiry-dates",
			Usage:     "When true, include unique options expiry dates for this instrument",
			QueryPath: "include_options_expiry_dates",
		},
	},
	Action:          handleActiveV1InstrumentsGetInstrumentByID,
	HideHelpCommand: true,
}

var activeV1InstrumentsGetInstruments = cli.Command{
	Name:    "get-instruments",
	Usage:   "Retrieves a list of tradeable instruments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "easy-to-borrow",
			Usage:     "Filter by easy to borrow status",
			QueryPath: "easy_to_borrow",
		},
		&requestflag.Flag[string]{
			Name:      "id-filter",
			Usage:     "Filter IDs to those containing this substring. For options, and when security_type is omitted and no security_id/security_id_source filters are provided, this is required.",
			QueryPath: "id_filter",
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
			Name:      "is-restricted",
			Usage:     "Filter by restricted status",
			QueryPath: "is_restricted",
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
			Usage:     "Filter by security type. If omitted, returns all types.",
			QueryPath: "security_type",
		},
	},
	Action:          handleActiveV1InstrumentsGetInstruments,
	HideHelpCommand: true,
}

func handleActiveV1InstrumentsGetInstrumentByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("security-id") && len(unusedArgs) > 0 {
		cmd.Set("security-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1InstrumentGetInstrumentByIDParams{
		SecurityIDSource: clearstreet.SecurityIDSource(cmd.Value("security-id-source").(string)),
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
	_, err = client.Active.V1.Instruments.GetInstrumentByID(
		ctx,
		cmd.Value("security-id").(string),
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
		Title:          "active:v1:instruments get-instrument-by-id",
		Transform:      transform,
	})
}

func handleActiveV1InstrumentsGetInstruments(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1InstrumentGetInstrumentsParams{}

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
	_, err = client.Active.V1.Instruments.GetInstruments(ctx, params, options...)
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
		Title:          "active:v1:instruments get-instruments",
		Transform:      transform,
	})
}
