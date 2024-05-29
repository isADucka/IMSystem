/*
 * @Author: cyy
 * @Description: --
 */
package discovery

import "encoding/json"

type EndpointInfo struct {
	IP       string                 `json:"ip"`
	Port     string                 `json:"port"`
	MetaData map[string]interface{} `json:"meta"`
}

/**
* 将json的数据类型转化为go的结构类型
 */
func UnMarshal(data []byte) (*EndpointInfo, error) {
	ed := &EndpointInfo{}
	err := json.Unmarshal(data, ed)
	if err != nil {
		return nil, err
	}
	return ed, nil
}

/**
* 将go的结构类型转化为json
 */
func (ed *EndpointInfo) Marshal() string {
	data, err := json.Marshal(ed)
	if err != nil {
		panic(err)
	}
	return string(data)
}
