package cmd

import (
	"crypto/md5"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "file-compare [...urls]",
	Short: `通过获取给定的URL文件内容判定其MD5值是否一致。`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// 如果没有提供URL参数则在终端提示
			return fmt.Errorf("请提供两个文件的URL地址")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			duration time.Duration
			body     []byte
			err      error
		)
		for _, url := range args {
			if duration, body, err = get(url); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("%v %x\n", duration, md5.Sum(body))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
