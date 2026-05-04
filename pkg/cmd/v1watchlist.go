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

var v1WatchlistsCreateWatchlist = cli.Command{
	Name:    "create-watchlist",
	Usage:   "Create Watchlist",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The desired watchlist name.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleV1WatchlistsCreateWatchlist,
	HideHelpCommand: true,
}

var v1WatchlistsDeleteWatchlist = cli.Command{
	Name:    "delete-watchlist",
	Usage:   "Delete a watchlist and all its items",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "watchlist-id",
			Required:  true,
			PathParam: "watchlist_id",
		},
	},
	Action:          handleV1WatchlistsDeleteWatchlist,
	HideHelpCommand: true,
}

var v1WatchlistsGetWatchlistByID = cli.Command{
	Name:    "get-watchlist-by-id",
	Usage:   "Get a watchlist by ID with all its items",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "watchlist-id",
			Required:  true,
			PathParam: "watchlist_id",
		},
	},
	Action:          handleV1WatchlistsGetWatchlistByID,
	HideHelpCommand: true,
}

var v1WatchlistsGetWatchlists = cli.Command{
	Name:    "get-watchlists",
	Usage:   "List watchlists for the authenticated user",
	Suggest: true,
	Flags: []cli.Flag{
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
	},
	Action:          handleV1WatchlistsGetWatchlists,
	HideHelpCommand: true,
}

func handleV1WatchlistsCreateWatchlist(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1WatchlistNewWatchlistParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Watchlists.NewWatchlist(ctx, params, options...)
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
		Title:          "v1:watchlists create-watchlist",
		Transform:      transform,
	})
}

func handleV1WatchlistsDeleteWatchlist(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("watchlist-id") && len(unusedArgs) > 0 {
		cmd.Set("watchlist-id", unusedArgs[0])
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
	_, err = client.V1.Watchlists.DeleteWatchlist(ctx, cmd.Value("watchlist-id").(string), options...)
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
		Title:          "v1:watchlists delete-watchlist",
		Transform:      transform,
	})
}

func handleV1WatchlistsGetWatchlistByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("watchlist-id") && len(unusedArgs) > 0 {
		cmd.Set("watchlist-id", unusedArgs[0])
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
	_, err = client.V1.Watchlists.GetWatchlistByID(ctx, cmd.Value("watchlist-id").(string), options...)
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
		Title:          "v1:watchlists get-watchlist-by-id",
		Transform:      transform,
	})
}

func handleV1WatchlistsGetWatchlists(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1WatchlistGetWatchlistsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Watchlists.GetWatchlists(ctx, params, options...)
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
		Title:          "v1:watchlists get-watchlists",
		Transform:      transform,
	})
}
