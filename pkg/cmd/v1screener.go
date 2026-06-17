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

var v1ScreenerCreateScreener = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-screener",
	Usage:   "Create a saved screener configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "column",
			Usage:    "Structured field references to include when running this screener",
			BodyPath: "columns",
		},
		&requestflag.Flag[any]{
			Name:     "field-filter",
			Usage:    "Deprecated: use `columns` instead. Ignored when `columns` is provided.",
			BodyPath: "field_filter",
		},
		&requestflag.Flag[any]{
			Name:     "filter",
			Usage:    "Structured search filter criteria",
			BodyPath: "filters",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "The name for this screener configuration",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "sort",
			Usage:    "Multi-field sort specifications",
			BodyPath: "sorts",
		},
	},
	Action:          handleV1ScreenerCreateScreener,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"column": {
		&requestflag.InnerFlag[string]{
			Name:                  "column.name",
			Usage:                 "The field name.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.lookback",
			Usage:                 "Historical lookback window for price/change fields.",
			InnerField:            "lookback",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.period",
			Usage:                 "Reporting period for financial data fields.",
			InnerField:            "period",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.value-type",
			Usage:                 "The data type of a screener field value.",
			InnerField:            "value_type",
			OuterIsArrayOfObjects: true,
		},
	},
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
	"sort": {
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "sort.field",
			Usage:                 "A reference to a screener field.",
			InnerField:            "field",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "sort.direction",
			Usage:                 "Sort direction sorted results",
			InnerField:            "direction",
			OuterIsArrayOfObjects: true,
		},
	},
})

var v1ScreenerDeleteScreener = cli.Command{
	Name:    "delete-screener",
	Usage:   "Delete a saved screener configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "screener-id",
			Required:  true,
			PathParam: "screener_id",
		},
	},
	Action:          handleV1ScreenerDeleteScreener,
	HideHelpCommand: true,
}

var v1ScreenerGetScreenerByID = cli.Command{
	Name:    "get-screener-by-id",
	Usage:   "Get a saved screener configuration by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "screener-id",
			Required:  true,
			PathParam: "screener_id",
		},
	},
	Action:          handleV1ScreenerGetScreenerByID,
	HideHelpCommand: true,
}

var v1ScreenerGetScreeners = cli.Command{
	Name:            "get-screeners",
	Usage:           "List saved screener configurations.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleV1ScreenerGetScreeners,
	HideHelpCommand: true,
}

var v1ScreenerReplaceScreener = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace-screener",
	Usage:   "Update a saved screener configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "screener-id",
			Required:  true,
			PathParam: "screener_id",
		},
		&requestflag.Flag[any]{
			Name:     "column",
			Usage:    "Structured field references to include when running this screener",
			BodyPath: "columns",
		},
		&requestflag.Flag[any]{
			Name:     "field-filter",
			Usage:    "Deprecated: use `columns` instead. Ignored when `columns` is provided.",
			BodyPath: "field_filter",
		},
		&requestflag.Flag[any]{
			Name:     "filter",
			Usage:    "Structured search filter criteria",
			BodyPath: "filters",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "The name for this screener configuration",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "sort",
			Usage:    "Multi-field sort specifications",
			BodyPath: "sorts",
		},
	},
	Action:          handleV1ScreenerReplaceScreener,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"column": {
		&requestflag.InnerFlag[string]{
			Name:                  "column.name",
			Usage:                 "The field name.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.lookback",
			Usage:                 "Historical lookback window for price/change fields.",
			InnerField:            "lookback",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.period",
			Usage:                 "Reporting period for financial data fields.",
			InnerField:            "period",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.value-type",
			Usage:                 "The data type of a screener field value.",
			InnerField:            "value_type",
			OuterIsArrayOfObjects: true,
		},
	},
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
	"sort": {
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "sort.field",
			Usage:                 "A reference to a screener field.",
			InnerField:            "field",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "sort.direction",
			Usage:                 "Sort direction sorted results",
			InnerField:            "direction",
			OuterIsArrayOfObjects: true,
		},
	},
})

var v1ScreenerSearchScreener = requestflag.WithInnerFlags(cli.Command{
	Name:    "search-screener",
	Usage:   "Search instruments using structured filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "column",
			Usage:    "Subset of fields to include in the response.",
			BodyPath: "columns",
		},
		&requestflag.Flag[any]{
			Name:     "field-filter",
			Usage:    "Deprecated: use `columns` instead. Ignored when `columns` is provided.",
			BodyPath: "field_filter",
		},
		&requestflag.Flag[any]{
			Name:     "filter",
			Usage:    "Filter conditions to apply.",
			BodyPath: "filters",
		},
		&requestflag.Flag[*int64]{
			Name:     "page-size",
			Usage:    "The number of items to return per page (only used when page_token is not provided)",
			BodyPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "Token for retrieving the next page of results. Contains encoded pagination state (limit + offset).\nWhen provided, page_size is ignored.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[*bool]{
			Name:     "sort-case-sensitive",
			Usage:    "Whether string sorts should be case-sensitive (default: false).",
			BodyPath: "sort_case_sensitive",
		},
		&requestflag.Flag[any]{
			Name:     "sort",
			Usage:    "Multi-field sort specifications.",
			BodyPath: "sorts",
		},
	},
	Action:          handleV1ScreenerSearchScreener,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"column": {
		&requestflag.InnerFlag[string]{
			Name:                  "column.name",
			Usage:                 "The field name.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.lookback",
			Usage:                 "Historical lookback window for price/change fields.",
			InnerField:            "lookback",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.period",
			Usage:                 "Reporting period for financial data fields.",
			InnerField:            "period",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "column.value-type",
			Usage:                 "The data type of a screener field value.",
			InnerField:            "value_type",
			OuterIsArrayOfObjects: true,
		},
	},
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
	"sort": {
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "sort.field",
			Usage:                 "A reference to a screener field.",
			InnerField:            "field",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "sort.direction",
			Usage:                 "Sort direction sorted results",
			InnerField:            "direction",
			OuterIsArrayOfObjects: true,
		},
	},
})

func handleV1ScreenerCreateScreener(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1ScreenerNewScreenerParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Screener.NewScreener(ctx, params, options...)
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
		Title:          "v1:screener create-screener",
		Transform:      transform,
	})
}

func handleV1ScreenerDeleteScreener(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("screener-id") && len(unusedArgs) > 0 {
		cmd.Set("screener-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	return client.V1.Screener.DeleteScreener(ctx, cmd.Value("screener-id").(string), options...)
}

func handleV1ScreenerGetScreenerByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("screener-id") && len(unusedArgs) > 0 {
		cmd.Set("screener-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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
	_, err = client.V1.Screener.GetScreenerByID(ctx, cmd.Value("screener-id").(string), options...)
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
		Title:          "v1:screener get-screener-by-id",
		Transform:      transform,
	})
}

func handleV1ScreenerGetScreeners(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Screener.GetScreeners(ctx, options...)
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
		Title:          "v1:screener get-screeners",
		Transform:      transform,
	})
}

func handleV1ScreenerReplaceScreener(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("screener-id") && len(unusedArgs) > 0 {
		cmd.Set("screener-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1ScreenerReplaceScreenerParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Screener.ReplaceScreener(
		ctx,
		cmd.Value("screener-id").(string),
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
		Title:          "v1:screener replace-screener",
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
		apiquery.ArrayQueryFormatComma,
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
