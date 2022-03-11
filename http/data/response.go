package data

import (
	"context"
	"fmt"
	smalllogger "github.com/PepperTalk/api/small-logger"
	logger2 "gorm.io/gorm/logger"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

type Response struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func ResponseFrom(message string, data map[string]interface{}) Response {
	if data == nil {
		data = map[string]interface{}{}
	}
	return Response{
		Message: message,
		Data:    data,
	}
}

func MapFrom2(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{key: value}
}

func MapFrom(halfPair ...interface{}) map[string]interface{} {
	if len(halfPair) == 0 {
		return map[string]interface{}{}
	}

	var logger logger2.Interface
	if goyave.IsReady() {
		logger = database.GetConnection().Config.Logger
	} else {
		logger = &smalllogger.SmallLogger
	}
	if len(halfPair)%2 != 0 {
		logger.Error(context.Background(), "Invalid number of arguments !")
		return nil
	}

	final := map[string]interface{}{}
	isKey := true
	for i, v := range halfPair {
		if isKey {
			_, ok := v.(string)
			if !ok {
				logger.Error(context.Background(), fmt.Sprintf("Argument '%v' (n°%d) is not of type string !", v, i))
				return nil
			}
		} else {
			key := halfPair[i-1].(string)
			final[key] = v
		}
		isKey = !isKey
	}
	return final
}
