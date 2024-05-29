/*
 * @Author: cyy
 * @Description: --
 */
/*
 * @Author: cyy
 * @Description: 用于启动数据源
 */
package source

import (
	"context"

	"github.com/awai/IMSystem/ipconf/discovery"
	"github.com/awai/IMSystem/ipconf/utils"
	"github.com/bytedance/gopkg/util/logger"
)

func Init() {
	//创建一个Event
	eventChan = make(chan *Event)
	ctx := context.Background() //创建一个上下文context
	go DataHandler(&ctx)

}

/**
*服务发现
 */
func DataHandler(ctx *context.Context) {
	disc := discovery.NewServiceDiscovery(ctx) //创建服务

	//退出时关闭服务
	defer disc.Close()

	//匿名函数set和del
	setFunc := func(key, value string) {
		if ed, err := discovery.UnMarshal([]byte(value)); err == nil {
			if event := NewEvent(ed); event != nil {
				event.Type = ADD_NODE_ENVENT
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.setFunc.err :%s", err.Error())
		}
	}

	delFunc := func(key, value string) {
		if ed, err := discovery.UnMarshal([]byte(value)); err == nil {
			if event := NewEvent(ed); ed != nil {
				event.Type = DEL_NODE_ENVENT
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.delFunc.err :%s", err.Error())
		}
	}
	//监听服务
	err := disc.WatcherService(utils.GetServicePath(), setFunc, delFunc)
	if err != nil {
		panic(err)
	}
}
