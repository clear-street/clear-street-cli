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

var v1InstrumentsOptionsGetOptionContracts = cli.Command{
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
			Default:   1000,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next page of results. Contains encoded pagination state (limit + offset).\nWhen provided, page_size is ignored.",
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
	Action:          handleV1InstrumentsOptionsGetOptionContracts,
	HideHelpCommand: true,
}

func handleV1InstrumentsOptionsGetOptionContracts(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentOptionGetOptionContractsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.Options.GetOptionContracts(ctx, params, options...)
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
		Title:          "v1:instruments:options get-option-contracts",
		Transform:      transform,
	})
}
