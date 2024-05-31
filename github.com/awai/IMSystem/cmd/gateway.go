/*
 * @Author: cyy
 * @Description: 网关部分
 */
package cmd

import (
	"github.com/awai/IMSystem/gateway"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
	Use: "gateway",
	Short:"网关部分",
	Run:GatewayHandler,
}

func GatewayHandler(cmd *cobra.Command,args []string){
	gateway.RunMain(ConfigPath)
}
