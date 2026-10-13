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
