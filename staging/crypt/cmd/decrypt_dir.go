package cmd

import (
	"fmt"
	"github.com/noovertime7/gin-mysqlbak/staging/crypt/pkg"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(decryptDir)
	decryptDir.Flags().StringP("fileType", "t", "txt", "File suffix before encryption")
}

var decryptDir = &cobra.Command{
	Use:     "decryptdir",
	Aliases: []string{"dedir"},
	Short:   "Decrypting Folder using key",
	Long:    "Use 16-bit (AES-128) 24-byte (AES-192) or 32-byte (AES-256) key to decrypt the file, the key must be the same as the key of the encrypted file",
	Example: "crypt decryptdir -k 0123456789abcdeasbgted3jikydj3ss -t sql {{Your_FolderName}} ",
	Run: func(cmd *cobra.Command, args []string) {
		//设置recover,
		defer func() {
			if err := recover(); err != nil { //产生了panic异常
				fmt.Println("encrypt err ", err)
			}
		}()
		if len(args) != 1 {
			HandleErr("You have entered the wrong parameter, Usage: crypt decryptdir -k 0123456789abcdeasbgted3jikydj3ss -t sql {{Your_FolderName}} ", errors.New("params error"))
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
		if !isDir(filePath) {
			fmt.Printf("decrypt err, %s is not a directory")
			return
		}
		if err := pkg.DecryptDir(key, filePath, fileType); err != nil {
			HandleErr("decrypt err", err)
		}
	},
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		HandleErr("path err", err)
	}
	return s.IsDir()
}
