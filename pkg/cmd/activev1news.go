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

var activeV1NewsGetNews = cli.Command{
	Name:    "get-news",
	Usage:   "Retrieves news items with optional filtering by security IDs, time range,\npublisher, type, and text query.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "exclude-publishers",
			Usage:     "Comma-separated list of publishers to exclude (mutually exclusive with include_publishers).",
			QueryPath: "exclude_publishers",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive start timestamp. Accepts `YYYY-MM-DD` or RFC3339 datetime.",
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "include-publishers",
			Usage:     "Comma-separated list of publishers to include (mutually exclusive with exclude_publishers).",
			QueryPath: "include_publishers",
		},
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-delimited OEMS instrument UUIDs to filter by.",
			QueryPath: "instrument_ids",
		},
		&requestflag.Flag[string]{
			Name:      "news-type",
			Usage:     "Filter by news type.",
			QueryPath: "news_type",
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
			Name:      "search-query",
			Usage:     "Free-text query matched against title/text and associated security IDs.",
			QueryPath: "search_query",
		},
		&requestflag.Flag[[]string]{
			Name:      "sector",
			Usage:     "Comma-separated sector values to filter by.",
			QueryPath: "sectors",
		},
		&requestflag.Flag[[]string]{
			Name:      "security-id",
			Usage:     "Filter by security ID(s). Accepts single value or indexed array.\n\nExamples:\n- Single: `security_id=037833100`\n- Multiple: `security_id[0]=037833100&security_id[1]=594918104`",
			QueryPath: "security_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "security-id-source",
			Usage:     "Source(s) for the security ID filter. Must match the count and order of security_id.\n\nExamples:\n- Single: `security_id_source=CUSIP`\n- Multiple: `security_id_source[0]=CUSIP&security_id_source[1]=FIGI`",
			QueryPath: "security_id_source",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive end timestamp. Accepts `YYYY-MM-DD` or RFC3339 datetime.",
			QueryPath: "to",
		},
	},
	Action:          handleActiveV1NewsGetNews,
	HideHelpCommand: true,
}

func handleActiveV1NewsGetNews(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.ActiveV1NewsGetNewsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Active.V1.News.GetNews(ctx, params, options...)
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
		Title:          "active:v1:news get-news",
		Transform:      transform,
	})
}
