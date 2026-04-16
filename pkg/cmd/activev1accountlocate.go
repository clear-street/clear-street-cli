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

var activeV1AccountsLocatesCreateLocateRequest = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-locate-request",
	Usage:   "Submits a new short stock locate request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleActiveV1AccountsLocatesCreateLocateRequest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[int64]{
			Name:       "body.quantity",
			Usage:      "The quantity of shares to locate",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.symbol",
			Usage:      "The symbol of the security to locate",
			InnerField: "symbol",
		},
		&requestflag.InnerFlag[any]{
			Name:       "body.comments",
			Usage:      "Optional comments to associate with the locate request",
			InnerField: "comments",
		},
		&requestflag.InnerFlag[any]{
			Name:       "body.reference-id",
			Usage:      "A client-provided reference ID to identify the locate order",
			InnerField: "reference_id",
		},
	},
})

var activeV1AccountsLocatesGetLocateRequests = cli.Command{
	Name:    "get-locate-requests",
	Usage:   "Retrieves all locate requests for the specified account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Default:   1000,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next page of results. Contains encoded pagination state (limit + offset).\nWhen provided, page_size is ignored.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[string]{
			Name:      "reference-id",
			Usage:     "Filter by client reference ID",
			QueryPath: "reference_id",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "The status of a locate order",
			QueryPath: "status",
		},
	},
	Action:          handleActiveV1AccountsLocatesGetLocateRequests,
	HideHelpCommand: true,
}

var activeV1AccountsLocatesUpdateLocateRequest = cli.Command{
	Name:    "update-locate-request",
	Usage:   "Modifies an existing locate request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "accept",
			Usage:    "Whether to accept (`true`) or decline (`false`) the locate offer",
			Required: true,
			BodyPath: "accept",
		},
	},
	Action:          handleActiveV1AccountsLocatesUpdateLocateRequest,
	HideHelpCommand: true,
}

func handleActiveV1AccountsLocatesCreateLocateRequest(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountLocateNewLocateRequestParams{}

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
	_, err = client.Active.V1.Accounts.Locates.NewLocateRequest(
		ctx,
		cmd.Value("account-id").(int64),
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
		Title:          "active:v1:accounts:locates create-locate-request",
		Transform:      transform,
	})
}

func handleActiveV1AccountsLocatesGetLocateRequests(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountLocateGetLocateRequestsParams{}

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
	_, err = client.Active.V1.Accounts.Locates.GetLocateRequests(
		ctx,
		cmd.Value("account-id").(int64),
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
		Title:          "active:v1:accounts:locates get-locate-requests",
		Transform:      transform,
	})
}

func handleActiveV1AccountsLocatesUpdateLocateRequest(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountLocateUpdateLocateRequestParams{}

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
	_, err = client.Active.V1.Accounts.Locates.UpdateLocateRequest(
		ctx,
		cmd.Value("account-id").(int64),
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
		Title:          "active:v1:accounts:locates update-locate-request",
		Transform:      transform,
	})
}
