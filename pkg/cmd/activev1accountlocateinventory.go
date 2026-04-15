// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/clear-street/clear-street-go"
	"github.com/clear-street/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-cli/internal/apiquery"
	"github.com/stainless-sdks/clear-street-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var activeV1AccountsLocatesInventoryGetLocateInventory = cli.Command{
	Name:    "get-locate-inventory",
	Usage:   "Retrieves available inventory for short stock locates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "account-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "symbol",
			Usage:     "The instrument symbol",
			Required:  true,
			QueryPath: "symbol",
		},
	},
	Action:          handleActiveV1AccountsLocatesInventoryGetLocateInventory,
	HideHelpCommand: true,
}

func handleActiveV1AccountsLocatesInventoryGetLocateInventory(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := clearstreet.ActiveV1AccountLocateInventoryGetLocateInventoryParams{}

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
	_, err = client.Active.V1.Accounts.Locates.Inventory.GetLocateInventory(
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
	return ShowJSON(os.Stdout, os.Stderr, "active:v1:accounts:locates:inventory get-locate-inventory", obj, format, explicitFormat, transform)
}
