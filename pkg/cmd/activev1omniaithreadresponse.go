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

var activeV1OmniAIThreadsResponseGetThreadResponse = cli.Command{
	Name:    "get-thread-response",
	Usage:   "Convenience endpoint to look up the currently active response for a thread\nwithout knowing the `response_id`. Useful when reloading a thread whose last\nfinalized message is a `USER` message — this indicates an assistant turn is\nlikely in progress.",
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
	},
	Action:          handleActiveV1OmniAIThreadsResponseGetThreadResponse,
	HideHelpCommand: true,
}

func handleActiveV1OmniAIThreadsResponseGetThreadResponse(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("thread-id") && len(unusedArgs) > 0 {
		cmd.Set("thread-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIThreadResponseGetThreadResponseParams{}

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
	_, err = client.Active.V1.OmniAI.Threads.Response.GetThreadResponse(
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
		Title:          "active:v1:omni-ai:threads:response get-thread-response",
		Transform:      transform,
	})
}
