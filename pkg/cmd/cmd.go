// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/clear-street/clear-street-cli/internal/autocomplete"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "clst",
		Usage:     "CLI for the clear-street API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:  "api-key",
				Usage: "A JWT issued by the authentication service.",
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "v1:accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AccountsGetAccountBalances,
					&v1AccountsGetAccountByID,
					&v1AccountsGetAccounts,
					&v1AccountsGetPortfolioHistory,
					&v1AccountsPatchAccountByID,
				},
			},
			{
				Name:     "v1:alerts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AlertsCreateAlert,
					&v1AlertsDeleteAlert,
					&v1AlertsGetAlertByID,
					&v1AlertsGetAlerts,
				},
			},
			{
				Name:     "v1:api-version",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1APIVersionGetVersion,
				},
			},
			{
				Name:     "v1:calendar",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1CalendarGetClock,
					&v1CalendarGetMarketHoursCalendar,
				},
			},
			{
				Name:     "v1:instrument-data",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentDataGetAllInstrumentEvents,
					&v1InstrumentDataGetInstrumentAnalystConsensus,
					&v1InstrumentDataGetInstrumentBalanceSheetStatements,
					&v1InstrumentDataGetInstrumentCashFlowStatements,
					&v1InstrumentDataGetInstrumentEvents,
					&v1InstrumentDataGetInstrumentFundamentals,
					&v1InstrumentDataGetInstrumentIncomeStatements,
				},
			},
			{
				Name:     "v1:instrument-data:market-data",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentDataMarketDataGetDailySummaries,
					&v1InstrumentDataMarketDataGetSnapshots,
				},
			},
			{
				Name:     "v1:instrument-data:news",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentDataNewsGetNews,
				},
			},
			{
				Name:     "v1:instruments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsGetInstrumentByID,
					&v1InstrumentsGetInstruments,
					&v1InstrumentsGetOptionContracts,
					&v1InstrumentsSearchInstruments,
				},
			},
			{
				Name:     "v1:omni-ai:entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OmniAIEntitlementsCreateEntitlements,
					&v1OmniAIEntitlementsDeleteEntitlement,
					&v1OmniAIEntitlementsGetEntitlementAgreements,
					&v1OmniAIEntitlementsGetEntitlements,
				},
			},
			{
				Name:     "v1:omni-ai:messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OmniAIMessagesGetMessageByID,
					&v1OmniAIMessagesSubmitFeedback,
				},
			},
			{
				Name:     "v1:omni-ai:responses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OmniAIResponsesCancelResponse,
					&v1OmniAIResponsesGetResponseByID,
				},
			},
			{
				Name:     "v1:omni-ai:threads",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OmniAIThreadsCreateMessage,
					&v1OmniAIThreadsCreateThread,
					&v1OmniAIThreadsGetMessages,
					&v1OmniAIThreadsGetThreadByID,
					&v1OmniAIThreadsGetThreadResponse,
					&v1OmniAIThreadsGetThreads,
				},
			},
			{
				Name:     "v1:orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OrdersCancelAllOpenOrders,
					&v1OrdersCancelOpenOrder,
					&v1OrdersGetExecutions,
					&v1OrdersGetOrderByID,
					&v1OrdersGetOrders,
					&v1OrdersReplaceOrder,
					&v1OrdersSubmitOrders,
				},
			},
			{
				Name:     "v1:positions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1PositionsCancelPositionInstruction,
					&v1PositionsClosePosition,
					&v1PositionsClosePositions,
					&v1PositionsGetPositionInstructions,
					&v1PositionsGetPositions,
					&v1PositionsSubmitPositionInstructions,
				},
			},
			{
				Name:     "v1:screener",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1ScreenerCreateScreener,
					&v1ScreenerDeleteScreener,
					&v1ScreenerGetScreenerByID,
					&v1ScreenerGetScreeners,
					&v1ScreenerReplaceScreener,
					&v1ScreenerSearchScreener,
				},
			},
			{
				Name:     "v1:watchlist",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1WatchlistAddWatchlistItem,
					&v1WatchlistCreateWatchlist,
					&v1WatchlistDeleteWatchlist,
					&v1WatchlistDeleteWatchlistItem,
					&v1WatchlistGetWatchlistByID,
					&v1WatchlistGetWatchlists,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "clst @manpages [-o clst.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "clst.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "clst.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
