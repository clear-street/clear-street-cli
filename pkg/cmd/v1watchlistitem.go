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

var v1WatchlistsItemsAddWatchlistItem = cli.Command{
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
	Action:          handleV1WatchlistsItemsAddWatchlistItem,
	HideHelpCommand: true,
}

var v1WatchlistsItemsDeleteWatchlistItem = cli.Command{
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
	Action:          handleV1WatchlistsItemsDeleteWatchlistItem,
	HideHelpCommand: true,
}

func handleV1WatchlistsItemsAddWatchlistItem(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1WatchlistItemAddWatchlistItemParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Watchlists.Items.AddWatchlistItem(
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
		Title:          "v1:watchlists:items add-watchlist-item",
		Transform:      transform,
	})
}

func handleV1WatchlistsItemsDeleteWatchlistItem(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1WatchlistItemDeleteWatchlistItemParams{
		WatchlistID: cmd.Value("watchlist-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Watchlists.Items.DeleteWatchlistItem(
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
		Title:          "v1:watchlists:items delete-watchlist-item",
		Transform:      transform,
	})
}
