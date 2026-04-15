// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/clear-street/clear-street-cli/internal/apiquery"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
	"github.com/clear-street/clear-street-go"
	"github.com/clear-street/clear-street-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var activeV1WatchlistsCreateWatchlist = cli.Command{
	Name:    "create-watchlist",
	Usage:   "Create a new watchlist",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The desired watchlist name.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleActiveV1WatchlistsCreateWatchlist,
	HideHelpCommand: true,
}

var activeV1WatchlistsDeleteWatchlist = cli.Command{
	Name:    "delete-watchlist",
	Usage:   "Delete a watchlist and all its items",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "watchlist-id",
			Required: true,
		},
	},
	Action:          handleActiveV1WatchlistsDeleteWatchlist,
	HideHelpCommand: true,
}

var activeV1WatchlistsGetWatchlistByID = cli.Command{
	Name:    "get-watchlist-by-id",
	Usage:   "Get a watchlist by ID with all its items",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "watchlist-id",
			Required: true,
		},
	},
	Action:          handleActiveV1WatchlistsGetWatchlistByID,
	HideHelpCommand: true,
}

var activeV1WatchlistsGetWatchlists = cli.Command{
	Name:            "get-watchlists",
	Usage:           "List watchlists for the authenticated user",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleActiveV1WatchlistsGetWatchlists,
	HideHelpCommand: true,
}

func handleActiveV1WatchlistsCreateWatchlist(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1WatchlistNewWatchlistParams{}

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
	_, err = client.Active.V1.Watchlists.NewWatchlist(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:watchlists create-watchlist", obj, format, explicitFormat, transform)
}

func handleActiveV1WatchlistsDeleteWatchlist(ctx context.Context, cmd *cli.Command) error {
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

	return client.Active.V1.Watchlists.DeleteWatchlist(ctx, cmd.Value("watchlist-id").(string), options...)
}

func handleActiveV1WatchlistsGetWatchlistByID(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Active.V1.Watchlists.GetWatchlistByID(ctx, cmd.Value("watchlist-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:watchlists get-watchlist-by-id", obj, format, explicitFormat, transform)
}

func handleActiveV1WatchlistsGetWatchlists(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Active.V1.Watchlists.GetWatchlists(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:watchlists get-watchlists", obj, format, explicitFormat, transform)
}
