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

var activeV1IrisRunsCancelRunDeprecated = cli.Command{
	Name:    "cancel-run-deprecated",
	Usage:   "**Deprecated**: Use `DELETE /omni-ai/runs/{run_id}` instead.",
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
	Action:          handleActiveV1IrisRunsCancelRunDeprecated,
	HideHelpCommand: true,
}

var activeV1IrisRunsGetRunDeprecated = cli.Command{
	Name:    "get-run-deprecated",
	Usage:   "**Deprecated**: Use `GET /omni-ai/runs/{run_id}` instead.",
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
	Action:          handleActiveV1IrisRunsGetRunDeprecated,
	HideHelpCommand: true,
}

var activeV1IrisRunsStartRunDeprecated = cli.Command{
	Name:    "start-run-deprecated",
	Usage:   "**Deprecated**: Use `POST /omni-ai/runs` instead.",
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
	Action:          handleActiveV1IrisRunsStartRunDeprecated,
	HideHelpCommand: true,
}

func handleActiveV1IrisRunsCancelRunDeprecated(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1IrisRunCancelRunDeprecatedParams{}

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
	_, err = client.Active.V1.Iris.Runs.CancelRunDeprecated(
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
		Title:          "active:v1:iris:runs cancel-run-deprecated",
		Transform:      transform,
	})
}

func handleActiveV1IrisRunsGetRunDeprecated(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1IrisRunGetRunDeprecatedParams{}

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
	_, err = client.Active.V1.Iris.Runs.GetRunDeprecated(
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
		Title:          "active:v1:iris:runs get-run-deprecated",
		Transform:      transform,
	})
}

func handleActiveV1IrisRunsStartRunDeprecated(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1IrisRunStartRunDeprecatedParams{}

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
	_, err = client.Active.V1.Iris.Runs.StartRunDeprecated(ctx, params, options...)
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
		Title:          "active:v1:iris:runs start-run-deprecated",
		Transform:      transform,
	})
}
