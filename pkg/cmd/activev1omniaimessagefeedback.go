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

var activeV1OmniAIMessagesFeedbackCreateFeedback = cli.Command{
	Name:    "create-feedback",
	Usage:   "Attaches a score and optional comment to a finalized assistant message. Feedback\nis only valid for messages with role `ASSISTANT` that have reached a terminal\noutcome.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Usage:    "Account ID for the request",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:     "score",
			Usage:    "Feedback score (-1, 0, +1 or 1-5)",
			Required: true,
			BodyPath: "score",
		},
		&requestflag.Flag[string]{
			Name:     "comment",
			Usage:    "Optional feedback comment",
			BodyPath: "comment",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Optional metadata",
			BodyPath: "metadata",
		},
	},
	Action:          handleActiveV1OmniAIMessagesFeedbackCreateFeedback,
	HideHelpCommand: true,
}

func handleActiveV1OmniAIMessagesFeedbackCreateFeedback(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIMessageFeedbackNewFeedbackParams{}

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
	_, err = client.Active.V1.OmniAI.Messages.Feedback.NewFeedback(
		ctx,
		cmd.Value("message-id").(string),
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
		Title:          "active:v1:omni-ai:messages:feedback create-feedback",
		Transform:      transform,
	})
}
