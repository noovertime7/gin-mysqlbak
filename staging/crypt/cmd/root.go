package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypt",
	Short: "An encryption and decryption tool using AES",
	Long:  `Encrypt or decrypt files using the GCM algorithm in AES encryption`,
}

func init() {
	rootCmd.PersistentFlags().StringP("key", "k", "0123456789abcdeasbgted3jikydj3ss", " crypt -k 0123456789abcdeasbgted3jikydj3ss ")
	if err := rootCmd.MarkPersistentFlagRequired("key"); err != nil {
		HandleErr("MarkPersistentFlagRequired err", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		HandleErr("Execute err", err)
	}
}

func HandleErr(msg string, err error) {
	fmt.Printf("%s:%v\n", msg, err)
	os.Exit(1)
}
