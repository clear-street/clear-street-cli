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

var v1WatchlistAddWatchlistItem = cli.Command{
	Name:    "add-watchlist-item",
	Usage:   "Add an instrument to a watchlist",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "watchlist-id",
			Required:  true,
			PathParam: "watchlist_id",
		},
		&requestflag.Flag[string]{
			Name:     "instrument-id",
			Usage:    "OEMS instrument UUID",
			Required: true,
			BodyPath: "instrument_id",
		},
	},
	Action:          handleV1WatchlistAddWatchlistItem,
	HideHelpCommand: true,
}

var v1WatchlistCreateWatchlist = cli.Command{
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
	Action:          handleV1WatchlistCreateWatchlist,
	HideHelpCommand: true,
}

var v1WatchlistDeleteWatchlist = cli.Command{
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
	Action:          handleV1WatchlistDeleteWatchlist,
	HideHelpCommand: true,
}

var v1WatchlistDeleteWatchlistItem = cli.Command{
	Name:    "delete-watchlist-item",
	Usage:   "Delete an instrument from a watchlist",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "watchlist-id",
			Required:  true,
			PathParam: "watchlist_id",
		},
		&requestflag.Flag[string]{
			Name:      "item-id",
			Required:  true,
			PathParam: "item_id",
		},
	},
	Action:          handleV1WatchlistDeleteWatchlistItem,
	HideHelpCommand: true,
}

var v1WatchlistGetWatchlistByID = cli.Command{
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
	Action:          handleV1WatchlistGetWatchlistByID,
	HideHelpCommand: true,
}

var v1WatchlistGetWatchlists = cli.Command{
	Name:    "get-watchlists",
	Usage:   "List watchlists for the authenticated user",
	Suggest: true,
	Flags: []cli.Flag{
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
	},
	Action:          handleV1WatchlistGetWatchlists,
	HideHelpCommand: true,
}

func handleV1WatchlistAddWatchlistItem(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := clearstreet.V1WatchlistAddWatchlistItemParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Watchlist.AddWatchlistItem(
		ctx,
		cmd.Value("watchlist-id").(string),
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
		Title:          "v1:watchlist add-watchlist-item",
		Transform:      transform,
	})
}

func handleV1WatchlistCreateWatchlist(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Watchlist.NewWatchlist(ctx, params, options...)
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
		Title:          "v1:watchlist create-watchlist",
		Transform:      transform,
	})
}

func handleV1WatchlistDeleteWatchlist(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Watchlist.DeleteWatchlist(ctx, cmd.Value("watchlist-id").(string), options...)
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
		Title:          "v1:watchlist delete-watchlist",
		Transform:      transform,
	})
}

func handleV1WatchlistDeleteWatchlistItem(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("item-id") && len(unusedArgs) > 0 {
		cmd.Set("item-id", unusedArgs[0])
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

	params := clearstreet.V1WatchlistDeleteWatchlistItemParams{
		WatchlistID: cmd.Value("watchlist-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Watchlist.DeleteWatchlistItem(
		ctx,
		cmd.Value("item-id").(string),
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
		Title:          "v1:watchlist delete-watchlist-item",
		Transform:      transform,
	})
}

func handleV1WatchlistGetWatchlistByID(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Watchlist.GetWatchlistByID(ctx, cmd.Value("watchlist-id").(string), options...)
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
		Title:          "v1:watchlist get-watchlist-by-id",
		Transform:      transform,
	})
}

func handleV1WatchlistGetWatchlists(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Watchlist.GetWatchlists(ctx, params, options...)
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
		Title:          "v1:watchlist get-watchlists",
		Transform:      transform,
	})
}
