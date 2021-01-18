package rdm

import (
	"fmt"
	"goqt-redis/libs/helper"
	"sort"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func RedisManagerConnectionTestForQt(data map[string]interface{}) (bool, error) {
	var config connection
	config.Ip = data["ip"].(string)
	config.Title = data["title"].(string)
	config.Port = data["port"].(string)
	config.Auth = data["auth"].(string)

	client, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", config.Ip, config.Port))
	if err != nil {
		return false, err
	}
	defer client.Close()
	if config.Auth != "" {
		_, err := client.Do("AUTH", config.Auth)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func RedisManagerConfigSaveForQt(data map[string]interface{}) int {
	var config connection
	config.Ip = data["ip"].(string)
	config.Title = data["title"].(string)
	config.Port = data["port"].(string)
	config.Auth = data["auth"].(string)
	totalConnection = totalConnection + 1
	config.ID = int64(totalConnection)

	// 判断存在
	for _, conn := range connectionList {
		if conn.Ip == config.Ip && conn.Port == config.Port {
			helper.ShowWarningMessage("错误提醒", "已经存在相同的连接, 名称为: "+config.Title)
			return -1
		}
	}
	connectionList = append(connectionList, config)
	err := writeConfigJSON()
	if err != nil {
		helper.ShowWarningMessage("错误提醒", "保存连接失败")
		return -1
	}
	return int(config.ID)
}

func RedisManagerRemoveConnectionForQt(data map[string]interface{}) bool {
	var configs []connection
	id := int64(getFromInterfaceOrFloat64ToInt(data["id"]))
	if id == 0 {
		return false
	}
	for _, v := range connectionList {
		if v.ID != id {
			configs = append(configs, v)
		}
	}
	connectionList = configs
	err := writeConfigJSON()
	if err != nil {
		helper.ShowWarningMessage("错误", "删除连接失败")
		return false
	}
	return true
}

func RedisManagerConnectionListForQt() []struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Flag  string `json:"flag"`
} {
	err := readConfigJSON()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var conns []struct {
		ID    int64  `json:"id"`
		Title string `json:"title"`
		Flag  string `json:"flag"`
	}
	for _, conn := range connectionList {
		conns = append(conns, struct {
			ID    int64  `json:"id"`
			Title string `json:"title"`
			Flag  string `json:"flag"`
		}{ID: conn.ID, Title: conn.Title, Flag: fmt.Sprintf("%d:%s:%s", conn.ID, conn.Ip, conn.Port)})
	}

	return conns
}

func RedisManagerDbListForQt(id interface{}, isQt bool) []int {
	redisClient, _ := getRedisClient(map[string]interface{}{"id": id, "isQt": isQt}, false, false)
	defer redisClient.Close()
	var dbs []int
	for i := 0; i < 20; i++ {
		_, err := redisClient.Do("SELECT", i)
		if err != nil {
			break
		}
		//读取数据总量
		total, _ := redis.Int(redisClient.Do("DBSIZE"))
		dbs = append(dbs, total)
	}
	return dbs
}

func RedisManagerConnectionSelectDbForQt(data RequestData) []string {
	client, _ := getRedisClient(data, false, false)
	defer client.Close()
	index := getFromInterfaceOrFloat64ToInt(data["index"])
	_, _ = client.Do("SELECT", index) //选择数据库
	var nextCur = "0"
	pattern, ok := data["pattern"]
	if !ok {
		pattern = "*"
	}
	var resKeys = []string{}
	for {
		repl, _ := client.Do("SCAN", nextCur, "MATCH", pattern.(string))
		nextCur = string(repl.([]interface{})[0].([]byte))
		keys, err := redis.Strings(repl.([]interface{})[1], nil)
		if err != nil {
			return nil
		}
		for _, v := range keys {
			resKeys = append(resKeys, v)
		}
		if nextCur == "0" || nextCur == "" {
			break
		}
	}
	sort.Strings(resKeys) // 排序key
	return resKeys
}

func RedisManagerBatchDeleteForQt(data RequestData) ResponseData {
	keys := RedisManagerConnectionSelectDbForQt(data) // 获取要删除的Key
	client, _ := getRedisClient(data, false, false)
	defer client.Close()
	if len(keys) > 0 {
		for _, key := range keys {
			client.Do("DEL", key)
		}
	}
	return ResponseData{SuccessCode, "", nil}
}

func RedisManagerConnectionServerForQt(data RequestData) string {
	client, _ := getRedisClient(data, false, false)
	defer client.Close()
	var err error
	action := strings.Trim(data["action"].(string), " ")
	switch action {
	case "get_value":
		index := getFromInterfaceOrFloat64ToInt(data["index"])
		_, err = client.Do("SELECT", index) //选择数据库
		if err != nil {
			return JSON(ResponseData{FailedCode, FailedMsg, nil})
		}
		key := data["key"].(string)
		if key == "" {
			return JSON(ResponseData{FailedCode, FailedMsg, nil})
		}
		typeStr, _ := redis.String(client.Do("TYPE", key))
		if typeStr == "none" {
			return JSON(ResponseData{5001, FailedMsg, nil})
		}
		ttl, _ := redis.Int64(client.Do("TTL", key))
		switch typeStr {
		case "list":
			val, err := redis.Strings(client.Do("LRANGE", key, 0, 1000))
			if err != nil {
				return JSON(ResponseData{FailedCode, "读取数据错误", err.Error()})
			} else {
				return JSON(ResponseData{SuccessCode, "读取所有key成功", RequestData{
					"type": typeStr,
					"data": val,
					"ttl":  ttl,
				}})
			}
		case "set":
			val, err := redis.Strings(client.Do("SMEMBERS", key))
			if err != nil {
				return JSON(ResponseData{FailedCode, "读取数据错误", err.Error()})
			} else {
				return JSON(ResponseData{SuccessCode, "读取所有key成功", RequestData{
					"type": typeStr,
					"data": val,
					"ttl":  ttl,
				}})
			}
		case "zset":
			val, err := redis.StringMap(client.Do("ZRANGEBYSCORE", key, "-inf", "+inf", "WITHSCORES"))
			if err != nil {
				return JSON(ResponseData{FailedCode, "读取数据错误", err.Error()})
			} else {
				var retData []map[string]string
				for k, v := range val {
					retData = append(retData, map[string]string{"value": k, "score": v})
				}
				return JSON(ResponseData{SuccessCode, "读取所有key成功", RequestData{
					"type": typeStr,
					"data": retData,
					"ttl":  ttl,
				}})
			}
		case "string":
			val, err := redis.String(client.Do("GET", key))
			if err != nil {
				return JSON(ResponseData{FailedCode, "读取数据错误", err.Error()})
			} else {
				return JSON(ResponseData{SuccessCode, "读取所有key成功", RequestData{
					"type": typeStr,
					"data": val,
					"ttl":  ttl,
				}})
			}
		case "hash":
			val, err := redis.StringMap(client.Do("HGETALL", key))
			if err != nil {
				return JSON(ResponseData{FailedCode, "读取数据错误", err.Error()})
			} else {
				return JSON(ResponseData{SuccessCode, "读取所有key成功", RequestData{
					"type": typeStr,
					"data": val,
					"ttl":  ttl,
				}})
			}
		}
	case "select_db":
		index := getFromInterfaceOrFloat64ToInt(data["index"])
		_, _ = client.Do("SELECT", index) //选择数据库
		var nextCur = "0"
		var resKeys = map[string][]string{}
		for {
			repl, _ := client.Do("SCAN", nextCur)
			nextCur = string(repl.([]interface{})[0].([]byte))
			keys, err := redis.Strings(repl.([]interface{})[1], nil)
			if err != nil {
				return JSON(ResponseData{FailedCode, "错误,无法解析SCAN返回值", nil})
			}
			for _, v := range keys {
				resKeys[v] = append(resKeys[v], v)
			}
			if nextCur == "0" || nextCur == "" {
				break
			}
		}
		return JSON(ResponseData{SuccessCode, "读取所有key成功", resKeys})
	}
	return JSON(ResponseData{FailedCode, "错误,无法解析到动作:" + action, nil})
}
