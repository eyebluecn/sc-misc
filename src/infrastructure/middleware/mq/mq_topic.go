package mq

// 消息发送的目标Topic名称，需要提前在RocketMQ中创建。
// 一般一个领域对象对应一个TOPIC.
type MqTopic string

const (
	MqTopicPayment  MqTopic = "SMART_CLASSROOM_PAYMENT"  //支付单 主题
	MqTopicAuthor   MqTopic = "SMART_CLASSROOM_AUTHOR"   //作者 主题
	MqTopicEditor   MqTopic = "SMART_CLASSROOM_EDITOR"   //小编 主题
	MqTopicReader   MqTopic = "SMART_CLASSROOM_READER"   //读者 主题
	MqTopicColumn   MqTopic = "SMART_CLASSROOM_COLUMN"   //专栏 主题
	MqTopicContract MqTopic = "SMART_CLASSROOM_CONTRACT" //合同 主题
)
