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

var activeV1OmniAIEntitlementsCreateEntitlements = cli.Command{
	Name:    "create-entitlements",
	Usage:   "Record consent and upsert one-or-more active grants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agreement-id",
			Required: true,
			BodyPath: "agreement_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "requested-entitlement-code",
			Required: true,
			BodyPath: "requested_entitlement_codes",
		},
		&requestflag.Flag[[]int64]{
			Name:     "trading-account-id",
			Required: true,
			BodyPath: "trading_account_ids",
		},
	},
	Action:          handleActiveV1OmniAIEntitlementsCreateEntitlements,
	HideHelpCommand: true,
}

var activeV1OmniAIEntitlementsDeleteEntitlement = cli.Command{
	Name:    "delete-entitlement",
	Usage:   "Revoke one entitlement grant by id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entitlement-id",
			Required: true,
		},
	},
	Action:          handleActiveV1OmniAIEntitlementsDeleteEntitlement,
	HideHelpCommand: true,
}

var activeV1OmniAIEntitlementsListEntitlements = cli.Command{
	Name:    "list-entitlements",
	Usage:   "List caller's active entitlement grants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "trading-account-id",
			QueryPath: "trading_account_id",
		},
	},
	Action:          handleActiveV1OmniAIEntitlementsListEntitlements,
	HideHelpCommand: true,
}

func handleActiveV1OmniAIEntitlementsCreateEntitlements(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIEntitlementNewEntitlementsParams{}

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
	_, err = client.Active.V1.OmniAI.Entitlements.NewEntitlements(ctx, params, options...)
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
		Title:          "active:v1:omni-ai:entitlements create-entitlements",
		Transform:      transform,
	})
}

func handleActiveV1OmniAIEntitlementsDeleteEntitlement(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entitlement-id") && len(unusedArgs) > 0 {
		cmd.Set("entitlement-id", unusedArgs[0])
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
	_, err = client.Active.V1.OmniAI.Entitlements.DeleteEntitlement(ctx, cmd.Value("entitlement-id").(string), options...)
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
		Title:          "active:v1:omni-ai:entitlements delete-entitlement",
		Transform:      transform,
	})
}

func handleActiveV1OmniAIEntitlementsListEntitlements(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1OmniAIEntitlementListEntitlementsParams{}

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
	_, err = client.Active.V1.OmniAI.Entitlements.ListEntitlements(ctx, params, options...)
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
		Title:          "active:v1:omni-ai:entitlements list-entitlements",
		Transform:      transform,
	})
}
