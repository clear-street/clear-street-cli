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

var activeV1OmniAIFeedbackCreateFeedback = cli.Command{
	Name:    "create-feedback",
	Usage:   "Submit user feedback (thumbs up/down, rating, comment) for an assistant message.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "Account ID for the request",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "message-id",
			Usage:    "Message to provide feedback on",
			Required: true,
			BodyPath: "message_id",
		},
		&requestflag.Flag[int64]{
			Name:     "score",
			Usage:    "Feedback score (-1, 0, +1 or 1-5)",
			Required: true,
			BodyPath: "score",
		},
		&requestflag.Flag[string]{
			Name:     "thread-id",
			Usage:    "Thread containing the message",
			Required: true,
			BodyPath: "thread_id",
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
	Action:          handleActiveV1OmniAIFeedbackCreateFeedback,
	HideHelpCommand: true,
}

func handleActiveV1OmniAIFeedbackCreateFeedback(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIFeedbackNewFeedbackParams{}

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
	_, err = client.Active.V1.OmniAI.Feedback.NewFeedback(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:omni-ai:feedback create-feedback", obj, format, explicitFormat, transform)
}
