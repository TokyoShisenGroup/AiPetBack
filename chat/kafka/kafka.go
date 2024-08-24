// Description: kafka操作封装
package kafka

import (
	"github.com/IBM/sarama"
	"strings"
)


var producer sarama.AsyncProducer
var topic string = "default_message"

func InitProducer(topicInput, hosts string) {
	topic = topicInput
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if nil != err {
		//log.Logger.Error("init kafka client error", log.Any("init kafka client error", err.Error()))
	}

	producer, err = sarama.NewAsyncProducerFromClient(client)
	if nil != err {
		//log.Logger.Error("init kafka async client error", log.Any("init kafka async client error", err.Error()))
	}
}

func Send(data []byte) {
	be := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: be}
}

func Close() {
	if producer != nil {
		producer.Close()
	}
}


var consumer sarama.Consumer

type ConsumerCallback func(data []byte)

// 初始化消费者
func InitConsumer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if nil != err {
		//log.Logger.Error("init kafka consumer client error", log.Any("init kafka consumer client error", err.Error()))
	}

	consumer, err = sarama.NewConsumerFromClient(client)
	if nil != err {
		//log.Logger.Error("init kafka consumer error", log.Any("init kafka consumer error", err.Error()))
	}
}

// 消费消息，通过回调函数进行
func ConsumerMsg(callBack ConsumerCallback) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if nil != err {
		//log.Logger.Error("iConsumePartition error", log.Any("ConsumePartition error", err.Error()))
		return
	}

	defer partitionConsumer.Close()
	for {
		msg := <-partitionConsumer.Messages()
		if nil != callBack {
			callBack(msg.Value)
		}
	}
}

func CloseConsumer() {
	if nil != consumer {
		consumer.Close()
	}
}
