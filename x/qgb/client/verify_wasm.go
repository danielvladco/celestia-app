//go:build wasm

package client

import (
	"github.com/spf13/cobra"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"os"
)

func VerifyCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "verify",
		Short: "Verifies that a transaction hash, a range of shares, or a blob referenced by its transaction hash were committed to by the QGB contract",
		Run: func(cmd *cobra.Command, args []string) {
			logger := tmlog.NewTMLogger(os.Stdout)
			logger.Error("verify is not supported for wasm")
		},
	}
	return command
}
