/*
 * @Author: cyy
 * @Description: 用于读取配置文件
 */
package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func Init(path string) {
	fmt.Print("我是配置文件初始化"+path,"         ",CONFIGFILE_TYPE)
	viper.SetConfigFile(path)
	viper.SetConfigType(CONFIGFILE_TYPE)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

const (
	DISCOVERY_ENDPOINTS string = "discovery.endpoints"
	DISCOVERY_TIMEOUT   string = "discovery.timeout"
	IPCONF_SERVICEPATH  string = "ip_conf.service_path"
	GLOBAL_ENV          string = "global.env"
	DEBUG               string = "debug"
	CONFIGFILE_TYPE            = "yaml"
)

/*
*
* 获取服务发现的地址
(其实就是从配置文件里面获取key的值，然后获取对应的value值)
*/
func GetEndpointsDiscovery() []string {
	return viper.GetStringSlice(DISCOVERY_ENDPOINTS)
}

/**
* 获取连接服务发现集群超时事件，加上单位为s
 */
func GetTimeoutDiscover() time.Duration {
	return viper.GetDuration(DISCOVERY_TIMEOUT) * time.Second
}

/**
* 获取地址
 */
func GetServicePath() string {
	fmt.Print(viper.GetString(IPCONF_SERVICEPATH))
	return viper.GetString(IPCONF_SERVICEPATH)
}

/**
* 判断是否是debug环境
 */
func IsDubug() bool {
	env := viper.GetString(GLOBAL_ENV)
	return env == DEBUG
}
