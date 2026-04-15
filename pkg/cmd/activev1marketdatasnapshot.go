// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/clear-street-cli/internal/apiquery"
	"github.com/stainless-sdks/clear-street-cli/internal/requestflag"
	"github.com/stainless-sdks/clear-street-go"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var activeV1MarketDataSnapshotGetSnapshots = cli.Command{
	Name:    "get-snapshots",
	Usage:   "Get market data snapshots for one or more securities.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ids",
			Usage:     "Comma-separated OEMS instrument UUIDs",
			QueryPath: "ids",
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
	},
	Action:          handleActiveV1MarketDataSnapshotGetSnapshots,
	HideHelpCommand: true,
}

func handleActiveV1MarketDataSnapshotGetSnapshots(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1MarketDataSnapshotGetSnapshotsParams{}

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
	_, err = client.Active.V1.MarketData.Snapshot.GetSnapshots(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:market-data:snapshot get-snapshots", obj, format, explicitFormat, transform)
}
