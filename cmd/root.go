/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/frank-yf/go-tool/internal/datecmd"
	"github.com/frank-yf/go-tool/internal/jsoncmd"
	"github.com/frank-yf/go-tool/internal/pkg/zlog"

	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "go-tool",
		Short: "基于Go语言开发的常用命令工具包",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 暂时无法生效
	debug := rootCmd.Flags().Bool("debug", false, "使用debug输出级别")
	if debug != nil && *debug {
		zlog.UseDebug()
	}
	zlog.Initial()

	datecmd.Init(rootCmd)
	jsoncmd.Init(rootCmd)
}
