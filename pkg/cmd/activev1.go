// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/clear-street/clear-street-cli/internal/apiquery"
	"github.com/clear-street/clear-street-go"
	"github.com/urfave/cli/v3"
)

var activeV1Ws = cli.Command{
	Name:            "ws",
	Usage:           "Upgrade the HTTP connection to a WebSocket and echo incoming messages.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleActiveV1Ws,
	HideHelpCommand: true,
}

func handleActiveV1Ws(ctx context.Context, cmd *cli.Command) error {
	client := clearstreet.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	return client.Active.V1.Ws(ctx, options...)
}
