/*
 * @Author: cyy
 * @Descripttion: 封装消息部分
 */
package sdk

const(
	MsgType="text"
)

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

func (chat *Chat) SendMessage(msg *Message) {
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
