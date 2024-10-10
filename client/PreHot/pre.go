package PreHot

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// 用来预热及测试
func Test() {
	time.Sleep(10 * time.Second)
	for i := 0; i < 10000; i++ {
		go func() {
			res, err := http.Get("http://localhost:8080/Ping")
			if err != nil {
				// 日志记录或处理错误
				fmt.Println(err)
				return
			}
			// 确保响应被关闭
			defer func(Body io.ReadCloser) {
				err = Body.Close()
				if err != nil {
					fmt.Println("close err")
					return
				}
			}(res.Body)
		}()
	}
}
