/*
 * @Author: cyy
 * @Description: --
 */
package ipconf

import (
	"github.com/awai/IMSystem/ipconf/source"
	"github.com/awai/IMSystem/ipconf/utils"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/**
* 启动web容器
 */
func RunMain(path string) {
	utils.Init(path) //这里先要读取配置文件
	source.Init()    // 这里是发现服务并且注册服务
	//调度层启动
	s := server.Default(server.WithHostPorts(":6789")) //记得规范，是"："+端口号
	s.GET("/ip/list", GetIpInfoList)
	s.Spin()
}
