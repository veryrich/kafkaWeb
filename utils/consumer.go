package utils

import (
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

func KafkaConsumer(addr string, topics string, times int, group *sync.WaitGroup, ch chan interface{}) {

	defer group.Done()
	address := StrToIp(addr)

	//创建一个配置对象
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V2_3_0_0

	//创建一个客户消费者
	client, err := sarama.NewClient(address, consumerConfig)
	if err != nil {
		log.Printf("unable to create kafka client: %q", err)
		ch <- err
		return
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Println(err)
		ch <- err
		return
	}
	defer consumer.Close()

	consumerLoop(consumer, topics, times, ch)
}

func consumerLoop(consumer sarama.Consumer, topic string, times int, ch chan interface{}) {
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Println("unable to fetch partition IDs for the topic", topic, err)
		ch <- err
		return
	}

	var wg sync.WaitGroup
	for partition := range partitions {
		wg.Add(1)
		go func() {
			consumePartition(consumer, int32(partition), times, topic, ch, &wg)
		}()
	}
	wg.Wait()
}

// 开始消费消息
func consumePartition(consumer sarama.Consumer, partition int32, times int, topic string, ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Receving on partition", partition)
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Println(err)
			ch <- err
		}
	}()

	consumed := 0

	for i := 0; i < times; i++ {
		msg := <-partitionConsumer.Messages()
		t := msg.Timestamp
		formatTime := t.Format("15:04:05")
		log.Printf("%s - %s\n", msg.Value, formatTime)
		consumed++
		ch <- string(msg.Value)
	}

	//log.Printf("Consumed: %d\n", consumed)

}
