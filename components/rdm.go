package components

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"

	"goqt-redis/libs/helper"
	"goqt-redis/libs/rdm"

	"github.com/gorilla/websocket"
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
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  10240,
		WriteBufferSize: 10240,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
	}
	mux.HandleFunc("/redis/connection/pubsub", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			data := make(map[string]interface{})
			_ = request.ParseForm()
			params := request.PostForm
			for param, values := range params {
				if len(values) > 0 {
					data[param] = values[0]
				} else {
					data[param] = nil
				}
			}
			_, _ = writer.Write([]byte(rdm.RedisPubSub(data)))
			return
		}

		ws, _ := upgrader.Upgrade(writer, request, nil)
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)
			}
		}()
		for {
			_, msg, err := ws.ReadMessage()
			if err == nil {
				data := make(map[string]interface{})
				if err := json.Unmarshal(msg, &data); err != nil {
					logger.Error(err)
					continue
				}
				data["ws"] = ws
				rdm.RedisPubSub(data)
			}
		}
	})
	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), handler)
	if err != nil {
		panic(err)
	}
}
