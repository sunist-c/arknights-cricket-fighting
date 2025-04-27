package main

import (
	"arknights-cricket-fighting/app/router"
	"os"
)

func main() {
	// 获取环境变量中的端口，如果未设置则使用默认端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "18080"
	}
	
	// 启动服务器，监听所有接口
	_ = router.GetEngine().Run(":" + port)
}
