package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func get(url string) (duration time.Duration, body []byte, err error) {
	var (
		resp  *http.Response
		start = time.Now() // 请求开始时间
	)

	// 根据URL获取资源
	if resp, err = http.Get(url); err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return duration, nil, err
	}

	// 关闭资源流
	defer resp.Body.Close()

	// 读取资源数据 body: []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return duration, nil, err
	}

	duration = time.Since(start) // 请求总计耗时

	return duration, body, nil
}
