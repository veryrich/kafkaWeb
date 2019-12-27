package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化web
	router := gin.Default()
	router.Static("/static/css", "./static/css")
	router.Static("/static/js", "./static/js")

	// 运行Controller
	Index(router)

	// 执行 计划任务
	//wg.Add(1)
	//go kafkaCheck(&wg)
	//wg.Wait()

	router.Run(":9090")
}
