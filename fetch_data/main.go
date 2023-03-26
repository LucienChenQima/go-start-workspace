package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    // 要刷新的网页URL
    url := "https://lucienchen.xyz/news/articles/1635822673517637633"
    // 循环刷新10000次
    for i := 0; i < 10000; i++ {
        fmt.Println("刷新页面：", url)
        // 发送GET请求
        _, err := http.Get(url)
        if err != nil {
            fmt.Println("请求失败:", err)
            return
        }
        // 等待1秒钟
        time.Sleep(time.Second)
    }
}