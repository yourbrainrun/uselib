package main

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/vcr"
)

func main() {
	testVcr()
}

func testVcr() {
	client,_ := vcr.NewClient("19e4425489f5a1", "ebd4c801", "vcr.bj.baidubce.com")
	// 配置不进行重试，默认为Back Off重试
	client.Config.Retry = bce.NewNoRetryPolicy()

	// 配置连接超时时间为30秒
	client.Config.ConnectionTimeoutInMillis = 30 * 1000

	addr := "http://play.roadofcloud.net/live/0a8070ea0/playlist.m3u8"
	err := client.SimplePutMedia(addr, "test ", "default", "vcr_callback")
	if err != nil {
		return
	}

	fmt.Println(client)
}
