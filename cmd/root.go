package cmd

import (
	"crypto/md5"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
	"time"
	"unicode/utf8"
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
			duration  time.Duration
			body      []byte
			size      string
			md5_value string
			// t        table.Table
			err error
		)

		t := table.NewWriter()
		t.AppendHeader(table.Row{"#", "MD5", "请求耗时", "文件大小", "请求地址"})

		for index, url := range args {
			if duration, body, err = get(url); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			md5_value = fmt.Sprintf("%x", md5.Sum(body))
			size = ByteCountDecimal(utf8.RuneCountInString(string(body)))

			t.AppendRows([]table.Row{
				{index, md5_value, duration, size, url},
			})

			// fmt.Printf("请求时间：%v 文件大小：%v 文件MD5：%x\n", duration, ByteCountDecimal(utf8.RuneCountInString(string(body))), md5.Sum(body))
		}
		fmt.Println(t.Render())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
