/*
 * @Author: cyy
 * @Description: --
 */
/*
 * @Author: cyy
 * @Description:和配置相关
 */
package cmd

import (
	"github.com/awai/IMSystem/ipconf"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipConfCmd)
}

var ipConfCmd = &cobra.Command{
	Use:   "ipconf",
	Short: "启动的命令行： ./项目 ipconf --config=./plato.yaml",
	Run:   IpConfHandler,
}

/**
*
 */
func IpConfHandler(cmd *cobra.Command, arg []string) {
	ipconf.RunMain(ConfigPath)
}
