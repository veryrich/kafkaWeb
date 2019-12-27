package main

import (
	"awesomeProject1/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type KafkaInfo struct {
	KafkaIP string `form:"kafkaIP"`
	Message MessageInfo
}

type MessageInfo struct {
	Content string `form:"messageContent"`
	Number  int    `form:"messageNumber"`
	Topic   string `form:"kafkaTopic"`
}

var ch = make(chan interface{}, 128)

func Index(router *gin.Engine) {
	var kafkaInfo KafkaInfo
	var wg = &sync.WaitGroup{}

	router.LoadHTMLGlob("templates/*")

	// 首页
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 开始生产/消费
	router.POST("/", func(c *gin.Context) {

		c.Bind(&kafkaInfo)

		if validator(kafkaInfo) {
			wg.Add(1)
			go utils.SyncProducer(kafkaInfo.KafkaIP, kafkaInfo.Message.Number, kafkaInfo.Message.Content, kafkaInfo.Message.Topic, wg, ch)

			wg.Add(1)
			go utils.KafkaConsumer(kafkaInfo.KafkaIP, kafkaInfo.Message.Topic, kafkaInfo.Message.Number, wg, ch)
		}

		ch <- "数据不合法"
		return

	})

	wg.Wait()

	router.GET("/getStreamData", func(c *gin.Context) {

		select {
		case data := <-ch:
			c.String(http.StatusOK, "%s", data)
		default:
			c.String(http.StatusOK, "Empty")

		}
	})

}

// 数据校验
func validator(k KafkaInfo) bool {
	if k.KafkaIP != "" && k.Message.Content != "" && k.Message.Topic != "" && 0 > k.Message.Number && k.Message.Number <= 60 && len(k.KafkaIP) >= 12 {
		return true
	}

	return false
}
