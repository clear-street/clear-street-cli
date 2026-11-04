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

var v1OmniAIResponsesCancelResponse = cli.Command{
	Name:    "cancel-response",
	Usage:   "Cancel a queued or running response. Canceling a running response (Stop) also\ncancels every queued response in the thread that answers a user message.\nCanceling a queued response (withdraw) cancels only that response; the running\nresponse and other queued responses continue. Cancellation is idempotent after\nthe response becomes terminal. Each canceled response still produces a finalized\nassistant message with outcome `canceled` in the thread history. Canceling a\nresponse that ended `steered` or `merged` follows its `continued_in_response_id`\nand stops the response that continues its work, with the same effect as Stop.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "response-id",
			Required:  true,
			PathParam: "response_id",
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Lists only conversations for this account, or unlinked conversations when omitted.\nOther reads authorize the resource's linked account.\nOmit when no account is selected; empty values and the string null are invalid.",
			QueryPath: "account_id",
		},
	},
	Action:          handleV1OmniAIResponsesCancelResponse,
	HideHelpCommand: true,
}

var v1OmniAIResponsesGetResponseByID = cli.Command{
	Name:    "get-response-by-id",
	Usage:   "Poll the current snapshot of an in-progress or completed assistant response.\nWhile its status is `queued` or `running`, content may be partial and include\nthinking parts. Continue polling until it becomes `succeeded`, `failed`,\n`canceled`, `steered`, or `merged`. Treat an unknown status as terminal. On\n`steered` or `merged`, poll the response in `continued_in_response_id` at once:\nit continues the work.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "response-id",
			Required:  true,
			PathParam: "response_id",
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Lists only conversations for this account, or unlinked conversations when omitted.\nOther reads authorize the resource's linked account.\nOmit when no account is selected; empty values and the string null are invalid.",
			QueryPath: "account_id",
		},
	},
	Action:          handleV1OmniAIResponsesGetResponseByID,
	HideHelpCommand: true,
}

func handleV1OmniAIResponsesCancelResponse(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("response-id") && len(unusedArgs) > 0 {
		cmd.Set("response-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := clearstreet.V1OmniAIResponseCancelResponseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.OmniAI.Responses.CancelResponse(
		ctx,
		cmd.Value("response-id").(string),
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
		Title:          "v1:omni-ai:responses cancel-response",
		Transform:      transform,
	})
}

func handleV1OmniAIResponsesGetResponseByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("response-id") && len(unusedArgs) > 0 {
		cmd.Set("response-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := clearstreet.V1OmniAIResponseGetResponseByIDParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.OmniAI.Responses.GetResponseByID(
		ctx,
		cmd.Value("response-id").(string),
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
		Title:          "v1:omni-ai:responses get-response-by-id",
		Transform:      transform,
	})
}
