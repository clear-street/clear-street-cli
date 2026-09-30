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

var v1CalendarGetEconomicEventsCalendar = requestflag.WithInnerFlags(cli.Command{
	Name:    "get-economic-events-calendar",
	Usage:   "Retrieves macroeconomic calendar events (e.g. CPI, jobs reports, central bank\nrate decisions), optionally filtered by country, impact, and event time range.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country",
			Usage:     "Comma-separated ISO 3166-1 alpha-2 country codes (or `EU`) to filter by. Defaults to `US` when omitted.",
			QueryPath: "country",
		},
		&requestflag.Flag[[]string]{
			Name:      "impact",
			Usage:     "Comma-separated impact levels to filter by.",
			QueryPath: "impact",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   1000,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "timestamp",
			QueryPath: "timestamp",
		},
	},
	Action:          handleV1CalendarGetEconomicEventsCalendar,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"timestamp": {
		&requestflag.InnerFlag[string]{
			Name:       "timestamp.gt",
			Usage:      "Return only rows where `timestamp` is strictly after the given value. A bare `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches from the start of the following day. See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "gt",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timestamp.gte",
			Usage:      "Return only rows where `timestamp` is on or after the given value. A bare `YYYY-MM-DD` date expands to the start of that day (UTC). See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timestamp.lt",
			Usage:      "Return only rows where `timestamp` is strictly before the given value. A bare `YYYY-MM-DD` date expands to the start of that day (UTC). See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "lt",
		},
		&requestflag.InnerFlag[string]{
			Name:       "timestamp.lte",
			Usage:      "Return only rows where `timestamp` is on or before the given value. A bare `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches through the end of that day. See [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters) for accepted formats, bare-date expansion, and combining bounds. Returns 400 if the resulting range is inverted.",
			InnerField: "lte",
		},
	},
})

var v1CalendarGetMarketHoursCalendar = cli.Command{
	Name:    "get-market-hours-calendar",
	Usage:   "Retrieves comprehensive trading hours including pre-market, regular, and\nafter-hours sessions. Returns market status, session times, and next session\nschedules.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "date",
			Usage:     "The date to query market hours for (YYYY-MM-DD). Defaults to today.",
			QueryPath: "date",
		},
		&requestflag.Flag[string]{
			Name:      "market",
			Usage:     "Market type to query (us_equities, us_options). If omitted, returns all markets.",
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
		apiquery.ArrayQueryFormatComma,
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

func handleV1CalendarGetEconomicEventsCalendar(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1CalendarGetEconomicEventsCalendarParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Calendar.GetEconomicEventsCalendar(ctx, params, options...)
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
		Title:          "v1:calendar get-economic-events-calendar",
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
		apiquery.ArrayQueryFormatComma,
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
