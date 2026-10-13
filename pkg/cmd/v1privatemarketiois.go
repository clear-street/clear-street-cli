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

var v1PrivateMarketsIoisCreateIoi = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-ioi",
	Usage:   "Create an IOI for a visible upcoming offering.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "notional-amount",
			Required: true,
			BodyPath: "notional_amount",
		},
		&requestflag.Flag[string]{
			Name:     "offering-id",
			Required: true,
			BodyPath: "offering_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "nda-acceptance",
			Usage:    "Required only when the offering's attached SPV has an NDA agreement.",
			BodyPath: "nda_acceptance",
		},
	},
	Action:          handleV1PrivateMarketsIoisCreateIoi,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"nda-acceptance": {
		&requestflag.InnerFlag[bool]{
			Name:       "nda-acceptance.accepted",
			Usage:      "Must be true; confirms affirmative assent.",
			InnerField: "accepted",
		},
		&requestflag.InnerFlag[string]{
			Name:       "nda-acceptance.agreement-id",
			Usage:      "Exact agreement id returned by offering detail.",
			InnerField: "agreement_id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "nda-acceptance.authority-confirmed",
			Usage:      "Must be true; confirms the signer may bind the account-holder entity.",
			InnerField: "authority_confirmed",
		},
	},
})

var v1PrivateMarketsIoisDeleteIoi = cli.Command{
	Name:    "delete-ioi",
	Usage:   "Withdraw a live IOI. Repeating a withdrawal returns 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ioi-id",
			Required:  true,
			PathParam: "ioi_id",
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			QueryPath: "account_id",
		},
	},
	Action:          handleV1PrivateMarketsIoisDeleteIoi,
	HideHelpCommand: true,
}

var v1PrivateMarketsIoisGetIois = cli.Command{
	Name:    "get-iois",
	Usage:   "List every live IOI for the caller's account-holder entity.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			QueryPath: "account_id",
		},
	},
	Action:          handleV1PrivateMarketsIoisGetIois,
	HideHelpCommand: true,
}

var v1PrivateMarketsIoisUpdateIoi = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-ioi",
	Usage:   "Update an IOI's notional, accepting the current NDA revision when required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ioi-id",
			Required:  true,
			PathParam: "ioi_id",
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "notional-amount",
			Required: true,
			BodyPath: "notional_amount",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "nda-acceptance",
			Usage:    "Required when the SPV's current NDA version is newer than the IOI's\nlatest acceptance. Irrelevant acceptances are rejected.",
			BodyPath: "nda_acceptance",
		},
	},
	Action:          handleV1PrivateMarketsIoisUpdateIoi,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"nda-acceptance": {
		&requestflag.InnerFlag[bool]{
			Name:       "nda-acceptance.accepted",
			Usage:      "Must be true; confirms affirmative assent.",
			InnerField: "accepted",
		},
		&requestflag.InnerFlag[string]{
			Name:       "nda-acceptance.agreement-id",
			Usage:      "Exact agreement id returned by offering detail.",
			InnerField: "agreement_id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "nda-acceptance.authority-confirmed",
			Usage:      "Must be true; confirms the signer may bind the account-holder entity.",
			InnerField: "authority_confirmed",
		},
	},
})

func handleV1PrivateMarketsIoisCreateIoi(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1PrivateMarketIoisNewIoiParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.Iois.NewIoi(ctx, params, options...)
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
		Title:          "v1:private-markets:iois create-ioi",
		Transform:      transform,
	})
}

func handleV1PrivateMarketsIoisDeleteIoi(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ioi-id") && len(unusedArgs) > 0 {
		cmd.Set("ioi-id", unusedArgs[0])
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

	params := clearstreet.V1PrivateMarketIoisDeleteIoiParams{}

	return client.V1.PrivateMarkets.Iois.DeleteIoi(
		ctx,
		cmd.Value("ioi-id").(string),
		params,
		options...,
	)
}

func handleV1PrivateMarketsIoisGetIois(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1PrivateMarketIoisGetIoisParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.Iois.GetIois(ctx, params, options...)
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
		Title:          "v1:private-markets:iois get-iois",
		Transform:      transform,
	})
}

func handleV1PrivateMarketsIoisUpdateIoi(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ioi-id") && len(unusedArgs) > 0 {
		cmd.Set("ioi-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := clearstreet.V1PrivateMarketIoisUpdateIoiParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.Iois.UpdateIoi(
		ctx,
		cmd.Value("ioi-id").(string),
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
		Title:          "v1:private-markets:iois update-ioi",
		Transform:      transform,
	})
}
