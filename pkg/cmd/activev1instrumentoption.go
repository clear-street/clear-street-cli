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

var activeV1InstrumentsOptionsContracts = cli.Command{
	Name:    "contracts",
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
			Name:      "underlier-instrument-id",
			Usage:     "OEMS instrument UUID of the underlying equity/index",
			QueryPath: "underlier_instrument_id",
		},
		&requestflag.Flag[string]{
			Name:      "underlier-security-id",
			Usage:     "Security identifier of the underlying (e.g., CUSIP, ISIN). Must be paired with underlier_security_id_source.",
			QueryPath: "underlier_security_id",
		},
		&requestflag.Flag[string]{
			Name:      "underlier-security-id-source",
			Usage:     "Security identifier source",
			QueryPath: "underlier_security_id_source",
		},
	},
	Action:          handleActiveV1InstrumentsOptionsContracts,
	HideHelpCommand: true,
}

func handleActiveV1InstrumentsOptionsContracts(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1InstrumentOptionContractsParams{}

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
	_, err = client.Active.V1.Instruments.Options.Contracts(ctx, params, options...)
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
		Title:          "active:v1:instruments:options contracts",
		Transform:      transform,
	})
}
