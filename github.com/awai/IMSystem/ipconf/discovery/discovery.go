/*
 * @Author: cyy
 * @Description: --
 */
/*
 * @Author: cyy
 * @Description: --
 */
package discovery

import (
	"context"
	"fmt"
	"sync"

	"github.com/awai/IMSystem/ipconf/utils"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
)

// 服务发现
type ServiceDiscovery struct {
	client  *clientv3.Client
	lock    sync.Mutex
	context *context.Context
}

// 新建发现的服务
func NewServiceDiscovery(ctx *context.Context) *ServiceDiscovery {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   utils.GetEndpointsDiscovery(),
		DialTimeout: utils.GetTimeoutDiscover(),
	})
	if err != nil {
		logger.Fatal(err)
	}
	return &ServiceDiscovery{
		client:  client,
		context: ctx,
	}

}

// 初始化服务列表，监听发现的服务
func (s *ServiceDiscovery) WatcherService(prefix string, set, del func(key, value string)) error {
	//初始化服务：①根据前缀获取现有的key
	resp, err := s.client.Get(*s.context, prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	//②完成初始化
	for _, ev := range resp.Kvs {
		fmt.Print("初始化：", ev)
		set(string(ev.Key), string(ev.Value))
	}

	//③监视
	s.watcher(prefix, set, del)
	return nil

}

/**
* 新增或删除监听某项服务
 */
func (s *ServiceDiscovery) watcher(prefix string, set, del func(key, value string)) {
	rch := s.client.Watch(*s.context, prefix, clientv3.WithPrefix())
	for watch := range rch {
		for _, ev := range watch.Events {
			switch ev.Type {
			case mvccpb.PUT:
				set(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE:
				del(string(ev.Kv.Key), string(ev.Kv.Value))
			}
		}
	}

}

//关闭服务

func (s *ServiceDiscovery) Close() error {
	return s.client.Close()
}
