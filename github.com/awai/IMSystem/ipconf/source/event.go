/*
 * @Author: cyy
 * @Description: 这是一个事件
 */
package source

import (
	"fmt"

	"github.com/awai/IMSystem/ipconf/discovery"
)


var eventChan chan *Event
func EventChan()<-chan *Event{
	return eventChan
}

// 事件类型：有两个值addNode和delNode
type EnventType string

const (
	ADD_NODE_ENVENT EnventType = "addNode"
	DEL_NODE_ENVENT EnventType = "delNode"
	CONNETC_NUM string="connect_num"
	MESSAGE_BYTES string="message_bytes"
)

type Event struct {
	Type         EnventType
	IP           string
	Port         string
	ConnectNum   float64
	MessageBytes float64
}

/**
* 创建一个新的Event
*/
func NewEvent(ed *discovery.EndpointInfo) *Event{
	if ed==nil || ed.MetaData==nil{
		return nil
	}
	var connectNum, messageBytes float64
	if data,ok:=ed.MetaData[CONNETC_NUM];ok{
		connectNum=data.(float64)
	}
	if data,ok:=ed.MetaData[MESSAGE_BYTES];ok{
		messageBytes=data.(float64)
	}
	return &Event{
		Type: ADD_NODE_ENVENT,
		IP: ed.IP,
		Port: ed.Port,
		ConnectNum: connectNum,
		MessageBytes: messageBytes,
	}
}

func(event *Event)Key() string{
	return fmt.Sprintf("%s:%s",event.IP,event.Port)
}