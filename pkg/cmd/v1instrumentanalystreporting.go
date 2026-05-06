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

var v1InstrumentsAnalystReportingGetInstrumentAnalystConsensus = cli.Command{
	Name:    "get-instrument-analyst-consensus",
	Usage:   "Retrieves analyst ratings and price targets for an instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "The start date for the query range, inclusive (YYYY-MM-DD)",
			QueryPath: "from",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "The end date for the query range, inclusive (YYYY-MM-DD)",
			QueryPath: "to",
		},
	},
	Action:          handleV1InstrumentsAnalystReportingGetInstrumentAnalystConsensus,
	HideHelpCommand: true,
}

func handleV1InstrumentsAnalystReportingGetInstrumentAnalystConsensus(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Instruments.AnalystReporting.GetInstrumentAnalystConsensus(
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
		Title:          "v1:instruments:analyst-reporting get-instrument-analyst-consensus",
		Transform:      transform,
	})
}
