package components

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"

	"goqt-redis/libs/helper"
	"goqt-redis/libs/rdm"

	"github.com/rs/cors"
	"github.com/xiusin/logger"
)

var mux *http.ServeMux

var serverPort int

func init() {
	f, err := os.OpenFile(helper.UserHomeDir("error.log"), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err == nil {
		logger.SetOutput(io.MultiWriter(os.Stdout, os.Stderr, f))
	}
	serverPort = rand.Intn(10000) + 6000
}

// InitRdm 初始化 Rdm的 web服务
func InitRdm() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("启动rdm服务失败", err, string(debug.Stack()))
		}
	}()
	var routes = map[string]rdm.HandleFunc{
		"/redis/connection/test":        rdm.RedisManagerConnectionTest,
		"/redis/connection/save":        rdm.RedisManagerConfigSave,
		"/redis/connection/list":        rdm.RedisManagerConnectionList,
		"/redis/connection/server":      rdm.RedisManagerConnectionServer,
		"/redis/connection/removekey":   rdm.RedisManagerRemoveKey,
		"/redis/connection/removerow":   rdm.RedisManagerRemoveRow,
		"/redis/connection/updatekey":   rdm.RedisManagerUpdateKey,
		"/redis/connection/addkey":      rdm.RedisManagerAddKey,
		"/redis/connection/flushDB":     rdm.RedisManagerFlushDB,
		"/redis/connection/remove":      rdm.RedisManagerRemoveConnection,
		"/redis/connection/command":     rdm.RedisManagerCommand,
		"/redis/connection/info":        rdm.RedisManagerGetInfo,
		"/redis/connection/get-command": rdm.RedisManagerGetCommandList,
	}
	mux = http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("dist")))
	for route, handle := range routes {
		mux.HandleFunc(route, func(handle rdm.HandleFunc) func(writer http.ResponseWriter, request *http.Request) {
			return func(writer http.ResponseWriter, request *http.Request) {
				defer func() {
					if err := recover(); err != nil {
						s := debug.Stack()
						logger.Errorf("Recovered Error: %s, ErrorStack: \n%s\n\n", err, string(s))
					}
				}()
				var params url.Values
				data := make(map[string]interface{})
				if request.Method == http.MethodPost {
					_ = request.ParseForm()
					params = request.PostForm
				} else {
					params = request.URL.Query()
				}
				for param, values := range params {
					if len(values) > 0 {
						data[param] = values[0]
					} else {
						data[param] = nil
					}
				}
				_, _ = writer.Write([]byte(handle(data)))
			}
		}(handle))
	}
	pubSub()
	handler := cors.Default().Handler(mux)

	for {
		_ = http.ListenAndServe(fmt.Sprintf(":%d", serverPort), handler) // 端口冲突则更换接口
		fmt.Println("端口被占用:", serverPort)
	}

}


