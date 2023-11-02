package rest

import (
	"bytes"
	"httpproject/util"
	"httpproject/util/logger"
)

func RequetApiMethod(apiUri string, method string, reqBody map[string]interface{}) ([]byte, int, error) {

	var reqBodyBuffer *bytes.Buffer
	if reqBody != nil {
		reqBodyBuffer, _ = util.JSONMarshalBuffer(reqBody)
	} else {
		reqBodyBuffer = &bytes.Buffer{}
	}

	resp, statusCode, err := ExecuteService(apiUri, method, reqBodyBuffer)
	if err != nil && statusCode != 200 {
		logger.Log.Error().Msgf("Rest Execute Service Error: %s", err.Error())
		return nil, statusCode, err
	}
	if statusCode != 200 {
		return nil, statusCode, nil
	}

	return resp, statusCode, nil
}
