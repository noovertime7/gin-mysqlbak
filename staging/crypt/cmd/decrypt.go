package cmd

import (
	"fmt"
	"github.com/noovertime7/gin-mysqlbak/staging/crypt/pkg"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decrypt)
	decrypt.Flags().StringP("fileType", "t", "txt", "File suffix before encryption")
}

var decrypt = &cobra.Command{
	Use:     "decrypt",
	Aliases: []string{"de"},
	Short:   "Decrypting files using key",
	Long:    "Use 16-bit (AES-128) 24-byte (AES-192) or 32-byte (AES-256) key to decrypt the file, the key must be the same as the key of the encrypted file",
	Example: "crypt decrypt -k 0123456789abcdeasbgted3jikydj3ss -t txt text.data => text.txt",
	Run: func(cmd *cobra.Command, args []string) {
		//设置recover,
		defer func() {
			if err := recover(); err != nil { //产生了panic异常
				fmt.Println("encrypt err ", err)
			}
		}()
		if len(args) != 1 {
			HandleErr("You have entered the wrong parameter, Usage: crypt decrypt -k 0123456789abcdeasbgted3jikydj3ss -t pdf text.bak", errors.New("params error"))
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
		deFile, err := pkg.Decrypt(key, filePath, fileType)
		if err != nil {
			HandleErr("decrypt err", err)
		}
		fmt.Printf("decrypt %v success, decrypt file : %v\n", filePath, deFile)
	},
}
