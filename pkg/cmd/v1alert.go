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

var v1AlertsCreateAlert = cli.Command{
	Name:    "create-alert",
	Usage:   "Create an alert that watches a market or portfolio condition on the account and\nnotifies when it triggers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "condition",
			Usage:    "The boolean condition tree, in the condition grammar.\n`\"instrument_id\"` references accept a ticker or an OEMS instrument id.",
			Required: true,
			BodyPath: "condition",
		},
		&requestflag.Flag[string]{
			Name:     "schedule",
			Usage:    "How often an alert's condition is evaluated.",
			Required: true,
			BodyPath: "schedule",
		},
		&requestflag.Flag[string]{
			Name:     "trigger",
			Usage:    "How an alert triggers. `once` alerts complete after their first trigger.",
			Required: true,
			BodyPath: "trigger",
		},
		&requestflag.Flag[*int64]{
			Name:     "account-id",
			Usage:    "The account whose `account.*` signals and holdings scopes the\ncondition reads. Optional: a market-only alert needs no account.",
			BodyPath: "account_id",
		},
	},
	Action:          handleV1AlertsCreateAlert,
	HideHelpCommand: true,
}

var v1AlertsDeleteAlert = cli.Command{
	Name:    "delete-alert",
	Usage:   "Delete an alert. It stops evaluating and disappears from this API; its trigger\nhistory is retained server-side.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "alert-id",
			Required:  true,
			PathParam: "alert_id",
		},
	},
	Action:          handleV1AlertsDeleteAlert,
	HideHelpCommand: true,
}

var v1AlertsGetAlertByID = cli.Command{
	Name:    "get-alert-by-id",
	Usage:   "Get one alert by id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "alert-id",
			Required:  true,
			PathParam: "alert_id",
		},
	},
	Action:          handleV1AlertsGetAlertByID,
	HideHelpCommand: true,
}

var v1AlertsGetAlerts = cli.Command{
	Name:    "get-alerts",
	Usage:   "List the caller's alerts, newest first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The number of items to return per page. Only used when page_token is not provided.",
			Default:   100,
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "page-token",
			Usage:     "Token for retrieving the next or previous page of results. Contains encoded pagination state; when provided, page_size is ignored.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[*string]{
			Name:      "status",
			Usage:     "Comma-separated status filter (`active`, `paused`, `completed`,\n`expired`). Unknown values are rejected. Absent = every status.",
			QueryPath: "status",
		},
	},
	Action:          handleV1AlertsGetAlerts,
	HideHelpCommand: true,
}

func handleV1AlertsCreateAlert(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1AlertNewAlertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Alerts.NewAlert(ctx, params, options...)
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
		Title:          "v1:alerts create-alert",
		Transform:      transform,
	})
}

func handleV1AlertsDeleteAlert(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("alert-id") && len(unusedArgs) > 0 {
		cmd.Set("alert-id", unusedArgs[0])
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

	return client.V1.Alerts.DeleteAlert(ctx, cmd.Value("alert-id").(string), options...)
}

func handleV1AlertsGetAlertByID(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("alert-id") && len(unusedArgs) > 0 {
		cmd.Set("alert-id", unusedArgs[0])
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
	_, err = client.V1.Alerts.GetAlertByID(ctx, cmd.Value("alert-id").(string), options...)
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
		Title:          "v1:alerts get-alert-by-id",
		Transform:      transform,
	})
}

func handleV1AlertsGetAlerts(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1AlertGetAlertsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Alerts.GetAlerts(ctx, params, options...)
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
		Title:          "v1:alerts get-alerts",
		Transform:      transform,
	})
}
