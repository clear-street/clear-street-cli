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

var activeV1WatchlistsItemsAddWatchlistItem = cli.Command{
	Name:    "add-watchlist-item",
	Usage:   "Add an instrument to a watchlist",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "watchlist-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "instrument-id",
			Usage:    "OEMS instrument ID (mutually exclusive with security_id/security_id_source)",
			BodyPath: "instrument_id",
		},
		&requestflag.Flag[any]{
			Name:     "security-id",
			Usage:    "Security identifier",
			BodyPath: "security_id",
		},
		&requestflag.Flag[string]{
			Name:     "security-id-source",
			Usage:    "Security identifier source",
			BodyPath: "security_id_source",
		},
	},
	Action:          handleActiveV1WatchlistsItemsAddWatchlistItem,
	HideHelpCommand: true,
}

var activeV1WatchlistsItemsDeleteWatchlistItem = cli.Command{
	Name:    "delete-watchlist-item",
	Usage:   "Delete an instrument from a watchlist",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "watchlist-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "item-id",
			Required: true,
		},
	},
	Action:          handleActiveV1WatchlistsItemsDeleteWatchlistItem,
	HideHelpCommand: true,
}

func handleActiveV1WatchlistsItemsAddWatchlistItem(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("watchlist-id") && len(unusedArgs) > 0 {
		cmd.Set("watchlist-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1WatchlistItemAddWatchlistItemParams{}

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
	_, err = client.Active.V1.Watchlists.Items.AddWatchlistItem(
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
		Title:          "active:v1:watchlists:items add-watchlist-item",
		Transform:      transform,
	})
}

func handleActiveV1WatchlistsItemsDeleteWatchlistItem(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("item-id") && len(unusedArgs) > 0 {
		cmd.Set("item-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1WatchlistItemDeleteWatchlistItemParams{
		WatchlistID: cmd.Value("watchlist-id").(string),
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

	return client.Active.V1.Watchlists.Items.DeleteWatchlistItem(
		ctx,
		cmd.Value("item-id").(string),
		params,
		options...,
	)
}
