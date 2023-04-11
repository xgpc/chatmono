package openAI

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/util"
	"strconv"
	"time"
)

// 通过hset来保存用户的会话数据

type SessionData struct {
	Request     ReqOpenAI `json:"request"`
	Title       string    `json:"title"`
	TemplateID  uint32    `json:"template_id"`
	SessionType int       `json:"session_type"` // 回话类型 1单次, 2连续
}

func (p *SessionData) String() string {
	marshal, err := util.JsonEncode(p)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	return string(marshal)
}

func SessionDecode(data string) *SessionData {
	var info SessionData
	util.JsonDecode([]byte(data), &info)
	return &info
}

func SetSession(userId uint32, sessionKey string, req *SessionData) {
	idStr := strconv.Itoa(int(userId))
	_, err := redis().HSet(context.Background(), idStr, sessionKey, req.String(), time.Hour*24*30).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
}

func GetSession(userId uint32, sessionKey string) (*SessionData, error) {
	idStr := strconv.Itoa(int(userId))
	result, err := redis().HGet(context.Background(), idStr, sessionKey).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	Req := SessionDecode(result)
	return Req, nil
}

func DelSession(userId uint32, sessionKey string) error {
	idStr := strconv.Itoa(int(userId))
	_, err := redis().HDel(context.Background(), idStr, sessionKey).Result()
	return err
}

func GetAll(userId uint32) map[string]SessionData {
	idStr := strconv.Itoa(int(userId))
	md := map[string]SessionData{}
	result, err := redis().HGetAll(context.Background(), string(idStr)).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	for k, v := range result {
		Req := SessionDecode(v)

		md[k] = *Req
	}
	return md
}
