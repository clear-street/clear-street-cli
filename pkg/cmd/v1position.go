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

var v1PositionsCancelPositionInstruction = cli.Command{
	Name:    "cancel-position-instruction",
	Usage:   "Cancel an outstanding exercise / DNE / CEA instruction by its server- assigned\n`id`. Returns the updated instruction with status `CANCEL_REQUESTED`; the\nterminal `CANCELLED` / `CANCEL_FAILED` state arrives asynchronously via\nsubsequent GETs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "instruction-id",
			Required:  true,
			PathParam: "instruction_id",
		},
	},
	Action:          handleV1PositionsCancelPositionInstruction,
	HideHelpCommand: true,
}

var v1PositionsClosePosition = cli.Command{
	Name:    "close-position",
	Usage:   "Delete a position within an account for an instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			Required:  true,
			PathParam: "instrument_id",
		},
		&requestflag.Flag[*bool]{
			Name:     "cancel-orders",
			BodyPath: "cancel_orders",
		},
	},
	Action:          handleV1PositionsClosePosition,
	HideHelpCommand: true,
}

var v1PositionsClosePositions = cli.Command{
	Name:    "close-positions",
	Usage:   "Delete all positions within an account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[*bool]{
			Name:     "cancel-orders",
			BodyPath: "cancel_orders",
		},
	},
	Action:          handleV1PositionsClosePositions,
	HideHelpCommand: true,
}

var v1PositionsGetPositionInstructions = cli.Command{
	Name:    "get-position-instructions",
	Usage:   "Returns the current lifecycle state of exercise / DNE / CEA instructions for the\naccount. Optionally filter by a specific instrument.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "instrument-id",
			Usage:     "OEMS instrument UUID",
			QueryPath: "instrument_id",
		},
	},
	Action:          handleV1PositionsGetPositionInstructions,
	HideHelpCommand: true,
}

var v1PositionsGetPositions = cli.Command{
	Name:    "get-positions",
	Usage:   "Retrieves all positions for the specified trading account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "instrument-id",
			Usage:     "Comma-separated OEMS instrument UUIDs",
			QueryPath: "instrument_ids",
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
			Name:      "sort-by",
			Usage:     "Field to sort by",
			QueryPath: "sort_by",
		},
		&requestflag.Flag[string]{
			Name:      "sort-direction",
			Usage:     "Sort direction",
			QueryPath: "sort_direction",
		},
	},
	Action:          handleV1PositionsGetPositions,
	HideHelpCommand: true,
}

var v1PositionsSubmitPositionInstructions = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit-position-instructions",
	Usage:   "Submit one or more option lifecycle instructions against the account. Each row\nis routed to `oems-csc` independently; per-row rejections are surfaced on the\ncorresponding response entry without failing the batch.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "instruction",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleV1PositionsSubmitPositionInstructions,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"instruction": {
		&requestflag.InnerFlag[string]{
			Name:       "instruction.instruction-type",
			Usage:      "The instruction type a caller wants `oems-csc` to take against an options position.\n\nMaps onto FIX `PosTransType` (tag 709) + `PosMaintAction` (tag 712) +\n`ContraryInstructionIndicator` (tag 719) per `oems-csc`'s `classify_action`.",
			InnerField: "instruction_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "instruction.instrument-id",
			Usage:      "OEMS instrument identifier. api-gw resolves this to `security_id` +\n`security_id_source` via the instrument cache before dispatching to\n`oems-csc`. Unknown ids return 404.",
			InnerField: "instrument_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "instruction.quantity",
			Usage:      "Quantity of contracts to exercise / DNE / CEA.",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "instruction.instruction-id",
			Usage:      "Caller-supplied instruction id. Echoed back on the response and used\nas the FIX `pos_req_id` (tag 710) for idempotency. If omitted the\nserver generates a UUID.",
			InnerField: "instruction_id",
		},
	},
})

func handleV1PositionsCancelPositionInstruction(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instruction-id") && len(unusedArgs) > 0 {
		cmd.Set("instruction-id", unusedArgs[0])
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

	params := clearstreet.V1PositionCancelPositionInstructionParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Positions.CancelPositionInstruction(
		ctx,
		cmd.Value("instruction-id").(string),
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
		Title:          "v1:positions cancel-position-instruction",
		Transform:      transform,
	})
}

func handleV1PositionsClosePosition(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instrument-id") && len(unusedArgs) > 0 {
		cmd.Set("instrument-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := clearstreet.V1PositionClosePositionParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Positions.ClosePosition(
		ctx,
		clearstreet.InstrumentIDOrSymbol(cmd.Value("instrument-id").(string)),
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
		Title:          "v1:positions close-position",
		Transform:      transform,
	})
}

func handleV1PositionsClosePositions(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := clearstreet.V1PositionClosePositionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Positions.ClosePositions(
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:positions close-positions",
		Transform:      transform,
	})
}

func handleV1PositionsGetPositionInstructions(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
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

	params := clearstreet.V1PositionGetPositionInstructionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Positions.GetPositionInstructions(
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:positions get-position-instructions",
		Transform:      transform,
	})
}

func handleV1PositionsGetPositions(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
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

	params := clearstreet.V1PositionGetPositionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Positions.GetPositions(
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:positions get-positions",
		Transform:      transform,
	})
}

func handleV1PositionsSubmitPositionInstructions(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := clearstreet.V1PositionSubmitPositionInstructionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Positions.SubmitPositionInstructions(
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:positions submit-position-instructions",
		Transform:      transform,
	})
}
