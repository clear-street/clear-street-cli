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

var activeV1ScreenerGetScreener = cli.Command{
	Name:    "get-screener",
	Usage:   "Searches for instruments matching specified criteria.",
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
			Usage:     "Number of items to return per page (default: 100, max: 10000)",
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next page of results. Contains encoded pagination state.",
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
	Action:          handleActiveV1ScreenerGetScreener,
	HideHelpCommand: true,
}

func handleActiveV1ScreenerGetScreener(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1ScreenerGetScreenerParams{}

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
	_, err = client.Active.V1.Screener.GetScreener(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:screener get-screener", obj, format, explicitFormat, transform)
}
