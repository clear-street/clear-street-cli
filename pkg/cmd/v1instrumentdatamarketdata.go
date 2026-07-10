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

var v1InstrumentDataMarketDataGetDailySummaries = cli.Command{
	Name:    "get-daily-summaries",
	Usage:   "Returns the most recent open, high, low, volume (OHLV) and current price for the\nrequested instruments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-ids",
			Usage:     "Comma-separated instrument identifiers (required, 1..=100)",
			Required:  true,
			QueryPath: "instrument_ids",
		},
	},
	Action:          handleV1InstrumentDataMarketDataGetDailySummaries,
	HideHelpCommand: true,
}

var v1InstrumentDataMarketDataGetSnapshots = cli.Command{
	Name:    "get-snapshots",
	Usage:   "Get market data snapshots for one or more securities.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-separated instrument identifiers.",
			QueryPath: "instrument_ids",
		},
	},
	Action:          handleV1InstrumentDataMarketDataGetSnapshots,
	HideHelpCommand: true,
}

func handleV1InstrumentDataMarketDataGetDailySummaries(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentDataMarketDataGetDailySummariesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.MarketData.GetDailySummaries(ctx, params, options...)
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
		Title:          "v1:instrument-data:market-data get-daily-summaries",
		Transform:      transform,
	})
}

func handleV1InstrumentDataMarketDataGetSnapshots(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentDataMarketDataGetSnapshotsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.InstrumentData.MarketData.GetSnapshots(ctx, params, options...)
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
		Title:          "v1:instrument-data:market-data get-snapshots",
		Transform:      transform,
	})
}
