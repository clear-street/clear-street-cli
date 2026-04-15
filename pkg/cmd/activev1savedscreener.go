// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/clear-street/clear-street-go"
	"github.com/clear-street/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-cli/internal/apiquery"
	"github.com/stainless-sdks/clear-street-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var activeV1SavedScreenersCreateScreener = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-screener",
	Usage:   "Persists a screener configuration for the authenticated user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "field-filter",
			Usage:    "List of field names to include when running this screener",
			BodyPath: "field_filter",
		},
		&requestflag.Flag[any]{
			Name:     "filter",
			Usage:    "Filter criteria for this screener",
			BodyPath: "filters",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "The name for this screener configuration",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "sort-by",
			Usage:    "Field name to sort results by",
			BodyPath: "sort_by",
		},
		&requestflag.Flag[any]{
			Name:     "sort-direction",
			Usage:    "Sort direction for results",
			BodyPath: "sort_direction",
		},
	},
	Action:          handleActiveV1SavedScreenersCreateScreener,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.field-name",
			Usage:      "The field name to filter on",
			InnerField: "field_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.operation",
			Usage:      "The filter operation (lt, lte, gt, gte, eq, rgx, bw, ew)",
			InnerField: "operation",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.value",
			Usage:      "The filter value",
			InnerField: "value",
		},
	},
})

var activeV1SavedScreenersDeleteScreener = cli.Command{
	Name:    "delete-screener",
	Usage:   "Deletes the screener configuration for the authenticated user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "screener-id",
			Required: true,
		},
	},
	Action:          handleActiveV1SavedScreenersDeleteScreener,
	HideHelpCommand: true,
}

var activeV1SavedScreenersGetScreenerByID = cli.Command{
	Name:    "get-screener-by-id",
	Usage:   "Returns a single screener configuration for the authenticated user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "screener-id",
			Required: true,
		},
	},
	Action:          handleActiveV1SavedScreenersGetScreenerByID,
	HideHelpCommand: true,
}

var activeV1SavedScreenersListScreeners = cli.Command{
	Name:            "list-screeners",
	Usage:           "Returns all screener configurations for the authenticated user.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleActiveV1SavedScreenersListScreeners,
	HideHelpCommand: true,
}

var activeV1SavedScreenersUpdateScreener = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-screener",
	Usage:   "Replaces the screener configuration for the authenticated user. If `name` is\nnull, the existing name is preserved.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "screener-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "field-filter",
			Usage:    "List of field names to include when running this screener",
			BodyPath: "field_filter",
		},
		&requestflag.Flag[any]{
			Name:     "filter",
			Usage:    "Filter criteria for this screener",
			BodyPath: "filters",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "The name for this screener configuration",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "sort-by",
			Usage:    "Field name to sort results by",
			BodyPath: "sort_by",
		},
		&requestflag.Flag[any]{
			Name:     "sort-direction",
			Usage:    "Sort direction for results",
			BodyPath: "sort_direction",
		},
	},
	Action:          handleActiveV1SavedScreenersUpdateScreener,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.field-name",
			Usage:      "The field name to filter on",
			InnerField: "field_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.operation",
			Usage:      "The filter operation (lt, lte, gt, gte, eq, rgx, bw, ew)",
			InnerField: "operation",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.value",
			Usage:      "The filter value",
			InnerField: "value",
		},
	},
})

func handleActiveV1SavedScreenersCreateScreener(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1SavedScreenerNewScreenerParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.SavedScreeners.NewScreener(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:saved-screeners create-screener", obj, format, explicitFormat, transform)
}

func handleActiveV1SavedScreenersDeleteScreener(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Active.V1.SavedScreeners.DeleteScreener(ctx, cmd.Value("screener-id").(string), options...)
}

func handleActiveV1SavedScreenersGetScreenerByID(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatIndices,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.SavedScreeners.GetScreenerByID(ctx, cmd.Value("screener-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:saved-screeners get-screener-by-id", obj, format, explicitFormat, transform)
}

func handleActiveV1SavedScreenersListScreeners(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Active.V1.SavedScreeners.ListScreeners(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:saved-screeners list-screeners", obj, format, explicitFormat, transform)
}

func handleActiveV1SavedScreenersUpdateScreener(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("screener-id") && len(unusedArgs) > 0 {
		cmd.Set("screener-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1SavedScreenerUpdateScreenerParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.SavedScreeners.UpdateScreener(
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
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:saved-screeners update-screener", obj, format, explicitFormat, transform)
}
