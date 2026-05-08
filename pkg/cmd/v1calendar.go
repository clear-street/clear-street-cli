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

var v1CalendarGetClock = cli.Command{
	Name:            "get-clock",
	Usage:           "Returns the current server time in UTC.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleV1CalendarGetClock,
	HideHelpCommand: true,
}

var v1CalendarGetMarketHoursCalendar = cli.Command{
	Name:    "get-market-hours-calendar",
	Usage:   "Retrieves comprehensive trading hours including pre-market, regular, and\nafter-hours sessions. Returns market status, session times, and next session\nschedules.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "date",
			Usage:     "The date to query market hours for (YYYY-MM-DD). Defaults to today.",
			Required:  true,
			QueryPath: "date",
		},
		&requestflag.Flag[string]{
			Name:      "market",
			Usage:     "Market type for market hours calendar endpoint",
			QueryPath: "market",
		},
	},
	Action:          handleV1CalendarGetMarketHoursCalendar,
	HideHelpCommand: true,
}

func handleV1CalendarGetClock(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Calendar.GetClock(ctx, options...)
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
		Title:          "v1:calendar get-clock",
		Transform:      transform,
	})
}

func handleV1CalendarGetMarketHoursCalendar(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1CalendarGetMarketHoursCalendarParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Calendar.GetMarketHoursCalendar(ctx, params, options...)
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
		Title:          "v1:calendar get-market-hours-calendar",
		Transform:      transform,
	})
}
