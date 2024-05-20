package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
)

func init() {
	cobra.OnInitialize(initConfig)
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
