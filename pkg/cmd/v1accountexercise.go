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

var v1AccountsExercisesCancelExercise = cli.Command{
	Name:    "cancel-exercise",
	Usage:   "Cancel an outstanding exercise / DNE / CEA instruction by its server- assigned\n`id`. Returns the updated instruction with status `CANCEL_REQUESTED`; the\nterminal `CANCELLED` / `CANCEL_FAILED` state arrives asynchronously via\nsubsequent GETs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "exercise-id",
			Required:  true,
			PathParam: "exercise_id",
		},
	},
	Action:          handleV1AccountsExercisesCancelExercise,
	HideHelpCommand: true,
}

var v1AccountsExercisesGetExercises = cli.Command{
	Name:    "get-exercises",
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
			Usage:     "Filter by OEMS instrument id.",
			QueryPath: "instrument_id",
		},
	},
	Action:          handleV1AccountsExercisesGetExercises,
	HideHelpCommand: true,
}

var v1AccountsExercisesSubmitExercises = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit-exercises",
	Usage:   "Submit one or more option lifecycle instructions against the account. Each row\nis routed to `oems-csc` independently; per-row rejections are surfaced on the\ncorresponding response entry without failing the batch.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "account-id",
			Required:  true,
			PathParam: "account_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "exercise",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleV1AccountsExercisesSubmitExercises,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"exercise": {
		&requestflag.InnerFlag[string]{
			Name:       "exercise.action",
			Usage:      "The action a caller wants `oems-csc` to take against an options position.\n\nMaps onto FIX `PosTransType` (tag 709) + `PosMaintAction` (tag 712) +\n`ContraryInstructionIndicator` (tag 719) per `oems-csc`'s `classify_action`.",
			InnerField: "action",
		},
		&requestflag.InnerFlag[string]{
			Name:       "exercise.instrument-id",
			Usage:      "OEMS instrument identifier. api-gw resolves this to `security_id` +\n`security_id_source` via the instrument cache before dispatching to\n`oems-csc`. Unknown ids return 404.",
			InnerField: "instrument_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "exercise.quantity",
			Usage:      "Quantity of contracts to exercise / DNE / CEA.",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "exercise.client-exercise-id",
			Usage:      "Caller-supplied correlation id. Echoed back on the response and used\nas the FIX `pos_req_id` (tag 710) for idempotency. If omitted the\nserver generates a UUID.",
			InnerField: "client_exercise_id",
		},
	},
})

func handleV1AccountsExercisesCancelExercise(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("exercise-id") && len(unusedArgs) > 0 {
		cmd.Set("exercise-id", unusedArgs[0])
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

	params := clearstreet.V1AccountExerciseCancelExerciseParams{
		AccountID: cmd.Value("account-id").(int64),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Exercises.CancelExercise(
		ctx,
		cmd.Value("exercise-id").(string),
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
		Title:          "v1:accounts:exercises cancel-exercise",
		Transform:      transform,
	})
}

func handleV1AccountsExercisesGetExercises(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1AccountExerciseGetExercisesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Exercises.GetExercises(
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
		Title:          "v1:accounts:exercises get-exercises",
		Transform:      transform,
	})
}

func handleV1AccountsExercisesSubmitExercises(ctx context.Context, cmd *cli.Command) error {
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

	params := clearstreet.V1AccountExerciseSubmitExercisesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Accounts.Exercises.SubmitExercises(
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
		Title:          "v1:accounts:exercises submit-exercises",
		Transform:      transform,
	})
}
