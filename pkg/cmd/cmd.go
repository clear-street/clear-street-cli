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
		Name:      "clear-street",
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
				Name:     "v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1WebsocketHandler,
				},
			},
			{
				Name:     "v1:accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AccountsGetAccountByID,
					&v1AccountsGetAccounts,
					&v1AccountsPatchAccountByID,
				},
			},
			{
				Name:     "v1:accounts:balances",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AccountsBalancesGetAccountBalances,
				},
			},
			{
				Name:     "v1:accounts:exercises",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AccountsExercisesCancelExercise,
					&v1AccountsExercisesGetExercises,
					&v1AccountsExercisesSubmitExercises,
				},
			},
			{
				Name:     "v1:accounts:orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AccountsOrdersCancelAllOpenOrders,
					&v1AccountsOrdersCancelOpenOrder,
					&v1AccountsOrdersGetOrderByID,
					&v1AccountsOrdersGetOrders,
					&v1AccountsOrdersReplaceOrder,
					&v1AccountsOrdersSubmitOrders,
				},
			},
			{
				Name:     "v1:accounts:portfolio-history",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AccountsPortfolioHistoryGetPortfolioHistory,
				},
			},
			{
				Name:     "v1:accounts:positions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1AccountsPositionsClosePosition,
					&v1AccountsPositionsClosePositions,
					&v1AccountsPositionsGetPositions,
				},
			},
			{
				Name:     "v1:calendars:market-hours",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1CalendarsMarketHoursGetMarketHoursCalendar,
				},
			},
			{
				Name:     "v1:clock",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1ClockGetClock,
				},
			},
			{
				Name:     "v1:instruments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsGetInstrumentByID,
					&v1InstrumentsGetInstruments,
					&v1InstrumentsSearchInstruments,
				},
			},
			{
				Name:     "v1:instruments:analyst-reporting",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsAnalystReportingGetInstrumentAnalystConsensus,
				},
			},
			{
				Name:     "v1:instruments:balance-sheets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsBalanceSheetsGetInstrumentBalanceSheetStatements,
				},
			},
			{
				Name:     "v1:instruments:cash-flow-statements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsCashFlowStatementsGetInstrumentCashFlowStatements,
				},
			},
			{
				Name:     "v1:instruments:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsEventsGetAllInstrumentEvents,
					&v1InstrumentsEventsGetInstrumentEvents,
				},
			},
			{
				Name:     "v1:instruments:fundamentals",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsFundamentalsGetInstrumentFundamentals,
				},
			},
			{
				Name:     "v1:instruments:income-statements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsIncomeStatementsGetInstrumentIncomeStatements,
				},
			},
			{
				Name:     "v1:instruments:options",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1InstrumentsOptionsGetOptionContracts,
				},
			},
			{
				Name:     "v1:market-data:daily-summary",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1MarketDataDailySummaryGetDailySummaries,
				},
			},
			{
				Name:     "v1:market-data:snapshot",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1MarketDataSnapshotGetSnapshots,
				},
			},
			{
				Name:     "v1:news",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1NewsGetNews,
				},
			},
			{
				Name:     "v1:omni-ai:entitlement-agreements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OmniAIEntitlementAgreementsGetEntitlementAgreements,
				},
			},
			{
				Name:     "v1:omni-ai:entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OmniAIEntitlementsCreateEntitlements,
					&v1OmniAIEntitlementsDeleteEntitlement,
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
					&v1OmniAIThreadsCreateThread,
					&v1OmniAIThreadsGetThreadByID,
					&v1OmniAIThreadsGetThreadResponse,
					&v1OmniAIThreadsGetThreads,
				},
			},
			{
				Name:     "v1:omni-ai:threads:messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1OmniAIThreadsMessagesCreateMessage,
					&v1OmniAIThreadsMessagesGetMessages,
				},
			},
			{
				Name:     "v1:saved-screeners",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1SavedScreenersCreateScreener,
					&v1SavedScreenersDeleteScreener,
					&v1SavedScreenersGetScreenerByID,
					&v1SavedScreenersGetScreeners,
					&v1SavedScreenersReplaceScreener,
				},
			},
			{
				Name:     "v1:screener",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1ScreenerGetScreener,
					&v1ScreenerSearchScreener,
				},
			},
			{
				Name:     "v1:version",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1VersionGetVersion,
				},
			},
			{
				Name:     "v1:watchlists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1WatchlistsCreateWatchlist,
					&v1WatchlistsDeleteWatchlist,
					&v1WatchlistsGetWatchlistByID,
					&v1WatchlistsGetWatchlists,
				},
			},
			{
				Name:     "v1:watchlists:items",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&v1WatchlistsItemsAddWatchlistItem,
					&v1WatchlistsItemsDeleteWatchlistItem,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "clear-street @manpages [-o clear-street.1] [--gzip]",
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
		file, err := os.Create(filepath.Join(dir, "man1", "clear-street.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "clear-street.1.gz"))
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
