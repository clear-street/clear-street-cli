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

var v1ScreenerGetScreener = cli.Command{
	Name:    "get-screener",
	Usage:   "Screen instruments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "field-filter",
			Usage:     "Comma-separated list of field names to include in the response",
			QueryPath: "field_filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Dynamic filters with dot notation (e.g., filter[price.gte]=50, filter[symbol.bw]=A)",
			QueryPath: "filter",
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
			Name:      "sort-by",
			Usage:     "Field to sort by",
			QueryPath: "sort_by",
		},
		&requestflag.Flag[string]{
			Name:      "sort-direction",
			Usage:     "Sort direction (ASC or DESC, defaults to DESC)",
			QueryPath: "sort_direction",
		},
	},
	Action:          handleV1ScreenerGetScreener,
	HideHelpCommand: true,
}

var v1ScreenerSearchScreener = requestflag.WithInnerFlags(cli.Command{
	Name:    "search-screener",
	Usage:   "Search instruments using structured filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "field-filter",
			Usage:    "Subset of fields to include in the response.",
			BodyPath: "field_filter",
		},
		&requestflag.Flag[any]{
			Name:     "filter",
			Usage:    "Filter conditions to apply.",
			BodyPath: "filters",
		},
		&requestflag.Flag[*int64]{
			Name:     "page-size",
			Usage:    "Maximum number of results per page.",
			BodyPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "Opaque token for cursor-based pagination.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sort-by",
			Usage:    "A reference to a screener field.",
			BodyPath: "sort_by",
		},
		&requestflag.Flag[*bool]{
			Name:     "sort-case-sensitive",
			Usage:    "Whether string sorts should be case-sensitive (default: false).",
			BodyPath: "sort_case_sensitive",
		},
		&requestflag.Flag[string]{
			Name:     "sort-direction",
			Usage:    "Sort direction (defaults to DESC).",
			BodyPath: "sort_direction",
		},
		&requestflag.Flag[any]{
			Name:     "sort",
			Usage:    "Multi-field sort specifications. When present, takes precedence over sort_by/sort_direction.",
			BodyPath: "sorts",
		},
	},
	Action:          handleV1ScreenerSearchScreener,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"field-filter": {
		&requestflag.InnerFlag[string]{
			Name:                  "field-filter.name",
			Usage:                 "The field name.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "field-filter.lookback",
			Usage:                 "Historical lookback window for price/change fields.",
			InnerField:            "lookback",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "field-filter.period",
			Usage:                 "Reporting period for financial data fields.",
			InnerField:            "period",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "field-filter.value-type",
			Usage:                 "The data type of a screener field value.",
			InnerField:            "value_type",
			OuterIsArrayOfObjects: true,
		},
	},
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "filter.left",
			Usage:                 "A reference to a screener field.",
			InnerField:            "left",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "filter.op",
			Usage:                 "Operator specification with optional behavioral arguments.",
			InnerField:            "op",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "filter.right",
			Usage:                 "The value(s) to compare against. Omit together with `op` for an unenabled filter.",
			InnerField:            "right",
			OuterIsArrayOfObjects: true,
		},
	},
	"sort-by": {
		&requestflag.InnerFlag[string]{
			Name:       "sort-by.name",
			Usage:      "The field name.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sort-by.lookback",
			Usage:      "Historical lookback window for price/change fields.",
			InnerField: "lookback",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sort-by.period",
			Usage:      "Reporting period for financial data fields.",
			InnerField: "period",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sort-by.value-type",
			Usage:      "The data type of a screener field value.",
			InnerField: "value_type",
		},
	},
	"sort": {
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "sort.field",
			Usage:                 "A reference to a screener field.",
			InnerField:            "field",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "sort.direction",
			Usage:                 "Sort direction (defaults to DESC).",
			InnerField:            "direction",
			OuterIsArrayOfObjects: true,
		},
	},
})

func handleV1ScreenerGetScreener(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1ScreenerGetScreenerParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Screener.GetScreener(ctx, params, options...)
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
		Title:          "v1:screener get-screener",
		Transform:      transform,
	})
}

func handleV1ScreenerSearchScreener(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatIndices,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1ScreenerSearchScreenerParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Screener.SearchScreener(ctx, params, options...)
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
		Title:          "v1:screener search-screener",
		Transform:      transform,
	})
}
