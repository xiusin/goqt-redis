package components

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/xiusin/logger"
	"goqt-redis/libs/rdm"
	"net/http"
)

// 发布订阅进程
func pubSub() {
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
}
