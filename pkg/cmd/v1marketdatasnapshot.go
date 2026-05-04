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

var v1MarketDataSnapshotGetSnapshots = cli.Command{
	Name:    "get-snapshots",
	Usage:   "Get market data snapshots for one or more securities.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-separated OEMS instrument UUIDs.",
			QueryPath: "instrument_ids",
		},
	},
	Action:          handleV1MarketDataSnapshotGetSnapshots,
	HideHelpCommand: true,
}

func handleV1MarketDataSnapshotGetSnapshots(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1MarketDataSnapshotGetSnapshotsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.MarketData.Snapshot.GetSnapshots(ctx, params, options...)
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
		Title:          "v1:market-data:snapshot get-snapshots",
		Transform:      transform,
	})
}
