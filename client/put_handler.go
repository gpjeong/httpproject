package client

import (
	"fmt"
	"httpproject/internal/config"
	"httpproject/rest"
	"httpproject/util/logger"
	"net/http"
	"time"
)

func PutApiData(id, name, balance string) {
	time.Sleep(2 * time.Second)
	url := "http://" + config.ServerConfig.ApiInfo.ApiHost + ":" + config.ServerConfig.ApiInfo.ApiPort + "/put?id=" + id

	reqData := make(map[string]interface{})
	reqData["name"] = name
	reqData["balance"] = balance

	resp, statusCode, err := rest.RequetApiMethod(url, "PUT", reqData)

	if statusCode != http.StatusOK {
		var statusMessage string
		if err != nil {
			statusMessage = fmt.Sprintf("StatusCode: %d, ErrMsg: %v", statusCode, err.Error())
		} else {
			// 400 관련 에러는 에러 메시지가 null 값으로 return 되어 statusCode만 로그에 출력
			statusMessage = fmt.Sprintf("StatusCode: %d", statusCode)
		}
		logger.Log.Error().Msgf("Put OJT Api failed :", statusMessage)
	}

	if resp != nil {
		fmt.Println(string(resp))
	}
}
