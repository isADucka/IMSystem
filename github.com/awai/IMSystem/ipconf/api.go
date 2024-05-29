/*
 * @Author: cyy
 * @Description: --
 */
package ipconf

import (
	"context"

	"github.com/awai/IMSystem/ipconf/domain"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

//封装返回的数据
type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}


// 对外的响应接口 
func GetIpInfoList(c context.Context, ctx *app.RequestContext) {
	defer func() {
		//recover()用于捕获panic
		if err := recover(); err != nil {
			ctx.JSON(consts.StatusBadRequest, utils.H{"err": err})
		}
	}()
	//① 构建客户请求信息
	ipConfContext := domain.NewIPConfContext(&c, ctx)
	//② 进行ip调度
	eds:= domain.Dispatch(ipConfContext)
	ipConfContext.AppCtx.JSON(consts.StatusOK,Result(Top5EndPoints(eds)))

}
