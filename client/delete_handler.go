package client

import (
	"fmt"
	"httpproject/internal/config"
	"httpproject/rest"
	"httpproject/util/logger"
	"net/http"
	"time"
)

func DeleteApiData(id string, name string) {
	time.Sleep(2 * time.Second)
	url := "http://" + config.ServerConfig.ApiInfo.ApiHost + ":" + config.ServerConfig.ApiInfo.ApiPort + "/delete"

	reqData := make(map[string]interface{})
	reqData["id"] = id
	reqData["name"] = name

	resp, statusCode, err := rest.RequetApiMethod(url, "DELETE", reqData)

	if statusCode != http.StatusOK {
		var statusMessage string
		if err != nil {
			statusMessage = fmt.Sprintf("StatusCode: %d, ErrMsg: %v", statusCode, err.Error())
		} else {
			// 400 관련 에러는 에러 메시지가 null 값으로 return 되어 statusCode만 로그에 출력
			statusMessage = fmt.Sprintf("StatusCode: %d", statusCode)
		}
		logger.Log.Error().Msgf("Delete OJT Api failed :", statusMessage)
	}

	if resp != nil {
		fmt.Println(string(resp))
	}
}
