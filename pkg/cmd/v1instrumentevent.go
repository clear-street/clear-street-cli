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

var v1InstrumentsEventsGetAllInstrumentEvents = cli.Command{
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
	Action:          handleV1InstrumentsEventsGetAllInstrumentEvents,
	HideHelpCommand: true,
}

var v1InstrumentsEventsGetInstrumentEvents = cli.Command{
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
	Action:          handleV1InstrumentsEventsGetInstrumentEvents,
	HideHelpCommand: true,
}

func handleV1InstrumentsEventsGetAllInstrumentEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentEventGetAllInstrumentEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.Events.GetAllInstrumentEvents(ctx, params, options...)
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
		Title:          "v1:instruments:events get-all-instrument-events",
		Transform:      transform,
	})
}

func handleV1InstrumentsEventsGetInstrumentEvents(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentEventGetInstrumentEventsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.Events.GetInstrumentEvents(
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
		Title:          "v1:instruments:events get-instrument-events",
		Transform:      transform,
	})
}
