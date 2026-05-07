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

var v1AccountsPositionsInstructionsCancelPositionInstruction = cli.Command{
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
	Action:          handleV1AccountsPositionsInstructionsCancelPositionInstruction,
	HideHelpCommand: true,
}

var v1AccountsPositionsInstructionsGetPositionInstructions = cli.Command{
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
	Action:          handleV1AccountsPositionsInstructionsGetPositionInstructions,
	HideHelpCommand: true,
}

var v1AccountsPositionsInstructionsSubmitPositionInstructions = requestflag.WithInnerFlags(cli.Command{
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
	Action:          handleV1AccountsPositionsInstructionsSubmitPositionInstructions,
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

func handleV1AccountsPositionsInstructionsCancelPositionInstruction(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1AccountPositionInstructionCancelPositionInstructionParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Positions.Instructions.CancelPositionInstruction(
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
		Title:          "v1:accounts:positions:instructions cancel-position-instruction",
		Transform:      transform,
	})
}

func handleV1AccountsPositionsInstructionsGetPositionInstructions(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1AccountPositionInstructionGetPositionInstructionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Positions.Instructions.GetPositionInstructions(
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
		Title:          "v1:accounts:positions:instructions get-position-instructions",
		Transform:      transform,
	})
}

func handleV1AccountsPositionsInstructionsSubmitPositionInstructions(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1AccountPositionInstructionSubmitPositionInstructionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Positions.Instructions.SubmitPositionInstructions(
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
		Title:          "v1:accounts:positions:instructions submit-position-instructions",
		Transform:      transform,
	})
}
