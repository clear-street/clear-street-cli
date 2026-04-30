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

var activeV1MarketDataDailySummaryGetDailySummaries = cli.Command{
	Name:    "get-daily-summaries",
	Usage:   "Returns the most recent OHLV and current price for the requested OEMS\ninstruments. Backed by the in-memory Polygon snapshot cache.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-ids",
			Usage:     "Comma-separated OEMS instrument UUIDs (required, 1..=100)",
			Required:  true,
			QueryPath: "instrument_ids",
		},
	},
	Action:          handleActiveV1MarketDataDailySummaryGetDailySummaries,
	HideHelpCommand: true,
}

func handleActiveV1MarketDataDailySummaryGetDailySummaries(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.ActiveV1MarketDataDailySummaryGetDailySummariesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.MarketData.DailySummary.GetDailySummaries(ctx, params, options...)
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
		Title:          "active:v1:market-data:daily-summary get-daily-summaries",
		Transform:      transform,
	})
}
