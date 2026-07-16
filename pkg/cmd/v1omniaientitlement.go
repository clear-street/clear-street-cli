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

var v1OmniAIEntitlementsCreateEntitlements = cli.Command{
	Name:    "create-entitlements",
	Usage:   "Record consent and upsert one-or-more active grants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]int64]{
			Name:     "account-id",
			Required: true,
			BodyPath: "account_ids",
		},
		&requestflag.Flag[string]{
			Name:     "agreement-id",
			Required: true,
			BodyPath: "agreement_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "entitlement-code",
			Required: true,
			BodyPath: "entitlement_codes",
		},
	},
	Action:          handleV1OmniAIEntitlementsCreateEntitlements,
	HideHelpCommand: true,
}

var v1OmniAIEntitlementsDeleteEntitlement = cli.Command{
	Name:    "delete-entitlement",
	Usage:   "Revoke one entitlement grant by id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entitlement-id",
			Required:  true,
			PathParam: "entitlement_id",
		},
	},
	Action:          handleV1OmniAIEntitlementsDeleteEntitlement,
	HideHelpCommand: true,
}

var v1OmniAIEntitlementsGetEntitlementAgreements = cli.Command{
	Name:            "get-entitlement-agreements",
	Usage:           "List current signable entitlement agreements for consent UX.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleV1OmniAIEntitlementsGetEntitlementAgreements,
	HideHelpCommand: true,
}

var v1OmniAIEntitlementsGetEntitlements = cli.Command{
	Name:    "get-entitlements",
	Usage:   "List caller's active entitlement grants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			QueryPath: "account_id",
		},
	},
	Action:          handleV1OmniAIEntitlementsGetEntitlements,
	HideHelpCommand: true,
}

func handleV1OmniAIEntitlementsCreateEntitlements(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1OmniAIEntitlementNewEntitlementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.OmniAI.Entitlements.NewEntitlements(ctx, params, options...)
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
		Title:          "v1:omni-ai:entitlements create-entitlements",
		Transform:      transform,
	})
}

func handleV1OmniAIEntitlementsDeleteEntitlement(ctx context.Context, cmd *cli.Command) error {
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
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.OmniAI.Entitlements.DeleteEntitlement(ctx, cmd.Value("entitlement-id").(string), options...)
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
		Title:          "v1:omni-ai:entitlements delete-entitlement",
		Transform:      transform,
	})
}

func handleV1OmniAIEntitlementsGetEntitlementAgreements(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.OmniAI.Entitlements.GetEntitlementAgreements(ctx, options...)
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
		Title:          "v1:omni-ai:entitlements get-entitlement-agreements",
		Transform:      transform,
	})
}

func handleV1OmniAIEntitlementsGetEntitlements(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1OmniAIEntitlementGetEntitlementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.OmniAI.Entitlements.GetEntitlements(ctx, params, options...)
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
		Title:          "v1:omni-ai:entitlements get-entitlements",
		Transform:      transform,
	})
}
