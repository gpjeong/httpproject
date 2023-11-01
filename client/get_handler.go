package client

import (
	"fmt"
	"httpproject/internal/config"
	"httpproject/rest"
	"httpproject/util/logger"
	"net/http"
)

func GetApiData(inputData string) {
	url := config.ServerConfig.ApiInfo.ApiHost + ":" + config.ServerConfig.ApiInfo.ApiPort + "/get?input=" + inputData

	resp, statusCode, err := rest.RequetApiMethod(url, "GET")

	if statusCode != http.StatusOK {
		var statusMessage string
		if err != nil {
			statusMessage = fmt.Sprintf("StatusCode: %d, ErrMsg: %v", statusCode, err.Error())
		} else {
			// 400 관련 에러는 에러 메시지가 null 값으로 return 되어 statusCode만 로그에 출력
			statusMessage = fmt.Sprintf("StatusCode: %d", statusCode)
		}
		logger.Log.Error().Msgf("GET OJT Api failed :", statusMessage)
	}

	if resp != nil {
		fmt.Println(resp)
	}
}
