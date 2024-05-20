package sdk

type connect struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}

func newConnet(serverAddr string) *connect {
	return &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *Message),
		recvChan:   make(chan *Message),
	}
}

/**
* 获取一个reveChan的对象
*/
func (c *connect) recv() <-chan *Message {
	return c.recvChan
}

/**
* 将数据直接发送给接收方
 */
func (c *connect) send(data *Message) {
	c.recvChan <- data
}

/**
* 关闭连接
 */
func (c *connect) close() {
}
