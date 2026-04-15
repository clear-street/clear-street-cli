// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/clear-street/clear-street-cli/internal/apiquery"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
	"github.com/clear-street/clear-street-go"
	"github.com/clear-street/clear-street-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var activeV1CalendarsMergersAcquisitionsGetMergersAndAcquisitionsCalendar = cli.Command{
	Name:    "get-mergers-and-acquisitions-calendar",
	Usage:   "Retrieves upcoming M&A events.",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleActiveV1CalendarsMergersAcquisitionsGetMergersAndAcquisitionsCalendar,
	HideHelpCommand: true,
}

func handleActiveV1CalendarsMergersAcquisitionsGetMergersAndAcquisitionsCalendar(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams{}

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
	_, err = client.Active.V1.Calendars.MergersAcquisitions.GetMergersAndAcquisitionsCalendar(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:calendars:mergers-acquisitions get-mergers-and-acquisitions-calendar", obj, format, explicitFormat, transform)
}
