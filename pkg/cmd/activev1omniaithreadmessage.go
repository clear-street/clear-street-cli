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

var activeV1OmniAIThreadsMessagesCreateMessage = cli.Command{
	Name:    "create-message",
	Usage:   "Appends a new user message to the thread and starts an assistant response. Only\none response may be active per thread at a time — if the previous turn is still\nin progress, this endpoint returns **409 Conflict**. Wait for the active\nresponse to reach a terminal status before submitting the next turn.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "thread-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[[]string]{
			Name:     "capability",
			BodyPath: "capabilities",
		},
	},
	Action:          handleActiveV1OmniAIThreadsMessagesCreateMessage,
	HideHelpCommand: true,
}

var activeV1OmniAIThreadsMessagesListMessages = cli.Command{
	Name:    "list-messages",
	Usage:   "Returns **finalized** messages in chronological order. Messages from in-progress\nassistant turns are excluded — use `GET /omni-ai/threads/{thread_id}/response`\nor `GET /omni-ai/responses/{response_id}` for live output.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "thread-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Account ID for the request",
			Required:  true,
			QueryPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   100,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next page of results. Contains encoded pagination state (limit + offset).\nWhen provided, page_size is ignored.",
			QueryPath: "page_token",
		},
	},
	Action:          handleActiveV1OmniAIThreadsMessagesListMessages,
	HideHelpCommand: true,
}

func handleActiveV1OmniAIThreadsMessagesCreateMessage(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("thread-id") && len(unusedArgs) > 0 {
		cmd.Set("thread-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIThreadMessageNewMessageParams{}

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
	_, err = client.Active.V1.OmniAI.Threads.Messages.NewMessage(
		ctx,
		cmd.Value("thread-id").(string),
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
		Title:          "active:v1:omni-ai:threads:messages create-message",
		Transform:      transform,
	})
}

func handleActiveV1OmniAIThreadsMessagesListMessages(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("thread-id") && len(unusedArgs) > 0 {
		cmd.Set("thread-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIThreadMessageListMessagesParams{}

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
	_, err = client.Active.V1.OmniAI.Threads.Messages.ListMessages(
		ctx,
		cmd.Value("thread-id").(string),
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
		Title:          "active:v1:omni-ai:threads:messages list-messages",
		Transform:      transform,
	})
}
