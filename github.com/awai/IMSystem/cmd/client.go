package cmd

import (
	"github.com/awai/IMSystem/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "clinet命令，启动方式为./项目 client",
	Run:   ClientHandle,
}

func ClientHandle(cmd *cobra.Command, arg []string) {
	client.RunMain()
}
