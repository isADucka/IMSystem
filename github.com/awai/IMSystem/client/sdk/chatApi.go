/*
 * @Author: cyy 2867025942@qq.com
 * @Date: 2024-05-21 11:39:26
 * @LastEditors: cyy 2867025942@qq.com
 * @LastEditTime: 2024-05-21 21:23:06
 * @FilePath: /IMSystem/client/sdk/chatApi.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package sdk

/*
* 对消息的封装
*
 */
type Message struct {
	Type       string
	Name       string
	FromUserId string
	ToUserId   string
	Content    string
	Session    string
}

/**
* 对聊天对象的封装
 */
type Chat struct {
	Nick      string
	UserId    string
	SeesionId string
	conn      *connect
}

/**
* 新建一个聊天
 */
func NewChat(serverAddr, nick, userId, sessionId string) *Chat {
	return &Chat{
		Nick:      nick,
		UserId:    userId,
		SeesionId: sessionId,
		conn:      newConnet(serverAddr),
	}
}

func (chat *Chat) sendMessage(msg *Message) {
	//通过chat的连接调用发送消息
	chat.conn.send(msg)
}

/**
* 获取connet对象的reveChan对象
 */
func (chat *Chat) Recv() <-chan *Message {
	return chat.conn.recv()
}

/**
* 关闭聊天
 */
func (chat *Chat) Close() {

}
