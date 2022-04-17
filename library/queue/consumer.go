package queue

type consumer interface {
	// @param queueName 监听的队列名
	// @param group 事件处理组，同 group 对于同一个事件只处理一次
	// @param handler 事件处理函数
	// consumer
	Listen()
}
