package cmd

import (
	"fmt"
	"github.com/noovertime7/gin-mysqlbak/staging/crypt/pkg"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encrypt)
	encrypt.Flags().StringP("fileType", "t", "txt", "File suffix after encryption")
}

var encrypt = &cobra.Command{
	Use:     "encrypt",
	Aliases: []string{"en"},
	Short:   "Encrypting files using key",
	Long:    "Use 16-bit (AES-128) 24-byte (AES-192) or 32-byte (AES-256) key to encrypt the file, the key must be the same as the key of the encrypted file",
	Example: "crypt encrypt -k 0123456789abcdeasbgted3jikydj3ss -t data text.txt => text.data",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			HandleErr("You have entered the wrong parameter, Usage: crypt encrypt -k 0123456789abcdeasbgted3jikydj3ss -t data text.txt", errors.New("params error"))
		}
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			HandleErr("read key err", err)
		}
		fileType, err := cmd.Flags().GetString("fileType")
		if err != nil {
			HandleErr("read key err", err)
		}
		filePath := args[0]
		enFile, err := pkg.Encrypt(key, filePath, fileType)
		if err != nil {
			HandleErr("encrypt err", err)
		}
		fmt.Printf("encrypt %v success, encrypt file : %v\n", filePath, enFile)
	},
}
