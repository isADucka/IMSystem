/*
 * @Author: cyy
 * @Description: 初始化的cmd
 */
package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

var (
	ConfigPath string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&ConfigPath,
		"config",
		"./imSystem.yaml",
		"config file(default ./ImSytem.yaml)")
}

var rootCmd = &cobra.Command{
	Use:   "ImSystem",
	Short: "这是一个IM系统",
	Run:   ImSystem,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func ImSystem(cmd *cobra.Command, args []string) {

}

func initConfig() {

}
