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

var v1PrivateMarketsCreateIoi = requestflag.WithInnerFlags(cli.Command{
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
	Action:          handleV1PrivateMarketsCreateIoi,
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

var v1PrivateMarketsDeleteIoi = cli.Command{
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
	Action:          handleV1PrivateMarketsDeleteIoi,
	HideHelpCommand: true,
}

var v1PrivateMarketsGetCompanyByID = cli.Command{
	Name:    "get-company-by-id",
	Usage:   "Fetch one published private-market company with its complete versioned profile.\nRequires the account holder to have attested. Returns `404` when the company\ndoes not exist or is not yet published.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "company-id",
			Required:  true,
			PathParam: "company_id",
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Account whose account-holder entity must hold an accreditation\nattestation to browse private-market offerings.",
			Required:  true,
			QueryPath: "account_id",
		},
	},
	Action:          handleV1PrivateMarketsGetCompanyByID,
	HideHelpCommand: true,
}

var v1PrivateMarketsGetIois = cli.Command{
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
	Action:          handleV1PrivateMarketsGetIois,
	HideHelpCommand: true,
}

var v1PrivateMarketsGetSpvByID = cli.Command{
	Name:    "get-spv-by-id",
	Usage:   "Fetch one private-market SPV's complete economics and fee schedule. Requires the\naccount holder to have attested. Returns `404` unless the SPV is `OPEN` and\nattached to a currently visible `ACTIVE` offering.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "spv-id",
			Required:  true,
			PathParam: "spv_id",
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Account whose account-holder entity must hold an accreditation\nattestation to browse private-market offerings.",
			Required:  true,
			QueryPath: "account_id",
		},
	},
	Action:          handleV1PrivateMarketsGetSpvByID,
	HideHelpCommand: true,
}

var v1PrivateMarketsUpdateIoi = requestflag.WithInnerFlags(cli.Command{
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
	Action:          handleV1PrivateMarketsUpdateIoi,
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

func handleV1PrivateMarketsCreateIoi(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1PrivateMarketNewIoiParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.NewIoi(ctx, params, options...)
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
		Title:          "v1:private-markets create-ioi",
		Transform:      transform,
	})
}

func handleV1PrivateMarketsDeleteIoi(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1PrivateMarketDeleteIoiParams{}

	return client.V1.PrivateMarkets.DeleteIoi(
		ctx,
		cmd.Value("ioi-id").(string),
		params,
		options...,
	)
}

func handleV1PrivateMarketsGetCompanyByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
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

	params := clearstreet.V1PrivateMarketGetCompanyByIDParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.GetCompanyByID(
		ctx,
		cmd.Value("company-id").(string),
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
		Title:          "v1:private-markets get-company-by-id",
		Transform:      transform,
	})
}

func handleV1PrivateMarketsGetIois(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1PrivateMarketGetIoisParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.GetIois(ctx, params, options...)
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
		Title:          "v1:private-markets get-iois",
		Transform:      transform,
	})
}

func handleV1PrivateMarketsGetSpvByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("spv-id") && len(unusedArgs) > 0 {
		cmd.Set("spv-id", unusedArgs[0])
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

	params := clearstreet.V1PrivateMarketGetSpvByIDParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.GetSpvByID(
		ctx,
		cmd.Value("spv-id").(string),
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
		Title:          "v1:private-markets get-spv-by-id",
		Transform:      transform,
	})
}

func handleV1PrivateMarketsUpdateIoi(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1PrivateMarketUpdateIoiParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.UpdateIoi(
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
		Title:          "v1:private-markets update-ioi",
		Transform:      transform,
	})
}
