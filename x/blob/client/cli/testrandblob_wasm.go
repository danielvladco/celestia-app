//go:build wasm

package cli

import (
	"log"

	"github.com/spf13/cobra"
)

// CmdTestRandBlob is triggered by testground's tests as part of apps' node scenario
// to increase the block size by user-defined amount.
//
// CAUTION: This func should not be used in production env!
func CmdTestRandBlob() *cobra.Command {
	return &cobra.Command{
		Use:   "TestRandBlob [blobSize]",
		Short: "Generates a random blob for a random namespace to be published to the Celestia blockchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("CLI not supported for wasm")
			return nil
		},
	}
}
