package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var zhipuapikey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" //===========input your key for 智普ai

func main() {
	allmes := []string{"你好1", "今天天气如何", "你会不会死", "你是谁啊", "我是谁", "你好1", "今天天气如何", "你会不会死", "你是谁啊", "我是谁"}
	var wg sync.WaitGroup
	wg.Add(len(allmes)) // 添加两个子协程
	t := time.Now()

	aa := func(a11 string) string { //go不支持函数嵌套, 但是可以写匿名函数,然后在外面给他命名aa即可.go语法很弱智.
		defer wg.Done()
		type yitiao struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}
		type yici struct {
			Model    string   `json:"model"`
			Messages []yitiao `json:"messages"`
		}
		a := yici{
			Model: "glm-4-0520",
			Messages: []yitiao{
				{Role: "user", Content: a11},
			},
		}
		a1, _ := json.Marshal(a)

		print(a1)

		targetUrl := "https://open.bigmodel.cn/api/paas/v4/chat/completions"

		client := &http.Client{}

		req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(a1))

		req.Header.Add("Authorization", "Bearer "+zhipuapikey)
		req.Header.Add("Content-Type", "application/json")

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))

		return string(body)
	}
	for _, x := range allmes {
		go aa(x)
	}

	wg.Wait() // 等待所有子协程完成
	t2 := time.Now()
	fmt.Println("实用了多长时间", t2.Sub(t))
}
