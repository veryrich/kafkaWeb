package utils

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"time"
)

// 同步模式生产者
func SyncProducer(addr string, times int, messages string, topic string, wg *sync.WaitGroup, ch chan interface{}) {
	defer wg.Done()

	address := StrToIp(addr)
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 3 * time.Second
	config.Version = sarama.V2_3_0_0
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Println( err)
		ch<-err
		return
	}
	defer p.Close()

	srcValue := messages
	for i := 0; i < times; i++ {
		value := fmt.Sprintf(srcValue, i)
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Key:       nil,
			Value:     sarama.StringEncoder(value),
			Headers:   nil,
			Metadata:  nil,
			Offset:    0,
			Partition: 0,
			Timestamp: time.Now(),
		}
		//part, offset, err := p.SendMessage(msg)
		_, _, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
			ch <- err
		} else {
			//fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
			//fmt.Println("生产者",value)
		}

	}

	return

}
