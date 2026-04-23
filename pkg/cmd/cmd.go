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
				Name:     "active:v1:accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1AccountsGetAccountByID,
					&activeV1AccountsGetAccounts,
					&activeV1AccountsPatchAccountByID,
				},
			},
			{
				Name:     "active:v1:accounts:balances",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1AccountsBalancesGetAccountBalances,
				},
			},
			{
				Name:     "active:v1:accounts:orders",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1AccountsOrdersCancelAllOpenOrders,
					&activeV1AccountsOrdersCancelOpenOrder,
					&activeV1AccountsOrdersGetOrderByID,
					&activeV1AccountsOrdersGetOrders,
					&activeV1AccountsOrdersReplaceOrder,
					&activeV1AccountsOrdersSubmitOrders,
				},
			},
			{
				Name:     "active:v1:accounts:portfolio-history",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1AccountsPortfolioHistoryGetPortfolioHistory,
				},
			},
			{
				Name:     "active:v1:accounts:positions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1AccountsPositionsClosePosition,
					&activeV1AccountsPositionsClosePositions,
					&activeV1AccountsPositionsGetPositions,
				},
			},
			{
				Name:     "active:v1:api-keys",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1APIKeysCreate,
					&activeV1APIKeysList,
					&activeV1APIKeysRevoke,
					&activeV1APIKeysRevokeAll,
				},
			},
			{
				Name:     "active:v1:calendars:market-hours",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1CalendarsMarketHoursGetMarketHoursCalendar,
				},
			},
			{
				Name:     "active:v1:clock",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1ClockGetClock,
				},
			},
			{
				Name:     "active:v1:instruments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1InstrumentsGetInstrumentByID,
					&activeV1InstrumentsGetInstruments,
				},
			},
			{
				Name:     "active:v1:instruments:analyst-reporting",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1InstrumentsAnalystReportingGetInstrumentAnalystConsensus,
				},
			},
			{
				Name:     "active:v1:instruments:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1InstrumentsEventsGetAllInstrumentEvents,
					&activeV1InstrumentsEventsGetInstrumentEvents,
				},
			},
			{
				Name:     "active:v1:instruments:options:contracts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1InstrumentsOptionsContractsGetOptionContracts,
				},
			},
			{
				Name:     "active:v1:market-data:snapshot",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1MarketDataSnapshotGetSnapshots,
				},
			},
			{
				Name:     "active:v1:news",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1NewsGetNews,
				},
			},
			{
				Name:     "active:v1:omni-ai:entitlement-agreements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIEntitlementAgreementsListEntitlementAgreements,
				},
			},
			{
				Name:     "active:v1:omni-ai:entitlements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIEntitlementsCreateEntitlements,
					&activeV1OmniAIEntitlementsDeleteEntitlement,
					&activeV1OmniAIEntitlementsListEntitlements,
				},
			},
			{
				Name:     "active:v1:omni-ai:messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIMessagesGetMessage,
				},
			},
			{
				Name:     "active:v1:omni-ai:messages:feedback",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIMessagesFeedbackCreateFeedback,
				},
			},
			{
				Name:     "active:v1:omni-ai:responses",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIResponsesCancelResponse,
					&activeV1OmniAIResponsesGetResponse,
				},
			},
			{
				Name:     "active:v1:omni-ai:threads",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIThreadsCreateThread,
					&activeV1OmniAIThreadsGetThread,
					&activeV1OmniAIThreadsListThreads,
				},
			},
			{
				Name:     "active:v1:omni-ai:threads:messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIThreadsMessagesCreateMessage,
					&activeV1OmniAIThreadsMessagesListMessages,
				},
			},
			{
				Name:     "active:v1:omni-ai:threads:response",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1OmniAIThreadsResponseGetThreadResponse,
				},
			},
			{
				Name:     "active:v1:saved-screeners",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1SavedScreenersCreateScreener,
					&activeV1SavedScreenersDeleteScreener,
					&activeV1SavedScreenersGetScreenerByID,
					&activeV1SavedScreenersGetScreeners,
					&activeV1SavedScreenersReplaceScreener,
				},
			},
			{
				Name:     "active:v1:screener",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1ScreenerGetScreener,
					&activeV1ScreenerSearchScreener,
				},
			},
			{
				Name:     "active:v1:version",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1VersionGetVersion,
				},
			},
			{
				Name:     "active:v1:watchlists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1WatchlistsCreateWatchlist,
					&activeV1WatchlistsDeleteWatchlist,
					&activeV1WatchlistsGetWatchlistByID,
					&activeV1WatchlistsGetWatchlists,
				},
			},
			{
				Name:     "active:v1:watchlists:items",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1WatchlistsItemsAddWatchlistItem,
					&activeV1WatchlistsItemsDeleteWatchlistItem,
				},
			},
			{
				Name:     "active:v1:ws",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&activeV1WsWebsocketHandler,
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
