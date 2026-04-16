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

var activeV1InstrumentsEventsGetAllInstrumentEvents = cli.Command{
	Name:    "get-all-instrument-events",
	Usage:   "Retrieves all instrument events grouped by date.",
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
			Name:      "to-date",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD).",
			QueryPath: "to_date",
		},
	},
	Action:          handleActiveV1InstrumentsEventsGetAllInstrumentEvents,
	HideHelpCommand: true,
}

var activeV1InstrumentsEventsGetInstrumentEvents = cli.Command{
	Name:    "get-instrument-events",
	Usage:   "Retrieves corporate events (dividends, splits, etc.) for an instrument, grouped\nby event type.",
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
	Action:          handleActiveV1InstrumentsEventsGetInstrumentEvents,
	HideHelpCommand: true,
}

func handleActiveV1InstrumentsEventsGetAllInstrumentEvents(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1InstrumentEventGetAllInstrumentEventsParams{}

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
	_, err = client.Active.V1.Instruments.Events.GetAllInstrumentEvents(ctx, params, options...)
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
		Title:          "active:v1:instruments:events get-all-instrument-events",
		Transform:      transform,
	})
}

func handleActiveV1InstrumentsEventsGetInstrumentEvents(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("security-id") && len(unusedArgs) > 0 {
		cmd.Set("security-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1InstrumentEventGetInstrumentEventsParams{
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
	_, err = client.Active.V1.Instruments.Events.GetInstrumentEvents(
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
		Title:          "active:v1:instruments:events get-instrument-events",
		Transform:      transform,
	})
}
