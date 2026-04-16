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

var activeV1OmniAIRunsCancelRun = cli.Command{
	Name:    "cancel-run",
	Usage:   "Cancel a running assistant run.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "run-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "Account ID for the request",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Reason for cancellation",
			BodyPath: "reason",
		},
	},
	Action:          handleActiveV1OmniAIRunsCancelRun,
	HideHelpCommand: true,
}

var activeV1OmniAIRunsGetRun = cli.Command{
	Name:    "get-run",
	Usage:   "Poll for the current status of a run and any new events since the last poll.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "run-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Account ID for the request",
			Required:  true,
			QueryPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Maximum events to return",
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Page token for incremental polling",
			QueryPath: "page_token",
		},
	},
	Action:          handleActiveV1OmniAIRunsGetRun,
	HideHelpCommand: true,
}

var activeV1OmniAIRunsStartRun = cli.Command{
	Name:    "start-run",
	Usage:   "Begins an agentic conversation run. If thread_id is provided, continues an\nexisting conversation; otherwise creates a new thread.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "Account ID for the request",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "command-text",
			Usage:    "The user's natural language command",
			Required: true,
			BodyPath: "command_text",
		},
		&requestflag.Flag[[]string]{
			Name:     "capability",
			Usage:    "Capabilities for structured actions",
			BodyPath: "capabilities",
		},
		&requestflag.Flag[any]{
			Name:     "thread-id",
			Usage:    "Optional thread ID to continue an existing conversation",
			BodyPath: "thread_id",
		},
		&requestflag.Flag[any]{
			Name:     "thread-title",
			Usage:    "Optional title for new threads",
			BodyPath: "thread_title",
		},
	},
	Action:          handleActiveV1OmniAIRunsStartRun,
	HideHelpCommand: true,
}

func handleActiveV1OmniAIRunsCancelRun(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIRunCancelRunParams{}

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
	_, err = client.Active.V1.OmniAI.Runs.CancelRun(
		ctx,
		cmd.Value("run-id").(string),
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
		Title:          "active:v1:omni-ai:runs cancel-run",
		Transform:      transform,
	})
}

func handleActiveV1OmniAIRunsGetRun(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIRunGetRunParams{}

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
	_, err = client.Active.V1.OmniAI.Runs.GetRun(
		ctx,
		cmd.Value("run-id").(string),
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
		Title:          "active:v1:omni-ai:runs get-run",
		Transform:      transform,
	})
}

func handleActiveV1OmniAIRunsStartRun(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIRunStartRunParams{}

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
	_, err = client.Active.V1.OmniAI.Runs.StartRun(ctx, params, options...)
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
		Title:          "active:v1:omni-ai:runs start-run",
		Transform:      transform,
	})
}
