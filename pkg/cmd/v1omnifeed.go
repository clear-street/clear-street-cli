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

var v1OmniFeedGetFeed = cli.Command{
	Name:    "get-feed",
	Usage:   "> **Alpha** — this endpoint is experimental and may change or be removed at any\n> time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Trading account to serve as context. Optional — the feed works\nwithout one.",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Id of the last item already received; the response continues from\nthe item after it. Omit to resume from the oldest unseen item.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return (1–100, default 20).",
			QueryPath: "limit",
		},
	},
	Action:          handleV1OmniFeedGetFeed,
	HideHelpCommand: true,
}

var v1OmniFeedPostFeedEvent = requestflag.WithInnerFlags(cli.Command{
	Name:    "post-feed-event",
	Usage:   "> **Alpha** — this endpoint is experimental and may change or be removed at any\n> time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "event",
			Usage:    "The event to record.",
			Required: true,
			BodyPath: "event",
		},
	},
	Action:          handleV1OmniFeedPostFeedEvent,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"event": {
		&requestflag.InnerFlag[string]{
			Name:       "event.item-id",
			Usage:      "The feed item the event concerns.",
			InnerField: "item_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.type",
			Usage:      "What happened.",
			InnerField: "type",
		},
	},
})

func handleV1OmniFeedGetFeed(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1OmniFeedGetFeedParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.OmniFeed.GetFeed(ctx, params, options...)
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
		Title:          "v1:omni-feed get-feed",
		Transform:      transform,
	})
}

func handleV1OmniFeedPostFeedEvent(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1OmniFeedPostFeedEventParams{}

	return client.V1.OmniFeed.PostFeedEvent(ctx, params, options...)
}
