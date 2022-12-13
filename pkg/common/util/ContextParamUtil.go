package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLoginUserId(c *gin.Context) (userId int) {
	val, exit := c.Get("userId")
	if exit {
		userId, _ = strconv.Atoi(fmt.Sprint(val))
	}
	return
}

func ObjToJSON(v interface{}) (jsonStr string) {
	jsonByte, _ := json.Marshal(v)
	jsonStr = string(jsonByte)
	return
}

func JSONToObject(v interface{}) (r interface{}) {
	jsonByte := fmt.Sprintf("%v", v.(interface{}))
	json.Unmarshal([]byte(jsonByte), &r)
	return
}