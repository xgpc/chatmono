package openAI

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/util"
	"time"
)

func SetSession(key string, req *ReqOpenAI) {
	_, err := redis().Set(context.Background(), key, req.String(), time.Hour*24*30).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
}

func GetSession(key string) (*ReqOpenAI, error) {
	result, err := redis().Get(context.Background(), key).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	var Req ReqOpenAI

	err = util.JsonDecode([]byte(result), &Req)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	return &Req, nil
}

func Exists(key string) bool {
	result, err := redis().Exists(context.Background(), key).Result()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	return result > 0
}
