package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
	"os"
)

func main() {

	Consumer()

	//groupName := "test-group"
	//mqAddr := "127.0.0.1:9876"
	//topic := "test-topic"
	//
	//var err error
	//MqProducer, err := rocketmq.NewProducer(
	//	producer.WithGroupName(groupName),
	//	producer.WithNameServer([]string{mqAddr}),
	//	producer.WithRetry(3),
	//	producer.WithQueueSelector(producer.NewHashQueueSelector()),
	//)
	//if err != nil {
	//	panic(fmt.Sprintf("init rocket mq producer err:%v", err))
	//	return
	//}
	//
	//err = MqProducer.Start()
	//if err != nil {
	//	panic(fmt.Sprintf("producer mq start err:%v", err))
	//	return
	//}
	////ch := make(chan struct{}, 3)
	//defer MqProducer.Shutdown()
	//
	//msg := &primitive.Message{
	//	Topic: topic,
	//	Body:  []byte("this is a message body"),
	//}
	//
	//msg.WithShardingKey("testShardingKey")
	////发送带tag的消息
	//msg.WithTag("testTag")
	//
	//for {
	//	res, err := MqProducer.SendSync(context.Background(), msg)
	//	fmt.Println(res, err)
	//}
}

func Consumer() {
	testChan := make(chan string, 0)
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("preview_sync_data"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"10.255.67.7:9876"})),
	)
	topic := "pre_data_calculate"

	err := c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			log.Printf("订阅到消息: body=%v, tag =%v \n", string(msgs[i].Body), msgs[i].GetTags())
			testChan <- string(msgs[i].Body)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		log.Printf(err.Error())
	}
	err = c.Start()
	defer c.Shutdown()
	if err != nil {
		log.Printf(err.Error())
		os.Exit(-1)
	}
	//time.Sleep(20*time.Minute)

	for {
		select {
		case message := <-testChan:
			fmt.Println(message)
		}
	}
}
