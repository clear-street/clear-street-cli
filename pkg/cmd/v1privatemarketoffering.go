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

var v1PrivateMarketsOfferingsGetOfferingByID = cli.Command{
	Name:    "get-offering-by-id",
	Usage:   "Fetch one visible private-market offering with its documents, participants, and\nany attached SPV. Requires the account holder to have attested. Returns `404`\nwhen the offering does not exist or is not currently visible.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "offering-id",
			Required:  true,
			PathParam: "offering_id",
		},
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Account whose account-holder entity must hold an accreditation\nattestation to browse private-market offerings.",
			Required:  true,
			QueryPath: "account_id",
		},
	},
	Action:          handleV1PrivateMarketsOfferingsGetOfferingByID,
	HideHelpCommand: true,
}

var v1PrivateMarketsOfferingsGetOfferings = cli.Command{
	Name:    "get-offerings",
	Usage:   "List every visible private-market offering as a card, with its derived class,\ncompany and SPV identity, and indicative terms. Requires the account holder to\nhave attested.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Usage:     "Account whose account-holder entity must hold an accreditation\nattestation to browse private-market offerings.",
			Required:  true,
			QueryPath: "account_id",
		},
	},
	Action:          handleV1PrivateMarketsOfferingsGetOfferings,
	HideHelpCommand: true,
}

func handleV1PrivateMarketsOfferingsGetOfferingByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("offering-id") && len(unusedArgs) > 0 {
		cmd.Set("offering-id", unusedArgs[0])
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

	params := clearstreet.V1PrivateMarketOfferingGetOfferingByIDParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.Offerings.GetOfferingByID(
		ctx,
		cmd.Value("offering-id").(string),
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
		Title:          "v1:private-markets:offerings get-offering-by-id",
		Transform:      transform,
	})
}

func handleV1PrivateMarketsOfferingsGetOfferings(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1PrivateMarketOfferingGetOfferingsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.PrivateMarkets.Offerings.GetOfferings(ctx, params, options...)
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
		Title:          "v1:private-markets:offerings get-offerings",
		Transform:      transform,
	})
}
