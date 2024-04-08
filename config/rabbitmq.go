package config

// rabbitMq
const (
	RabbitUserName = "root"
	RabbitUserPwd  = "lrbrabbit123qwer"
	RabbitUri      = "amqp://root:lrbrabbit123qwer@124.71.63.137:5672/lrhost"
)

// push mq  推送服务队列一些配置
const (
	RQ_PushQueue      = "LrPushQueue"
	RQ_PushExchange   = "LrPushExchange"
	RQ_PushBindingKey = "LrPushKey"
)

// 发送邮件
const (
	RQ_EmailQueue      = "LrEmailQueue"
	RQ_EmailExchange   = "LrEmailExchange"
	RQ_EmailBindingKey = "LrEmailKey"
)
